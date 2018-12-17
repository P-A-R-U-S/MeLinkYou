package main

import (
	pb "MeLinkYou/Services/Authentication/Server/Module/Registration"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	ADDRESS              = "localhost:50051"
	REGISTER_FILE_NAME   = "register.json"
	UNREGISTER_FILE_NAME = "unregister.json"
)

var (
	registerClient pb.RegisterUserClient
)

func parseFile(file string) (*pb.RegisterUserRequest, error) {
	var registerUser *pb.RegisterUserRequest
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &registerUser)
	return registerUser, err
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(ADDRESS, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	registerClient = pb.NewRegisterUserClient(conn)

	// Print welcome message.
	fmt.Println("Cloud Datastore Task List")
	fmt.Println()
	usage()

	// Read commands from stdin.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")

	for scanner.Scan() {
		cmd, args := parseCmd(scanner.Text())
		switch cmd {
		case "register":

			Register(args)

		case "unregister":

			UnRegister(args)

		default:
			log.Printf("Unknown command %q", cmd)
			usage()
		}

		fmt.Print("> ")
	}

}

func usage() {
	fmt.Print(`Usage:
  ----------------------------------------------
  register <file>  			- register new user, by default: register.json
  register <file>			- marks a task as done, by default: unregister.json
  signIn					- lists all tasks by creation time
  signOut 			       	- deletes a task
`)
}

// parseCmd splits a line into a command and optional extra args.
func parseCmd(line string) (cmd, args string) {
	if f := strings.Fields(line); len(f) > 0 {
		cmd = f[0]
		args = strings.Join(f[1:], " ")
	}
	return cmd, args
}

func Register(fileName string) {
	//Contact the server and print out its response.

	if fileName == "" {
		fileName = REGISTER_FILE_NAME
	}

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	user, err := parseFile(fileName)
	if err != nil {
		log.Fatalf("Could not parse the file: %v", fileName)
	}

	r, err := registerClient.Register(context.Background(), user)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("User: %v, Tocken:%v, Status: %t, Error:%s",
		user.Email, r.Token, r.Status.Success, r.Status.Error)

}

func UnRegister(fileName string) {
	fmt.Println("Not implemented...")
}
