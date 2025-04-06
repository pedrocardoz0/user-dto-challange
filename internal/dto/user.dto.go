package dto

import (
	"errors"
	"time"

	"user-dto-challange/internal/repository"

	_ "github.com/lib/pq"
)

type UserDTO struct {
	userResponse *UserResponseDTO
	userCreate   *UserCreateDTO
}

type UserResponseDTO struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Fullname    string
	Age         int
}

type UserCreateDTO struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

type UserUpdateDTO struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	DateOfBirth time.Time `json:"date_of_birth"`
}

func NewUserDTO() *UserDTO {
	return &UserDTO{}
}

func (u *UserDTO) findUserByID(userID string) (*UserResponseDTO, error) {
	if userID == "" {
		return nil, errors.New("userID is empty")
	}

	userRepository := repository.NewUserRepository()
	user, err := userRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}

	u.userResponse = &UserResponseDTO{
		ID:          user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Fullname:    user.FirstName + " " + user.LastName,
		Age:         time.Now().Year() - user.DateOfBirth.Year(),
	}

	return u.userResponse, nil
}

func (u *UserDTO) createUser(userCreate *UserCreateDTO) (bool, error) {
	if userCreate.FirstName == "" || userCreate.LastName == "" || userCreate.Email == "" || userCreate.Password == "" || userCreate.DateOfBirth.IsZero() {
		return false, errors.New("invalid user data")
	}

	userRepository := repository.NewUserRepository()
	_, err := userRepository.Create(&repository.UserCreate{
		FirstName:   userCreate.FirstName,
		LastName:    userCreate.LastName,
		Email:       userCreate.Email,
		Password:    userCreate.Password,
		DateOfBirth: userCreate.DateOfBirth,
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
