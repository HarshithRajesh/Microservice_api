package main 


import (
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/initializers"
  "log"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/model"
)


func init(){
  initializers.LoadEnvs()
  initializers.ConnectDB()
}

func main(){
  log.Println("Starting the migration")
  err := initializers.DB.AutoMigrate(&model.Users{})
  if err != nil{
    log.Fatal("Migration is failed")
  } else {
    log.Println("Migration completed")
  }
}
