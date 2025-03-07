package main

import (
  // "fmt"
  "github.com/gin-gonic/gin"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/controllers"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/initializers"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/grpcserver"

)
func init(){
  initializers.LoadEnvs()
  initializers.ConnectDB()
  grpcserver.GrpcServer()
}
func main(){
  r:=gin.Default()
  r.GET("/health",controllers.HealthCheck)
  r.POST("/add",controllers.AddUser)
  r.Run()
}
