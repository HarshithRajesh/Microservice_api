package controllers

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func HealthCheck(c *gin.Context){
  c.JSON(http.StatusOK,gin.H{"staus":"Product service is working"}) 
}
