package initializers

import (
  "log"
  "google.golang.org/grpc"
  // pb "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/proto/userpb"
  // pb "../../proto/userpb"
  "ecommerce-microservice/proto/userpb"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/grpcserver"
)

func GrpcServer(){
  lis,err := net.Listen("tcp",":50051")
  if err != nil {
    log.Fatalf("failed to listen: %v",err)
  }
  grpcServer := grpc.NewServer()
  pb.RegisterUserServiceServer(grpcServer, &grpcserver.UserServerService)
  log.Println("User Service gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
