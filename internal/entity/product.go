package entity

import (
	"errors"
	"time"

	"github.com/misael28/GoApi/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInvalidID       = errors.New("invalid id")
	ErrInvalidPrice    = errors.New("invalid price")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID((p.ID.String())); err != nil {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewId(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}
	return product, nil
}
