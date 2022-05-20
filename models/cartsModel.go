package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	Carts struct {
		ID            uuid.UUID `gorm:"primary_key" json:"id"`
		PaymentStatus string    `json:"payment_status"`
		Quantity      int       `json:"quantity"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`

		ProductsID uuid.UUID `gorm:"index"`
		Products   Products  `gorm:"foreignKey:ProductsID"`

		UsersID uuid.UUID `gorm:"index"`
		Users   Users     `gorm:"foreignKey:UsersID"`
	}
	CartsInput struct {
		PaymentStatus string    `json:"payment_status"`
		Quantity      int       `json:"quantity"`
		ProductID     uuid.UUID `gorm:"index" json:"product_id"`
	}
)
