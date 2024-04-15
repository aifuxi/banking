package domain

import (
	"database/sql"
	"errors"
	"github.com/aifuxi/banking/errs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("error while querying customers ", err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("error while scanning customers ", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppErr) {

	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundErr("Customer not found")
		} else {
			log.Println("error while scanning customers ", err.Error())
			return nil, errs.NewUnexpectErr("Unexpect databases error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:123456@/banking")
	if err != nil {
		log.Fatalln("sql open error: ", err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
