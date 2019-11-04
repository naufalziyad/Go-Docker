package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Testing auto reload")
	})

	r.HandleFunc("/cek1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Page Cek 1")
	})

	r.HandleFunc("/cek2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Page Cek 2")
	})

	fmt.Println("Server Listening !")
	http.ListenAndServe(":80", r)
}
