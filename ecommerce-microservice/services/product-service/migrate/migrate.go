package main 

import (
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/product-service/model"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/product-service/initializers"
  "log"
)

func init (){
  initializers.LoadEnvs()
  initializers.ConnectDB()
}

func main(){
  log.Println("Starting the migration")
  err := initializers.DB.AutoMigrate(&model.Products{})
  if err != nil {
    log.Fatal("Failed migration")
  }else {
    log.Println("migration completed successfully")
  }
}
