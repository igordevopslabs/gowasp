package controllers

import (
	"net/http"
	"sqlinjection/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterRoutes(r *gin.Engine, database *gorm.DB) {
	db = database

	r.POST("/products", createProduct)
	r.GET("/products/search", searchProduct)
}

func createProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

/*
func searchProduct(c *gin.Context) {
	name := c.Query("name")
	var products []models.Product
	//essa query esta bem problematica pois esta concatenando dados diretamente, ou seja
	//caso explorado irá executar diretamente o comando no banco
	query := "SELECT * FROM products WHERE name LIKE '%" + name + "%'"
	if err := db.Raw(query).Scan(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
*/

// função usando os parametros de consulta corretamente ao inves de concatenar
func searchProduct(c *gin.Context) {
	name := c.Query("name")
	var products []models.Product
	if err := db.Where("name LIKE ?", "%"+name+"%").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
