package main

import (
	"github.com/gin-gonic/gin"
	"testGin/router"
)

func main() {
	r := gin.Default()
	r.GET("/get", get.NameQuery)
	r.Run() // listen and serve on 0.0.0.0:8080
}