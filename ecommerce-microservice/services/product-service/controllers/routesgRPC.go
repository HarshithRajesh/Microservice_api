package controllers

import (
  "context"
  "net/http"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/proto/userpb"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/product-service/initializers"
  "time"
  "github.com/gin-gonic/gin"
)

func GetUserDetails(c *gin.Context){
  userId := c.Param("user_id")
  ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
  defer cancel

  res,err := initializers.UserClient.GetUserInfo(ctx,&userpb.UserRequest{Id:userId})
  if err != nil {
    c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to fetch user details"})
    return
  }
  c.JSON(http.StatusOk,res)
}
