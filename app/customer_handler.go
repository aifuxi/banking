package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/aifuxi/banking/logger"
	"github.com/aifuxi/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
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

func (ch *CustomerHandler) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	logger.Info("customer_id from path: " + id)

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, code int, data any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("json encode error: " + err.Error())
	}
}
