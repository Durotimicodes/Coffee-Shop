package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/durotimicodes/microservices/handlers"
	"github.com/gorilla/mux"
)

// var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	//create a logger
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	//reference the handler
	ph := handlers.NewProduct(l)

	//define a new serve mux and register the new handler into it
	sm := mux.NewRouter()

	//router for Get request
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	//router for put request
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareProductValidation)

	//route for post request
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/addproduct", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	/*in order to prevent blocked connections due to clients interruption set a timeout by creating
	a server */
	//create a new server
	s := &http.Server{
		Addr:         ":9091",           //configure the bind address
		Handler:      sm,                //set the deafult handler
		ErrorLog:     l,                 //set the logger for the server
		IdleTimeout:  120 * time.Second, // max time for connection using TCP keep-alive
		ReadTimeout:  5 * time.Second,   // set max time to read request fro client
		WriteTimeout: 10 * time.Second,  // set max time to write response to client
	}

	//start the sever
	go func() {
		log.Println("Server running at port 9090...")
		//since this will block, put it in a goroutine
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error is starting server %s\n", err)
			os.Exit(1)
		}
	}()

	//this will basically broadcast the message during a operation kill or interrup command in passed
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	//graceful shutdown
	//to allow smooth disconnection of client
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
