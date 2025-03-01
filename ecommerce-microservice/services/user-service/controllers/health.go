package controllers

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/initializers"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/user-service/model"
)

func HealthCheck(c *gin.Context){
  c.JSON(http.StatusOK,gin.H{"status":"User service is running"})
}
//check if i can get healthcheck library
//
func AddUser(c *gin.Context){
  var authInput model.SignUp
  
  if err := c.ShouldBindJSON(&authInput); err != nil{
    c.JSON(http.StatusBadRequest,gin.H{"error":"Input is blank"})
    return
  }
  var users model.Users
  initializers.DB.Where("email = ?",authInput.Email).First(&users)
  if users.ID != 0{
    c.JSON(http.StatusBadRequest,gin.H{"error":"User already exists"})
    return
  }
  user := model.Users{
    FullName : authInput.FullName,
    Email : authInput.Email,
    Password : authInput.Password,
  }

  if err := initializers.DB.Create(&user).Error;err != nil{
    c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to create the user"})
    return 
  }

  c.JSON(http.StatusOK,gin.H{"status":1,"User":user})
  

}
