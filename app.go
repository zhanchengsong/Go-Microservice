package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zhanchengsong/Go-Microservice.git/handlers"
)

func main() {
	fmt.Println("Test module file!")
	r := gin.Default()
	r.GET("/ping", handlers.Ping)
	r.Run()
}
