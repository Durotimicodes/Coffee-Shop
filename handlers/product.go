package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/durotimicodes/microservices/product-api/data"
	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.TOJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to unmarshall json", http.StatusInternalServerError)
	}
}

func (p *Product) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product!")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal message", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Product) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert the id to integer", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT request", id)

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
	}
}


type KeyProduct struct{}

func (p *Product) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(rw http.ResponseWriter, r *http.Request){
		prod := &data.Product{}

	//decode the product from JSON, if not return an error message
		err = prod.FromJSON(r.Body)
			if err != nil {
			http.Error(rw, "Unable to unmarshal json body", http.StatusBadRequest)
			return
		}

		ctx := r.Context().WithValue(KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, re)
	}
}
