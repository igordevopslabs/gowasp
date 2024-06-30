package controllers

import (
	"net/http"
	"ratelimiting/models"

	"github.com/didip/tollbooth/v6"
	"github.com/didip/tollbooth/v6/limiter"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func RegisterRoutes(r *gin.Engine, database *gorm.DB) {
	db = database

	//limita a 10 req por segundo
	limiter := tollbooth.NewLimiter(10, nil)

	r.GET("/ping", limitMiddleware(limiter), handlePing)

}

func handlePing(c *gin.Context) {
	ip := c.ClientIP()
	endpoint := "/ping"

	log := models.RequestLog{
		IP:       ip,
		Endpoint: endpoint,
	}

	if err := db.Create(&log).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "pong!"})
}

//Necessário implementar algum middleware para service como ratelimit dessa rota.
//do contrario, podemos ter um DoS, seja por um atacante ou até mesmo fogo amigo.

func limitMiddleware(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if httpError != nil {
			c.Header("Content-Type", "application/json")
			c.AbortWithStatusJSON(httpError.StatusCode, gin.H{"error": httpError.Message})
			return
		}
		c.Next()
	}
}
