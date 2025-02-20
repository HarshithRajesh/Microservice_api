package initializers

import (
  "log"
  "github.com/joho/godotenv"
)

func LoadEnvs(){
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("Error loading the env %v",err)
  }
}
