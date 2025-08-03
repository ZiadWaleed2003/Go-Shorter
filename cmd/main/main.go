package main

import (
	"net/http"
	"fmt"
)


func main(){


	fmt.Println("hey shorty")

	router := http.NewServeMux()

	router.HandleFunc("GET /", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome MF! \n"))
	})

	server := http.Server{
		Addr:   ":8080",
		Handler: router,
	}

	fmt.Println("Starting server on :8080")

	server.ListenAndServe()
}