package initializers

import(
  "github.com/joho/godotenv"
  "log"
)

func LoadEnvs(){
  err := godotenv.Load()
  if err != nil{
    log.Fatalf("Error loading the environment variables %v",err)
  }
}
