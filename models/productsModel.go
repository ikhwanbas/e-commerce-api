package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Products struct {
		ID          uuid.UUID `gorm:"primary_key" json:"id"`
		ProductName string    `json:"product_name"`
		Stock       int       `json:"stock"`
		Price       int       `json:"price"`
		Weight      int       `json:"weight"`
		Description string    `json:"description"`
		ImageURL    string    `json:"image_url"`

		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		// migration
		ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID"`
		ProductCategoryID uuid.UUID       `gorm:"index" json:"-"`

		Stores   Stores    `gorm:"foreignKey:StoresID"`
		StoresID uuid.UUID `gorm:"index"`

		CartsID []Carts `json:"-"`
	}

	ProductCategory struct {
		ID           uuid.UUID `gorm:"primary_key" json:"id"`
		CategoryName string    `json:"category_name"`

		// migration
		ProductsID []Products
	}

	ProductCategoryInput struct {
		CategoryName string `json:"category_name"`
	}

	ProductInput struct {
		ProductName       string    `json:"product_name"`
		Stock             int       `json:"stock"`
		Price             int       `json:"price"`
		Weight            int       `json:"weight"`
		Description       string    `json:"description"`
		ImageURL          string    `json:"image_url"`
		ProductCategoryID uuid.UUID `json:"product_category_id"`
	}
)
