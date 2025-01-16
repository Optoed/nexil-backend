package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка при загрузке .env файла: %v", err)
	}

	r := gin.Default()

	r.GET("/ping")
	r.GET("/pulls")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("не удалось запустить сервер: ", err)
	}
}
