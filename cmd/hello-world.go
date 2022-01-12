package main

import (
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!"))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}