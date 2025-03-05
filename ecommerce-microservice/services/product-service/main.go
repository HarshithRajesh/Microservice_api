package main

import (
  "github.com/gin-gonic/gin"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/product-service/controllers"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/product-service/initializers"
)

func init(){
  initializers.LoadEnvs()
  initializers.ConnectDB()
  initializers.GRPCClient()

}

func main(){
  r := gin.Default()
  r.GET("/health",controllers.HealthCheck)
  r.POST("/add",controllers.AddProduct)
  r.GET("/users/:user_id",controllers.GetuserDetails)
  r.Run(":8081")
}
