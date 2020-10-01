# goBase

## Install Dependencies
```
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

git clone https://github.com/grpc/grpc-go.git  
cd grpc-go/cmd/protoc-gen-go-grpc  
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

## Run the proyect

```
./generate.sh 

go run greet/greet_server/server.go
go run greet/greet_client/client.go
```
