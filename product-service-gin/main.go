package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pobsaeng.com/product-api/config"
	"pobsaeng.com/product-api/controller"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const CustomHeaderValue = "Phob@#12$"

func init() {
	if err := godotenv.Load(); err != nil {
	    log.Fatalf("Error loading .env file")
	}
 }
 
func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
	
	config.InitDB()

	r := gin.Default()
	r.Use(corsMiddleware())

	v1 := r.Group("/api/v1")
	v1.GET("/products", controller.GetProducts)
	v1.GET("/products/:id", controller.GetProductByID)
	v1.POST("/products/search", controller.SearchProducts)
	v1.POST("/products", controller.CreateProduct)
	v1.PUT("/products/:id", controller.UpdateProduct)
	v1.DELETE("/products/:id", controller.DeleteProduct)

	r.Run(":" + os.Getenv("APP_PORT"))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// var headers struct {
		// 	CustomHeader string `header:"X-Custom-Header"`
		// }

		// if err := c.ShouldBindHeader(&headers); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing headers"})
		// 	c.Abort() // Abort processing
		// 	return
		// }

		// if headers.CustomHeader != CustomHeaderValue {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid X-Custom-Header"})
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}
