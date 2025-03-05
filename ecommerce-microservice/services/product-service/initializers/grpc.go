package initializers

import (
  "log"
  pb "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/proto/userpb"
  "google.golang.org/grpc"
)

var UserClient pb.UserServiceClient

func GRPCClient(){
  conn,err := grpc.Dial("user-service:50051",grpc.WithTransportCredentials(insecure.NewCredentials())
  if err != nil {
    log.Fatalf("failed to connect to grpc server at localhost : 50051 -> %v",err)
  }
  defer conn.Close()

  UserClient = pb.NewUserServieClient(conn)
  log.Println("Connected to user service through gRPC")
}
