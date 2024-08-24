package repository

import (
	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
)


type CategoryRepository interface {
	GetAll() ([]products.Category,error)
}

type UserRepository interface {
	Create(user models.User) ()
}

type Repository struct {
	CategoryRepository CategoryRepository
	UserRepository UserRepository
}


