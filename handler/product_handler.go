package handler

import (
	"go-cloud-run/repository"
	"go-cloud-run/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductList(c *gin.Context) {
	products := repository.FindProductList()
	response := response.GenerateResponseAPI(
		http.StatusOK,
		"success",
		products,
	)

	c.JSON(http.StatusOK, response)
}

func GetProductDetail(c *gin.Context) {
	productName := c.Param("name")

	productDetail, err := repository.FindProductDetail(productName)
	if err != nil {
		response := response.GenerateResponseAPI(
			http.StatusBadRequest,
			"error",
			err.Error(),
		)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := response.GenerateResponseAPI(
		http.StatusOK,
		"success",
		productDetail,
	)

	c.JSON(http.StatusOK, response)
}
