package main

import (
	get "testGin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/get", get.NameQuery)
	r.Run("8088") // listen and serve on 0.0.0.0:8080
}
