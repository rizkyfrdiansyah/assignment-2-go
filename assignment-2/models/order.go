package models

import "time"

type Order struct {
	ID           uint `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
	Items []Item
}

type Item struct {
	ID uint `gprm:"primaryKey"`
	Code string
	Description string
	Quantity uint
	OrderID uint
}