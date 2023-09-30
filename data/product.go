package data

type product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var productList = map[string]product{
	"mazda": {
		ID:    1,
		Name:  "mazda",
		Price: 4000000000,
		Stock: 12,
	},
	"supra": {
		ID:    2,
		Name:  "supra",
		Price: 6000000000,
		Stock: 2,
	},
	"gtr": {
		ID:    3,
		Name:  "gtr",
		Price: 1000000000,
		Stock: 5,
	},
}

func LoadProductList() map[string]product {
	return productList
}
