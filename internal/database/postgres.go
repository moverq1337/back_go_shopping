package database

import (
    "fmt"
    "product_api/internal/config"
    "product_api/internal/models"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDB(cfg config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Автоматическая миграция
    err = db.AutoMigrate(&models.Product{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
