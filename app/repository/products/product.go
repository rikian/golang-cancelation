package products

import (
	"go/cancel/app/config"
	"go/cancel/app/entity"
	"time"

	"github.com/google/uuid"
)

func InitProduct() productContract {
	return &productSchema{}
}

type productContract interface {
	InsertProduct(product *entity.Product) error
}

type productSchema struct {
	Product  *entity.Product
	Products []entity.Product
}

func (p *productSchema) InsertProduct(product *entity.Product) error {
	product.ProductID = uuid.New().String()
	product.ProductCreate = time.Now().Format("20060102150405")
	db := config.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(product).Error; err != nil {
		tx.Rollback()
	}

	return tx.Commit().Error
}
