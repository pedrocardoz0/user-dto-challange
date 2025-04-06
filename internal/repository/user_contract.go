package repository

import (
	"time"
)

type UserRepository interface {
	Update(user *UserUpdate) error
	FindByID(userID string) (*UserResponse, error)
	Create(userCreate *UserCreate) (bool, error)
	FindAll() ([]*UserResponse, error)
}

type UserUpdate struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type UserCreate struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type UserResponse struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
