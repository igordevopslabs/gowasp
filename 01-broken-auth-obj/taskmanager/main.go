package main

import (
	"log"
	"taskmanager/controllers"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	//inicia a conexão com o banco de dados.
	dsn := "host=localhost user=admin password=admin dbname=task_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error to connect DB:", err)
	}

	db.Migrator().DropTable(&models.User{}, &models.Task{})
	db.AutoMigrate(&models.User{}, &models.Task{})

	//Inicia o API server
	r := gin.Default()
	controllers.RegisterRoutes(r, db)

	r.Run(":3000")
}
