package app

import (
	"github.com/aifuxi/banking/domain"
	"github.com/aifuxi/banking/logger"
	"github.com/aifuxi/banking/service"
	"github.com/gorilla/mux" // third party package must use full name
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

func Start() {

	r := mux.NewRouter()

	client, err := sqlx.Open("mysql", "root:123456@/banking")
	if err != nil {
		logger.Error("sql open error: " + err.Error())
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	customerRepoDb := domain.NewCustomerRepositoryDb(client)
	accountRepoDb := domain.NewAccountRepositoryDb(client)

	// wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStud())}
	ch := CustomerHandler{service: service.NewCustomerService(customerRepoDb)}
	ah := AccountHandler{service: service.NewAccountService(accountRepoDb)}

	r.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	r.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
