package main

import (
  // "fmt"
  "github.com/gin-gonic/gin"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/controllers"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/initializers"

)
func init(){
  initializers.LoadEnvs()
  initializers.ConnectDB()
  initializers.GrpcServer()
}
func main(){
  r:=gin.Default()
  r.GET("/health",controllers.HealthCheck)
  r.POST("/add",controllers.AddUser)
  r.Run()
}
