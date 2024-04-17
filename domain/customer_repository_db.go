package domain

import (
	"database/sql"
	"errors"
	"github.com/aifuxi/banking/errs"
	"github.com/aifuxi/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	customers := make([]Customer, 0)

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	if err := d.client.Select(&customers, findAllSql); err != nil {
		logger.Error("error while querying customers ", zap.Field{Key: "error", String: err.Error()})
		return nil, err
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppErr) {
	var c Customer

	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	if err := d.client.Get(&c, customerSql, id); err != nil {
		logger.Debug("get customer debug" + err.Error())
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

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: dbClient}
}
