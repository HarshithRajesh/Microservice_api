package initializers

import (
  "log"
  "os"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB 

func ConnectDB(){
  dsn := os.Getenv("DATABASE_URL")
  if dsn == ""{
    log.Fatal("The user database url is not fetched")
  } 
  log.Println("Connecting to the database ...")

  var err error 
  DB,err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
  if err != nil{
    log.Fatalf("Failed to connect to the database : %v",err)
  }
  log.Println("Successfully connected to the database")
}
