package main

import (
	"fmt"
	"go-cloud-run/handler"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/", handler.GetProductList)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	port = fmt.Sprintf(":%v", port)

	if err := engine.Run(port); err != nil {
		log.Fatal(err)
	}
}
