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
    } else {
        log.Println("Успешно подключено к базе данных")
    }

    productRepo := repository.NewProductRepository(db)
    productHandler := handlers.NewProductHandler(productRepo)

    r := router.SetupRouter(productHandler)

    if err := r.Run(":" + cfg.ServerPort); err != nil {
        log.Fatalf("Не удалось запустить сервер: %v", err)
    }
}
