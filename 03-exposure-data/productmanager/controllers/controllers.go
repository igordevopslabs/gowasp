package controllers

import (
	"net/http"
	"productmanager/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterRoutes(r *gin.Engine, database *gorm.DB) {

	db = database

	r.POST("/products", createProduct)
	r.GET("/products", listProducts)
}

func createProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := db.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, product)

}

/*
//Essa função lista todos os produtos e todos os seus campos, o que expoe dados que por regra de negocio sao sensiveis
func listProducts(c *gin.Context) {
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Expoe todos os dados inclusive o dado sensivel CostToPrice
	c.JSON(http.StatusOK, products)
}
*/

//função que lista todos os produtos porem ocultando o campo sensivel

func listProducts(c *gin.Context) {
	var products []models.Product

	//busca no banco por produtos

	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	//trata os dados, criando uma nova estrutura para receber os produtos sem o campo sensivel.

	var publicProducts []models.PublicProduct

	for _, product := range products {
		publicProducts = append(publicProducts, models.PublicProduct{
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}

	c.JSON(http.StatusOK, publicProducts)

}
