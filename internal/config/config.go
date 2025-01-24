package config

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DBHost     string
    DBPort     string
    DBUser     string
    DBPassword string
    DBName     string
    DBSSLMode  string
    ServerPort string
}

func LoadConfig() Config {
    // Загрузка .env файла
    err := godotenv.Load()
    if err != nil {
        log.Println("Файл .env не найден, используется конфигурация по умолчанию")
    }

    return Config{
        DBHost:     getEnv("DB_HOST", "localhost"),
        DBPort:     getEnv("DB_PORT", "5432"),
        DBUser:     getEnv("DB_USER", "postgres"),
        DBPassword: getEnv("DB_PASSWORD", "password"),
        DBName:     getEnv("DB_NAME", "products_db"),
        DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
        ServerPort: getEnv("SERVER_PORT", "8080"),
    }
}

func getEnv(key, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultVal
}
