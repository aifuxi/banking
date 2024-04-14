package app

import (
	"github.com/aifuxi/banking/domain"
	"github.com/aifuxi/banking/service"
	"github.com/gorilla/mux" // third party package must use full name
	"log"
	"net/http"
)

func Start() {

	r := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStud())}

	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
