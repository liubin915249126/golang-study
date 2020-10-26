package main

import "github.com/gin-gonic/gin"
import "gin/router"

func main() {
	r := gin.Default()
	r.GET("/ping", get.Get)
	r.Run() // listen and serve on 0.0.0.0:8080
}