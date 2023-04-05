package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-auth-product/models"
	"go-auth-product/repository"
	"testing"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductService_GetOneProduct_NotFound(t *testing.T) {
	productRepository.Mock.On("FindById", uint(1)).Return(nil)
	product, err := productService.GetOneProduct(uint(1))
	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error())
}

func TestProductService_GetOneProduct(t *testing.T) {
	product := models.Product{
		Title:       "test",
		Description: "test",
	}
	productRepository.Mock.On("FindById", product.ID).Return(&product)
	result, err := productService.GetOneProduct(product.ID)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.Title, result.Title)
	assert.Equal(t, product.Description, result.Description)
	//assert.Equal(t, product, result)
}
