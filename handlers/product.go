package handlers

import (
	"log"
	"net/http"
	"regexp"

	"github.com/durotimicodes/microservices/product-api/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		//expect the id in the URL

		reg := regexp.MustCompile(`/[0-9]+`)
		reg.FindAllStringSubmatch(r.URL.Path)

	}

	//else catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.TOJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusInternalServerError)
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product!")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal message", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}
