package service

import (
	"errors"
	"go-auth-product/models"
	"go-auth-product/repository"
)

type ProductService struct {
	Repository repository.ProductRepository
}

func (service ProductService) GetOneProduct(id uint) (*models.Product, error) {
	product := service.Repository.FindById(id)
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (service ProductService) GetAllProducts() (*[]models.Product, error) {
	products := service.Repository.FindAll()
	if products == nil {
		return nil, errors.New("products not found")
	}
	return products, nil
}
