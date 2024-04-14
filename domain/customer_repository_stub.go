package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStud() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "001", Name: "Tom", City: "1001", Zipcode: "2000-01-01", DateOfBirth: "1"},
		{Id: "001", Name: "Tom", City: "1001", Zipcode: "2000-01-01", DateOfBirth: "1"},
	}

	return CustomerRepositoryStub{customers}
}
