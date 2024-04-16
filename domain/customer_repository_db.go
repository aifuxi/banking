package domain

import (
	"database/sql"
	"errors"
	"github.com/aifuxi/banking/errs"
	"github.com/aifuxi/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("error while querying customers ", zap.Field{Key: "error", String: err.Error()})
		return nil, err
	}
	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Info("error while scanning customers ", zap.Field{Key: "error", String: err.Error()})
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
		// check query row error
		if errors.Is(err, sql.ErrNoRows) {
			logger.Error("sql.ErrNoRows" + sql.ErrNoRows.Error())
			return nil, errs.NewNotFoundErr("Customer not found")
		} else {
			logger.Error("error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectErr("Unexpect databases error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:123456@/banking")
	if err != nil {
		logger.Error("sql open error: " + err.Error())
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
