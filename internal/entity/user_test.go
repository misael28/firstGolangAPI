package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "12345")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, "John Doe", user.Name)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "j@j.com", "12345")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("12345"))
	assert.False(t, user.ValidatePassword("1234"))
	assert.NotEqual(t, "12345", user.Password)
}
