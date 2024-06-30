package models

import "gorm.io/gorm"

type RequestLog struct {
	gorm.Model
	IP       string `json:"ip"`
	Endpoint string `json:"endpoint"`
}
