package dto

import (
	"database/sql"
	"errors"
	"time"

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

func NewUserDTO() *UserDTO {
	return &UserDTO{}
}

func (u *UserDTO) findUserByID(userID string) (*UserResponseDTO, error) {
	if userID == "" {
		return nil, errors.New("userID is empty")
	}

	u.userResponse = &UserResponseDTO{}

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgre password=admin dbname=user sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, first_name, last_name, email, date_of_birth FROM users WHERE id = $1", userID)
	if err != nil {
		panic(err)
	}

	if !rows.Next() {
		return nil, errors.New("user not found")
	}

	err = rows.Scan(&u.userResponse.ID, &u.userResponse.FirstName, &u.userResponse.LastName, &u.userResponse.Email, &u.userResponse.DateOfBirth)

	if err != nil {
		panic(err)
	}

	u.userResponse.Fullname = u.userResponse.FirstName + " " + u.userResponse.LastName
	u.userResponse.Age = time.Now().Year() - u.userResponse.DateOfBirth.Year()

	return u.userResponse, nil
}

func (u *UserDTO) createUser(userCreate *UserCreateDTO) (bool, error) {
	if userCreate.FirstName == "" || userCreate.LastName == "" || userCreate.Email == "" || userCreate.Password == "" || userCreate.DateOfBirth.IsZero() {
		return false, errors.New("invalid user data")
	}

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgre password=admin dbname=user sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("INSERT INTO users (first_name, last_name, email, password, date_of_birth) VALUES ($1, $2, $3, $4, $5) RETURNING id", userCreate.FirstName, userCreate.LastName, userCreate.Email, userCreate.Password, userCreate.DateOfBirth)
	if err != nil {
		panic(err)
	}

	if !rows.Next() {
		return false, errors.New("failed to create user")
	}

	return true, nil
}
