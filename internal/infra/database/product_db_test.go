package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/misael28/GoApi/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{}) // no cache here for the tests
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate((&entity.Product{}))
	
	product, err := entity.NewProduct("Product 1", 10.00)
	assert.NoError(t, err)
	
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)

}

func TestFindAllProducts(t *testing.T){
	db, err := gorm.Open(sqlite.Open("file:memory:"), &gorm.Config{}) // no cache here for the tests
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate((&entity.Product{}))
	
	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}
	
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)
}
