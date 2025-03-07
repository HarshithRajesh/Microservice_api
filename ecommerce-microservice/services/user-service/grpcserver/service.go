package grpcserver

import (
  "context"
  "log"
  pb "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/proto/userpb"
  // pb "../../proto/userpb"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/model"
)

type UserServiceServer struct (
  pb.UnimplementedUserServerService
)

func (s *UserServiceServer) GetUserInfo (ctx context.Context, req *pb.UserRequest)(*pb.UserResponse,error){
  log.Printf("Recevied GetUserInfo request for Id : %v",req.Id)

  var user model.Users 
  u ,err := initializers.DB.First(&user,req.Id)
  if err != nil{
    if err == gorm.ErrRecordNotFound {
      return nil, errors.New("user not found")
    }
    log.Printf("Database error : %v",err)
    return nil, errors.New("Internal Server Error")
  }
   userGrpc := &pb.UserResponse{
    Id : req.Id,
    Name: u.FullName,
    Email: u.Email,
  }
  return userGrpc,nil
}
