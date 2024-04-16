package service

import (
	"github.com/aifuxi/banking/domain"
	"github.com/aifuxi/banking/dto"
	"github.com/aifuxi/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, error)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, error) {
	customers, err := s.repo.FindAll()

	if err != nil {
		return nil, err
	}

	var customersResp []dto.CustomerResponse

	for _, c := range customers {
		customersResp = append(customersResp, c.ToDto())
	}

	return customersResp, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppErr) {
	customer, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := customer.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
