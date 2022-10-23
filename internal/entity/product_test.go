package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product Test", 12)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.Name)
	assert.NotEmpty(t, p.Price)
	assert.Equal(t, "Product Test", p.Name)
	assert.Equal(t, 12.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T){
	p, err := NewProduct("",10)
	assert.Nil(t,p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenProductIsRequired(t *testing.T){
	p, err := NewProduct("Product Test", 0)
	assert.Nil(t,p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T){
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
