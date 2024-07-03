package main

import (
	"finder/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	controllers.RegisterRoutes(r)

	r.Run()
}
