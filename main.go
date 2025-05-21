package main  // Главный пакет (точка входа)

import (
    "log"       // Логирование
    "net/http"  // HTTP-сервер
    "github.com/ziggurd/go-web-server/config"   // Конфигурация
    "github.com/ziggurd/go-web-server/handlers" // Обработчики
    "github.com/gorilla/mux"                     // Роутер
)

func main() {
    cfg := config.LoadConfig()  // Загружает конфиг

    r := mux.NewRouter()  // Создаёт роутер (для маршрутизации)
    r.HandleFunc("/health", handlers.HealthHandler).Methods("GET")  // /health → HealthHandler
    r.HandleFunc("/user", handlers.UserHandler).Methods("GET")      // /user → UserHandler

    log.Printf("Server started on :%s\n", cfg.Port)  // Лог порта
    log.Fatal(http.ListenAndServe(":"+cfg.Port, r))  // Запуск сервера
}