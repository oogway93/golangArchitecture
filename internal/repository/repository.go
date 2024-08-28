package repository

import (
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
)

type CategoryRepository interface {
	Create(requestData models.Product)
	GetAll() []map[string]interface{}
	Delete(categoryID string, productID string) error
	Get(categoryID string, productID string) string
	Update(categoryID string, productID string, newCategory models.Product) error
}

type ProductRepository interface {
	Create(requestData models.Category)
	GetAll() []map[string]interface{}
	Delete(categoryID string) error
	Get(categoryID string) string
	Update(categoryID string, newCategory models.Category) error
}

type UserRepository interface {
	Create(newUser models.User)
	GetAll() []map[string]interface{}
	Update(loginID string, newUser models.User) error
	Delete(loginID string) error
}

type Repository struct {
	CategoryRepository CategoryRepository
	UserRepository     UserRepository
	ProductRepository  ProductRepository
}
