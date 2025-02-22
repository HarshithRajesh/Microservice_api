package controllers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func HealthCheck(c *gin.Context){
  c.JSON(http.StatusOK,gin.H{"status":"User service is running"})
}
//check if i can get healthcheck library
