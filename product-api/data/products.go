package data

import (
	"encoding/json"
	"io"
	"time"
)

//Product define the structure for an API product
type Product struct {
	ID          int     `json: "id"`
	Name        string  `json: "name"`
	Description string  `json: "description`
	Price       float32 `json: "price`
	SKU         string  `json: "sku`
	CreatedOn   string  `json: "_"`
	UpdatedOn   string  `json: "_"`
	DeletedOn   string  `json: "_"`
}

type Products []*Product

func (p *Products) TOJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) FromJSON(r io.Reader) error{
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func GetProducts() Products {
	return productList
}

func AddProduct(p *Product){

	p.ID = getNextID()
	productList = append(productList, p)
 
}

func getNextID() int{
	lp := productList[len(productList) - 1]

	return lp.ID+1 
}

//initializing a product
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky Coffee",
		Price:       4500.5,
		SKU:         "abc234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Esppresso",
		Description: "Short and strong coffee without milk",
		Price:       2500.5,
		SKU:         "def234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
