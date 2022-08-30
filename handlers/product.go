package handlers

import (
	"fmt"
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

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&prod)
}

func (p *Product) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert the id to integer", http.StatusBadRequest)
		return
	}
	p.l.Println("Handle PUT request", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
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
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		prod := &data.Product{}

	//decode the product from JSON, if not return an error message
		err := prod.FromJSON(r.Body)
			if err != nil {
			p.l.Println("[Error in deserilizing product]", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		//validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[Error in validating product]", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating product %s", err),
				http.StatusBadRequest,
			)
			return
		}

		//add prodcut to the context
		ctx := r.Context().WithValue(KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
