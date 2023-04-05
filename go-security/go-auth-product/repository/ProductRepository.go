package repository

import (
	"go-auth-product/models"
)

type ProductRepository interface {
	FindAll() *[]models.Product
	FindById(id uint) *models.Product
}
