package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetUsers(t *testing.T) {
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUsers)
	handler.ServeHTTP(rr, req)

	// Проверяем, что код ответа 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler вернул неверный статус-код: получили %v ожидалось %v", status, http.StatusOK)
	}

	// Проверяем, что возвращенный JSON корректен
	expected := `[{"id":1,"name":"John Doe","age":30}]`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler вернул неожиданный результат: получили %v ожидалось %v", rr.Body.String(), expected)
	}
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getUser)
	handler.ServeHTTP(rr, req)

	// Проверяем статус ответа
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler вернул неверный статус-код: получили %v ожидалось %v", status, http.StatusOK)
	}

	// Проверка на правильный JSON
	expected := `{"id":1,"name":"John Doe","age":30}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler вернул неожиданный результат: получили %v ожидалось %v", rr.Body.String(), expected)
	}
}

func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`{"name":"Alice","age":25}`)
	req, err := http.NewRequest("POST", "/users", strings.NewReader(string(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(createUser)
	handler.ServeHTTP(rr, req)

	// Проверяем статус 201
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler вернул неверный статус-код: получили %v ожидалось %v", status, http.StatusCreated)
	}

	// Проверка правильности возвращенного JSON
	expected := `{"id":2,"name":"Alice","age":25}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler вернул неожиданный результат: получили %v ожидалось %v", rr.Body.String(), expected)
	}
}

func TestUpdateUser(t *testing.T) {
	var jsonStr = []byte(`{"name":"Alice Updated","age":26}`)
	req, err := http.NewRequest("PUT", "/users/1", strings.NewReader(string(jsonStr)))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateUser)
	handler.ServeHTTP(rr, req)

	// Проверка, что статус ответа - 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler вернул неверный статус-код: получили %v ожидалось %v", status, http.StatusOK)
	}

	// Проверка возвращенного JSON
	expected := `{"id":1,"name":"Alice Updated","age":26}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler вернул неожиданный результат: получили %v ожидалось %v", rr.Body.String(), expected)
	}
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(deleteUser)
	handler.ServeHTTP(rr, req)

	// Проверка статуса 204 No Content
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler вернул неверный статус-код: получили %v ожидалось %v", status, http.StatusNoContent)
	}
}
