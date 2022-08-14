package entity

type Product struct {
	ProductID     string `gorm:"size:128;not null;uniqueIndex;primary_key" json:"product_id"`
	ProductName   string `gorm:"size:128;not null;uniqueIndex" json:"product_name"`
	ProductStock  int    `gorm:"not null" json:"product_stoct"`
	ProductPrice  int    `gorm:"not null" json:"product_price"`
	ProductCreate string `gorm:"size:128;not null" json:"product_create"`
	ProductUpdate string `gorm:"size:128;not null" json:"product_update"`
}
