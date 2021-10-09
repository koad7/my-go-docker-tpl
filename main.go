package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// fs := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fs)
	var router *gin.Engine
	router = gin.Default()
	router.Static("/", "./static")
	router.Run()
}
