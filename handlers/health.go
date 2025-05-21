package handlers  // Пакет для обработчиков HTTP-запросов

import (
    "encoding/json"  // Для кодирования JSON
    "net/http"       // Для работы с HTTP
)

type HealthResponse struct {  // Структура ответа /health
    Status  string `json:"status"`   // Статус сервиса (OK/Error)
    Version string `json:"version"`  // Версия приложения
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {  // Обработчик /health
    response := HealthResponse{
        Status:  "OK",          // Сервис работает
        Version: "1.0.0",       // Версия (можно брать из конфига)
    }

    w.Header().Set("Content-Type", "application/json")  // Устанавливает заголовок
    w.WriteHeader(http.StatusOK)                       // HTTP 200
    json.NewEncoder(w).Encode(response)                // Отправляет JSON
}