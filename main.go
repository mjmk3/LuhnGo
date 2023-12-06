package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/validate/:cardNumber", validateCardNumberHandler)

	fmt.Println("Server is running on port 8080")
	router.Run(":8080")
}
