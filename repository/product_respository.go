package repository

import (
	"errors"
	"go-cloud-run/data"
)

func FindProductList() map[string]data.Product {
	return data.ProductList
}

func FindProductDetail(productName string) (data.Product, error) {
	products := FindProductList()

	if _, isExist := products[productName]; isExist {
		return products[productName], nil
	}

	return data.Product{}, errors.New("products available only mazda, supra, gtr")
}
