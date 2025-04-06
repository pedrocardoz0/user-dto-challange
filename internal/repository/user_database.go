package repository

import (
	"database/sql"
	"time"
)

type UserDatabaseRepository struct {
	db *sql.DB
}

func NewUserRepository() *UserDatabaseRepository {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgre password=admin dbname=user sslmode=disable")
	if err != nil {
		panic(err)
	}

	return &UserDatabaseRepository{db: db}
}

func (u *UserDatabaseRepository) Update(id string, user *UserUpdate) error {
	query := `UPDATE users SET first_name = $1, last_name = $2, email = $3, password = $4, date_of_birth = $5, updated_at = $6 WHERE id = $7`
	_, err := u.db.Exec(query, user.FirstName, user.LastName, user.Email, user.Password, user.DateOfBirth, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserDatabaseRepository) FindByID(userID string) (*UserResponse, error) {
	query := `SELECT id, first_name, last_name, email, date_of_birth FROM users WHERE id = $1`
	var user UserResponse
	err := u.db.QueryRow(query, userID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateOfBirth)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserDatabaseRepository) Create(userCreate *UserCreate) (bool, error) {
	query := `INSERT INTO users (first_name, last_name, email, password, date_of_birth) VALUES ($1, $2, $3, $4, $5)`
	_, err := u.db.Exec(query, userCreate.FirstName, userCreate.LastName, userCreate.Email, userCreate.Password, userCreate.DateOfBirth)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (u *UserDatabaseRepository) FindAll() ([]*UserResponse, error) {
	query := `SELECT id, first_name, last_name, email, date_of_birth FROM users LIMIT 100`
	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	users := []*UserResponse{}
	for rows.Next() {
		var user UserResponse
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateOfBirth)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
