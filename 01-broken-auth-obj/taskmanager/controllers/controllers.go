package controllers

import (
	"net/http"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

// registra as rotas existentes
func RegisterRoutes(r *gin.Engine, database *gorm.DB) {
	db = database

	r.POST("/users", createUser)
	r.POST("/tasks", createTask)
	r.GET("tasks/:id", getTask)
}

func createUser(c *gin.Context) {
	var user models.User

	//recebe a req
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//cria user no banco
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func createTask(c *gin.Context) {
	username, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.UserID = user.ID

	if err := db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func getTask(c *gin.Context) {

	username, password, hasAuth := c.Request.BasicAuth()
	if !hasAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := db.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var task models.Task
	id := c.Param("id")

	if err := db.Find(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found for this id", "id": id})
		return
	}

	/*	if task.UserID != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have access to this task"})
		return
	}*/

	c.JSON(http.StatusOK, task)
}
