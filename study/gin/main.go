package main

import "github.com/gin-gonic/gin"
import "get"

func main() {
	r := gin.Default()
	r.GET("/ping", get.Main )
	r.Run() // listen and serve on 0.0.0.0:8080
}