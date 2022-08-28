package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Hello struct {
	l *log.Logger
}

func NewLogger(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello mICROSERVICES")

	dt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		os.Exit(1) 
	}
	h.l.Printf("Data %s\n", dt)

	fmt.Fprintf(rw, "Hello %s", dt)
}
