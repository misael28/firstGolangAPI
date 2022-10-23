package database

import "github.com/misael28/GoApi/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(emailId string) (*entity.User, error)
}

