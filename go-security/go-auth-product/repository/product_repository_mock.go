package repository

import (
	"github.com/stretchr/testify/mock"
	"go-auth-product/models"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (repository *ProductRepositoryMock) FindAll() *[]models.Product {
	args := repository.Mock.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*[]models.Product)
}

func (repository *ProductRepositoryMock) FindById(id uint) *models.Product {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*models.Product)
}
