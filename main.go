package main

import (
	"log"

	"example.com/internal/lyrics/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	g := gin.Default()

	router.RegisterRoutes(g)

	g.Run()
}
