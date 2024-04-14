package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	ZipCode string
}

func main() {

	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World ~")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Tom", "bj", "001"},
		{"Jerry", "sh", "002"},
		{"GG Bound", "hz", "0010"},
	}
	// add content type header
	w.Header().Add("Content-Type", "application/json")

	// customers -> response
	json.NewEncoder(w).Encode(customers)
}
