package handler

import (
	"go-cloud-run/data"
	"go-cloud-run/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductList(c *gin.Context) {
	products := data.LoadProductList()

	response := response.GenerateResponseAPI(
		http.StatusOK,
		"success",
		products,
	)

	c.JSON(http.StatusOK, response)
}
