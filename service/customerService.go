package service

import (
	"capi/domain"
	"capi/dto"
	"capi/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]dto.CustomerResponse, *errs.AppErr)
	GetCustomerByID(string) (*dto.CustomerResponse, *errs.AppErr)
}

type DefaultCustomerService struct {
	repository domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]dto.CustomerResponse, *errs.AppErr) {
	// * add process here
	custs, err := s.repository.FindAll()
	if err != nil {
		return nil, err
		// logger.Error("error scanning customer data")
		// return nil, errs.NewUnexpectedError("unexpected database error")
	}

	var response []dto.CustomerResponse
	for _, c := range custs {
		response = append(response, c.ToDTO())
	}
	return response, nil
}

func (s DefaultCustomerService) GetCustomerByID(customerID string) (*dto.CustomerResponse, *errs.AppErr) {
	cust, err := s.repository.FindByID(customerID)
	if err != nil {
		return nil, err
	}

	response := cust.ToDTO()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository: repository}
}
