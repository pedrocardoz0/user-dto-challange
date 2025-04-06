package repository

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	userRepository := NewUserRepository()
	assert.NotNil(t, userRepository)
}

func TestUserRepository_Update(t *testing.T) {
	userRepository := NewUserRepository()
	user := &UserUpdate{
		FirstName: "Updated",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
	}
	err := userRepository.Update("2d846fc0-3315-4fe2-b9e2-d76505757e38", user)
	assert.Nil(t, err)
}

func TestUserRepository_FindByID(t *testing.T) {
	userRepository := NewUserRepository()
	user, err := userRepository.FindByID("2d846fc0-3315-4fe2-b9e2-d76505757e38")
	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestUserRepository_Create(t *testing.T) {
	userRepository := NewUserRepository()
	user := &UserCreate{
		FirstName:   "John",
		LastName:    "Doe",
		Email:       "test5@test.com",
		Password:    "password",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	created, err := userRepository.Create(user)
	assert.Nil(t, err)
	assert.True(t, created)
}

func TestUserRepository_FindAll(t *testing.T) {
	userRepository := NewUserRepository()
	users, err := userRepository.FindAll()
	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.Equal(t, len(users), 6)
}
