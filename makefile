build:
	protoc -I=. --proto_path=./Services/Authentication/Server/Module --go_out=plugins=grpc:. ./Services/Authentication/Server/Module/Registration/Registration.proto
	protoc -I=. --proto_path=./Services/Authentication/Server/Module --go_out=plugins=grpc:. ./Services/Authentication/Server/Module/Login/Login.proto