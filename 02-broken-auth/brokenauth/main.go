package main

import (
	"brokenauth/controllers"
	"brokenauth/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	//conexao com o banco
	dsn := "host=localhost user=admin password=admin dbname=brokenauth_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect db", err)
	}

	db.AutoMigrate(&models.User{}, &models.Token{})

	r := gin.Default()

	controllers.RegisterRoutes(r, db)

	r.Run(":3000")

}
