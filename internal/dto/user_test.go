package dto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetUserResponseDTO(t *testing.T) {
	userID := "2d846fc0-3315-4fe2-b9e2-d76505757e38"

	userDTO := NewUserDTO()
	userResponse, err := userDTO.findUserByID(userID)

	assert.NotNil(t, userResponse)
	assert.Equal(t, userResponse.ID, userID)
	assert.Equal(t, userResponse.Fullname, "John Doe")
	assert.Equal(t, userResponse.Age, 35)
	assert.NoError(t, err)
}

func TestGetUserResponseDTO_NotFound(t *testing.T) {
	userID := "3d246fc0-3315-4fe2-b9e2-d76505757e38"

	userDTO := NewUserDTO()
	userResponse, err := userDTO.findUserByID(userID)

	assert.Nil(t, userResponse)
	assert.Error(t, err)
}

func TestGetUserResponseDTO_EmptyUserID(t *testing.T) {
	userID := ""

	userDTO := NewUserDTO()
	userResponse, err := userDTO.findUserByID(userID)

	assert.Nil(t, userResponse)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "userID is empty")
}

func TestUserCreateDTO_Create(t *testing.T) {
	userCreate := &UserCreateDTO{
		FirstName:   "Test",
		LastName:    "Test",
		Email:       "test2@test.com",
		Password:    "password",
		DateOfBirth: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	userDTO := NewUserDTO()
	userResponse, err := userDTO.createUser(userCreate)

	assert.NotNil(t, userResponse)
	assert.NoError(t, err)
}

func TestUserCreateDTO_Create_InvalidData(t *testing.T) {
	userCreate := &UserCreateDTO{
		FirstName:   "",
		LastName:    "",
		Email:       "",
		Password:    "",
		DateOfBirth: time.Time{},
	}

	userDTO := NewUserDTO()
	userResponse, err := userDTO.createUser(userCreate)

	assert.False(t, userResponse)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "invalid user data")
}
