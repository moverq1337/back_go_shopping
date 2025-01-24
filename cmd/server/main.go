package main

import (
	"log"
	"product_api/internal/config"
	"product_api/internal/database"
	"product_api/internal/handlers"
	"product_api/internal/repository"
	"product_api/internal/router"
)

func main() {
    cfg := config.LoadConfig()

    db, err := database.ConnectDB(cfg)
    if err != nil {
        log.Fatalf("Не удалось подключиться к базе данных: %v", err)
    }

    // Создаем репозитории
    productRepo := repository.NewProductRepository(db)
    userRepo := repository.NewUserRepository(db)
    cartRepo := repository.NewCartRepository(db)

    productHandler := handlers.NewProductHandler(productRepo)
    authHandler := handlers.NewAuthHandler(userRepo)
    cartHandler := handlers.NewCartHandler(cartRepo)

    // Настройка маршрутов
    r := router.SetupRouter(productHandler, authHandler, cartHandler)

    // Запуск сервера
    if err := r.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Не удалось запустить сервер: %v", err)
    }
}
