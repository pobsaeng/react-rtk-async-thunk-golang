package model

import (
	"math/big"
	"time"
)

type Product struct {
	ID          uint64      `json:"id" gorm:"primaryKey;autoIncrement"`
	Code        string      `json:"code" binding:"required"`
	Name        string      `json:"name" binding:"required"`
	Description string      `json:"description"`
	Active      bool        `json:"active"`
	Price       big.Float   `json:"price" binding:"required,gt=0"`
	Stock       int         `json:"stock" binding:"required,gt=0"`
	Weight      big.Float   `json:"weight"`
	Brand       string      `json:"brand"`
	Color       string      `json:"color"`
	Size        string      `json:"size"`
	Length      big.Float   `json:"length"`
	Width       big.Float   `json:"width"`
	Height      big.Float   `json:"height"`
	Image       string      `json:"image"`
	CategoryID  uint64      `json:"category_id"`
	SupplierID  uint64      `json:"supplier_id"`
	CreatedBy   string      `json:"created_by"`
	UpdatedBy   string      `json:"updated_by"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}
