package models

import (
	"time"
)

// Product Struct
type Product struct {
	SKU         string    `gorm:"type:varchar(50);column:sku;primary_key:true" json:"sku"`
	ProductName string    `gorm:"type:varchar(50);column:product_name" json:"product_name"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
}

func (c Product) TableName() string {
	return "product"
}