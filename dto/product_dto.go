package dto

type ProductDto struct {
	SKU         string `json:"sku" form:"sku" param:"sku" validate:"required"`
	ProductName string `json:"product_name" form:"product_name" param:"product_name" validate:"required"`
}

type ProductEditDto struct {
	SKU         string `json:"sku" form:"sku" param:"sku" validate:"required"`
	ProductName string `json:"product_name" form:"product_name" param:"product_name" validate:"required"`
}
