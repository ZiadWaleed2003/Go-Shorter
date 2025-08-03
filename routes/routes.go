package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShorterRequest struct {
	URL string 
	Tags string
}
func NewRouter() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", IndexHandler)
	mux.HandleFunc("POST /shortener", ShotenerHandler)
	return mux

}


func IndexHandler(w http.ResponseWriter , r *http.Request){

	fmt.Fprintln(w, "Welcome MF!")
	w.Write([]byte("Welcome to the Index Page!\n"))
}


func ShotenerHandler(w http.ResponseWriter ,r *http.Request){

	req_data := ShorterRequest{}

	fmt.Fprintln(w, "Shortener Handler")
	w.Write([]byte("This is the Shortener Handler!\n"))

	err := json.NewDecoder(r.Body).Decode(&req_data)

	if err != nil {
		panic(err)
	}

	w.Write([]byte(req_data.URL +"\n"))
	w.WriteHeader(http.StatusOK)
}