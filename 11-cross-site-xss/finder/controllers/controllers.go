package controllers

import (
	"net/url"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/search", getSomething)
}

func getSomething(c *gin.Context) {
	query := c.Query("q")
	safeQuery := url.QueryEscape(query)              // safety way to receive the user input
	c.String(200, "Your search results: "+safeQuery) //just only this is vulnerable, because is not sanitized and accepts everything that user sends
	//c.JSON(http.StatusAccepted, query)
}
