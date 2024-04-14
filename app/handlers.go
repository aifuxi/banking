package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipCode" xml:"zipCode"`
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

	// get  content type for request header
	if r.Header.Get("Content-Type") == "application/xml" {
		// add content type header
		w.Header().Add("Content-Type", "application/xml")
		// customers -> response
		xml.NewEncoder(w).Encode(customers)
	} else {
		// add content type header
		w.Header().Add("Content-Type", "application/json")
		// customers -> response
		json.NewEncoder(w).Encode(customers)
	}

}
