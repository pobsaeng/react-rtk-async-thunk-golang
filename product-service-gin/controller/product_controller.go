package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pobsaeng.com/product-api/model"
	"pobsaeng.com/product-api/repository"
)

func GetProducts(c *gin.Context) {
	products, err := repository.GetAllProducts()
	if err != nil {
		log.Println("Error:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Fatalf("Error converting string to uint64: %v", err)
	}
	
	product, err := repository.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var product model.Product
	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	if err := repository.CreateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully"})
}

func UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	var updatedProduct model.Product
	if err := c.BindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request payload"})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Fatalf("Error converting string to uint64: %v", err)
	}

	if err := repository.UpdateProduct(id, updatedProduct); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

func DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Fatalf("Error converting string to uint64: %v", err)
	}

	if err := repository.DeleteProduct(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
