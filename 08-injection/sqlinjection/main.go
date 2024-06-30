package main

import (
	"log"
	"sqlinjection/controllers"
	"sqlinjection/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=admin password=admin dbname=sqlinjection_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&models.Product{})

	r := gin.Default()

	controllers.RegisterRoutes(r, db)

	r.Run(":3000")
}
