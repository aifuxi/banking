package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/aifuxi/banking/service"
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
