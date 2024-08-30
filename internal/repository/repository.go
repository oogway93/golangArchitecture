package repository

import (
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
)

type ProductRepository interface {
	Create(categoryID string, requestData models.Product) error
	GetAll(categoryID string) []map[string]interface{}
	Delete(categoryID string, productID string) error
	Get(categoryID string, productID string) map[string]interface{}
	Update(newCategoryName string, productID string, newProduct models.Product) error
}

type CategoryRepository interface {
	Create(requestData models.Category)
	GetAll() []map[string]interface{}
	Delete(categoryID string) error
	Get(categoryID string) map[string]interface{}
	Update(categoryID string, newCategory models.Category) error
}

type UserRepository interface {
	Create(newUser models.User)
	GetAll() []map[string]interface{}
	Get(loginID string) map[string]interface{}
	Update(loginID string, newUser models.User) error
	Delete(loginID string) error
}

type AuthRepository interface {
	Login(loginID string) map[string]interface{}
}

type Repository struct {
	CategoryRepository CategoryRepository
	UserRepository     UserRepository
	ProductRepository  ProductRepository
	AuthRepository     AuthRepository
}
