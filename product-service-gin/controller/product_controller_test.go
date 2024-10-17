// controller_test.go

package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"pobsaeng.com/product-api/repository"
)

func TestGetProducts(t *testing.T) {
	// Initialize a new Gin router
	router := gin.Default()

	// Set up a mock repository
	repository.MockInit()

	// Define the route
	router.GET("/api/v1/products", GetProducts)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/api/v1/products", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	println("StatusOK : ", http.StatusOK)
	println("Code : ", w.Code)
	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateProduct(t *testing.T) {
	// Initialize a new Gin router
	router := gin.Default()

	// Set up a mock repository
	repository.MockInit()

	// Define the route
	router.POST("/api/v1/products", CreateProduct)

	// Define a sample product
	var jsonStr = []byte(`{
		"id": "1",
		"name": "Sample Product",
		"image": "sample.jpg",
		"price": 10.0,
		"store": 100,
		"type": "Sample Type"
	}`)

	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// Set the content type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the HTTP request
	router.ServeHTTP(w, req)

	// Check the HTTP status code
	assert.Equal(t, http.StatusOK, w.Code)
}
