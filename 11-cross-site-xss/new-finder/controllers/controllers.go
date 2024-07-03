package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/finder", finder)
}

func finder(c *gin.Context) {
	input := c.Query("input")
	p := bluemonday.UGCPolicy()
	sanitizedEntry := p.Sanitize(input)
	c.JSON(http.StatusAccepted, sanitizedEntry)
}
