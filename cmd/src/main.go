package main

import (
	"errors"
	"example.com/mod/internal/handlers"
	"example.com/mod/internal/repositories"
	"example.com/mod/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка при загрузке .env файла: %v", err)
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Пинг базы данных
	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка при пинге базы данных:", err)
	} else {
		fmt.Println("Подключение к базе данных успешно!")
	}
	defer db.Close()

	// Настройка миграций
	m, err := migrate.New(
		"file://migrations", // Путь к папке с миграциями
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatal("Error initializing migration:", err)
	}

	// Применение миграций
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal("Failed to apply migrations:", err)
	}

	procurementRepo := repositories.NewProcurementRepository(db)
	procurementService := services.NewProcurementService(procurementRepo)
	procurementHandler := handlers.NewProcurementHandler(procurementService)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/procurements", procurementHandler.GetProcurements)
	}

	if err := r.Run(":8081"); err != nil {
		log.Fatal("не удалось запустить сервер: ", err)
	}
}
