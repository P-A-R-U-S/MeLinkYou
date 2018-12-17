package main

import (
	app "MeLinkYou/Services/Authentication/Server/App"
	command "MeLinkYou/Services/Authentication/Server/Commands"
	handlers "MeLinkYou/Services/Authentication/Server/Handlers"
	registration "MeLinkYou/Services/Authentication/Server/Module/Registration"
	"MeLinkYou/Services/Authentication/Server/Queries"
	"cloud.google.com/go/datastore"
	"crypto/rand"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

const (
	port = ":50051"
)

var (
	eventBus app.EventBus

	cloudProjectID         string
	dataStoreClient        *datastore.Client
	dataStoreContext       context.Context
	addUserHandler         handlers.AddUserHandler
	updateUserTokenHandler handlers.UpdateUserTokenHandler
)

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type registrationServer struct {
}

func (s *registrationServer) Register(context context.Context, request *registration.RegisterUserRequest) (*registration.RegisterUserResponse, error) {

	log.Println("Register...")
	ret := registration.RegisterUserResponse{Status: &registration.Status{Success: true, Error: ""}}

	//Validation
	query := Queries.UserExistQuery{Client: dataStoreClient, Context: dataStoreContext}

	isUserExist, err := query.Query(request.Email)

	if err != nil {
		ret.Status.Success = false
		ret.Status.Error = fmt.Sprint(err)
		return &ret, err
	}

	if isUserExist {
		ret.Status.Success = false
		ret.Status.Error = "User already exist."
		return &ret, nil
	}

	// Command
	command := command.AddUserCommand{
		Email:    request.Email,
		Password: request.Password,
	}

	eventBus.Publish(command)

	b := make([]byte, 16)
	_, err = rand.Read(b)
	if err != nil {
		log.Println("Error: ", err)
	}

	ret.Token = fmt.Sprintf("%X%X%X%X%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	ret.Status.Success = true
	ret.Status.Error = fmt.Sprint(err)

	return &ret, nil
}

func (s *registrationServer) UnRegister(context.Context, *registration.UnRegisterUserRequest) (*registration.UnRegisterUserResponse, error) {
	log.Println("UnRegister...")

	status := &registration.Status{Success: true, Error: ""}

	return &registration.UnRegisterUserResponse{Status: status}, nil
}

func main() {

	// Register handlers
	eventBus = app.New()

	InitDataStoreClient()

	//SignUp handler
	addUserHandler = handlers.AddUserHandler{Client: dataStoreClient, Context: dataStoreContext}
	eventBus.Subscribe(&addUserHandler)
	defer eventBus.Unsubscribe(addUserHandler.Event())

	updateUserTokenHandler = handlers.UpdateUserTokenHandler{Client: dataStoreClient, Context: dataStoreContext}
	eventBus.Subscribe(&updateUserTokenHandler)
	defer eventBus.Unsubscribe(updateUserTokenHandler.Event())

	//set-up gRPC server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listem: %v", err)
	}
	s := grpc.NewServer()

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	registration.RegisterRegisterUserServer(s, &registrationServer{})

	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("RPC:%s server started.", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func InitDataStoreClient() {

	log.Printf("Google Cloud Setting: DATASTORE_DATASET = %s", os.Getenv("DATASTORE_DATASET"))
	log.Printf("Google Cloud Setting: DATASTORE_EMULATOR_HOST = %s", os.Getenv("DATASTORE_EMULATOR_HOST"))
	log.Printf("Google Cloud Setting: DATASTORE_EMULATOR_HOST_PATH = %s", os.Getenv("DATASTORE_EMULATOR_HOST_PATH"))
	log.Printf("Google Cloud Setting: DATASTORE_HOST = %s", os.Getenv("DATASTORE_HOST"))
	log.Printf("Google Cloud Setting: DATASTORE_PROJECT_ID = %s", os.Getenv("DATASTORE_PROJECT_ID"))

	cloudProjectID = os.Getenv("DATASTORE_PROJECT_ID")
	if cloudProjectID == "" {
		log.Fatal(`You need to set the environment variable "DATASTORE_PROJECT_ID"`)
		return
	}

	// [START build_service]
	dataStoreContext = context.Background()

	// Create a datastore client. In a typical application, you would create
	// a single client which is reused for every datastore operation.
	client, err := datastore.NewClient(dataStoreContext, cloudProjectID)

	// [END build_service]
	if err != nil {
		log.Fatalf("Could not create datastore client: %v", err)
	}

	dataStoreClient = client
}
