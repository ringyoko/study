package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB
var validate *validator.Validate

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age" validate:"gte=0"`
}

// Подключение к базе данных
func initDB() {
	var err error
	connStr := "user=alice password=#### dbname=godatabase sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Ошибка проверки подключения к базе данных:", err)
	}

	log.Println("Успешно подключено к базе данных!")
}

// Получение списка пользователей с пагинацией и фильтрацией
func getUsers(w http.ResponseWriter, r *http.Request) {
	limit := 10
	offset := 0
	nameFilter := r.URL.Query().Get("name")

	// Получение параметров пагинации
	if l, err := strconv.Atoi(r.URL.Query().Get("limit")); err == nil {
		limit = l
	}
	if o, err := strconv.Atoi(r.URL.Query().Get("offset")); err == nil {
		offset = o
	}

	// SQL запрос с фильтрацией
	query := "SELECT id, name, age FROM users WHERE 1=1"
	args := []interface{}{}
	if nameFilter != "" {
		query += " AND name ILIKE $1"
		args = append(args, "%"+nameFilter+"%")
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		log.Println("Ошибка выполнения запроса к базе данных:", err)
		http.Error(w, "Ошибка выполнения запроса", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			log.Println("Ошибка сканирования результата:", err)
			http.Error(w, "Ошибка чтения данных", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	// Проверка на ошибки после завершения цикла rows.Next
	if err := rows.Err(); err != nil {
		log.Println("Ошибка обработки данных:", err)
		http.Error(w, "Ошибка обработки данных", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Получение информации о конкретном пользователе
func getUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var user User

	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Age)
	if err == sql.ErrNoRows {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	} else if err != nil {
		log.Println("Ошибка выполнения запроса к базе данных:", err)
		http.Error(w, "Ошибка получения пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Добавление нового пользователя
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		http.Error(w, "Некорректные данные", http.StatusBadRequest)
		return
	}

	// Валидация данных
	if err := validate.Struct(newUser); err != nil {
		log.Println("Ошибка валидации данных:", err)
		http.Error(w, "Ошибка валидации: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := db.QueryRow("INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id", newUser.Name, newUser.Age).Scan(&newUser.ID)
	if err != nil {
		log.Println("Ошибка добавления пользователя в базу данных:", err)
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// Обновление информации о пользователе
func updateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updateUser User

	if err := json.NewDecoder(r.Body).Decode(&updateUser); err != nil {
		log.Println("Ошибка декодирования JSON:", err)
		http.Error(w, "Некорректные данные", http.StatusBadRequest)
		return
	}

	// Валидация данных
	if err := validate.Struct(updateUser); err != nil {
		log.Println("Ошибка валидации данных:", err)
		http.Error(w, "Ошибка валидации: "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err := db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", updateUser.Name, updateUser.Age, id)
	if err != nil {
		log.Println("Ошибка обновления пользователя в базе данных:", err)
		http.Error(w, "Ошибка обновления пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateUser)
}

// Удаление пользователя
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Println("Ошибка удаления пользователя из базы данных:", err)
		http.Error(w, "Ошибка удаления пользователя", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {

	validate = validator.New()

	initDB()

	// Настройка маршрутов
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUser).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	log.Println("Сервер запущен на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
