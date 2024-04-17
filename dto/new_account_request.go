package dto

import (
	"github.com/aifuxi/banking/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppErr {
	if r.Amount < 5000 {
		return errs.NewValidationErr("To open a new account you need to deposit at least 5000.00")
	}

	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationErr("Account type should be checking or saving")
	}

	return nil
}
