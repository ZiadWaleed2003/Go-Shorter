package routes


import (
	"net/http"
	"fmt"
)



func NewRouter() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)

	return mux

}


func IndexHandler(w http.ResponseWriter , r * http.Request){

	fmt.Fprintln(w, "Welcome MF!")
	w.Write([]byte("Welcome to the Index Page!\n"))
}
