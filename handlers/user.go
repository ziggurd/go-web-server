package handlers

import (
    "encoding/json"  // Для работы с JSON
    "net/http"       // HTTP-запросы
    "time"           // Для работы с датой
)

type User struct {  // Структура пользователя
    FirstName   string    `json:"first_name"`     // Имя
    LastName    string    `json:"last_name"`      // Фамилия
    MiddleName  string    `json:"middle_name,omitempty"`  // Отчество (необязательное)
    BirthDate   time.Time `json:"birth_date"`     // Дата рождения
    BirthYear   int       `json:"birth_year"`     // Год рождения (дублирует BirthDate)
    BirthCity   string    `json:"birth_city"`     // Город рождения
}

func UserHandler(w http.ResponseWriter, r *http.Request) {  // Обработчик /user
    user := User{
        FirstName:  "Иван",
        LastName:   "Иванов",
        MiddleName: "Иванович",
        BirthDate:  time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),  // Дата в формате Go
        BirthYear:  1990,
        BirthCity:  "Москва",
    }

    w.Header().Set("Content-Type", "application/json")  // Указываем тип ответа
    w.WriteHeader(http.StatusOK)                       // HTTP 200
    json.NewEncoder(w).Encode(user)                   // Отправляем JSON
}