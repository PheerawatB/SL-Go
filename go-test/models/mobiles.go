package models

import (
	"time"

	"gorm.io/gorm"
)

type Mobile struct {
	gorm.Model
	Name          string
	Price         string
	ProductNumber string
	Details       string
	ProduceDate   time.Time
	GuarunteeYear uint
}

type TradingDetails struct {
	gorm.Model
	Mid           Mobile
	Uid           User
	Sid           Shop
	Discount      int
	TotalPrice    float32
	TatalDiscount float32
	Details       string
	GuarunteeTime time.Time
}
