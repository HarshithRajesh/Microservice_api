package controllers

import(
  "github.com/gin-gonic/gin"
  "net/http"
   "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/product-service/initializers"
  "github.com/HarshithRajesh/Microservice_api/ecommerce-microservice/product-service/model" 
)

func AddProduct(c *gin.Context){
  var input model.CreateProduct

  if err := c.ShouldBindJSON(&input);err != nil{
    c.JSON(http.StatusBadRequest,gin.H{"Error":err})
    return
  }

  var products model.Product
  initializers.DB.Where("name = ?",input.Name).First(&products)
  if products.ID != 0{
    c.JSON(http.StatusBadRequest,gin.H{"error":"Product already exists"})
    return
  }
  product := model.Product{
   Name  :  input.Name,
   Description :  input.Description,
   Price :  input.Price,
   Stock :  input.Stock,
  }
  if err := initializers.DB.Create(&product).Error;err != nil{
    c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to create the product"})
    return
  }
  c.JSON(http.StatusOK,gin.H{"status":1,"message":product})
}
