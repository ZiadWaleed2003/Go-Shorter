package main

import (
	"net/http"
	"fmt"
	"go-shorter/routes"
)


func main(){

	fmt.Println("hey shorty")
	router := routes.NewRouter()


	server := http.Server{
		Addr:   ":8080",
		Handler: router,
	}

	fmt.Println("Starting server on :8080")

	server.ListenAndServe()
}