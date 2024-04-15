package service

import (
	"github.com/aifuxi/banking/domain"
	"github.com/aifuxi/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(id string) (*domain.Customer, *errs.AppErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repo.FindAll()
}

func (s DefaultCustomerService) GetCustomer(id string) (*domain.Customer, *errs.AppErr) {
	return s.repo.ById(id)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
