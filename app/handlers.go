package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/aifuxi/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zipCode" xml:"zipCode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

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

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	log.Printf("customer_id from path: %v\n", id)

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
		return
	}

	// add content type header
	w.Header().Add("Content-Type", "application/json")
	// customers -> response
	json.NewEncoder(w).Encode(customer)

}
