package config  // Пакет для работы с конфигурацией

import (
    "log"       // Для логирования ошибок
    "os"        // Для чтения переменных окружения
    "github.com/joho/godotenv"  // Для загрузки .env
)

type Config struct {  // Структура для хранения настроек
    Port    string   // Порт сервера
    Version string   // Версия приложения
}

func LoadConfig() *Config {  // Загружает конфиг из .env или переменных окружения
    err := godotenv.Load()  // Пытается загрузить .env
    if err != nil {
        log.Println("No .env file found, using environment variables")  // Если .env нет, использует системные переменные
    }

    return &Config{
        Port:    getEnv("PORT", "8080"),      // Берёт PORT из .env или использует 8080
        Version: getEnv("VERSION", "1.0.0"),  // Аналогично для VERSION
    }
}

func getEnv(key, defaultValue string) string {  // Вспомогательная функция
    value := os.Getenv(key)  // Читает переменную окружения
    if value == "" {
        return defaultValue  // Возвращает значение по умолчанию, если переменной нет
    }
    return value
}