package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Price         float64 `json:"price"`
	CostToProduce float64 `json:"cost_to_produce"`
}

type PublicProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
