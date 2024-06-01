package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var taskCounter int

//var TIME_ADDITION_MS int
//var TIME_SUBTRACTION_MS int
//var TIME_MULTIPLICATION_MS int
//var TIME_DIVISION_MS int

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to our API!")
	})
	//Если в браузере: http://localhost:8080/api/v1/expressions
	// в командной строке: curl http://127.0.0.1:8080/api/v1/expressions
	r.HandleFunc("/api/v1/expressions", getExpressions).Methods("GET")

	//Если в браузере: http://localhost:8080/api/v1/expressions/:1
	// в командной строке: curl http://127.0.0.1:8080/api/v1/expressions/:1

	r.HandleFunc("/api/v1/expressions/:{id}", getExpression).Methods("GET")
	r.HandleFunc("/internal/task", getTask).Methods("GET")
	r.HandleFunc("/internal/task", setResult).Methods("POST")

	//Если в браузере: http://localhost:8080/api/v1/tasks
	// в командной строке: curl http://127.0.0.1:8080/api/v1/tasks
	r.HandleFunc("/api/v1/tasks", getTasks).Methods("GET")
	//r.HandleFunc("/api/v1/tasks/:{id}", getExpression).Methods("GET")

	//Если в браузере: http://localhost:8080/api/v1/datas
	// в командной строке: curl http://127.0.0.1:8080/api/v1/datas
	r.HandleFunc("/api/v1/datas", getDatas).Methods("GET")
	//r.HandleFunc("/api/v1/data/:{id}", getTasks).Methods("GET")
	http.ListenAndServe(":8080", r)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve a task for execution
	// Based on TIME_ADDITION_MS, TIME_SUBTRACTION_MS, TIME_MULTIPLICATION_MS, TIME_DIVISION_MS
	// Return a task in the specified format
}

func setResult(w http.ResponseWriter, r *http.Request) {
	var user Expression
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func addExpression(w http.ResponseWriter, r *http.Request) {
	var user Expression
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskCounter++
	user.ID = taskCounter
	response := map[string]int{"id": user.ID}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

////////////////////////////////

func getExpressions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// в пакете package main есть ещё файл - models.go, там и смотрим
	// значение переменной expressions
	//Берем оттуда значения, декодируем и помещаем в w- ответ
	json.NewEncoder(w).Encode(expressions)
}
func getDatas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// в пакете package main есть ещё файл - models.go, там и смотрим
	// значение переменной expressions
	//Берем оттуда значения, декодируем и помещаем в w- ответ
	json.NewEncoder(w).Encode(data)
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// в пакете package main есть ещё файл - models.go, там и смотрим
	// значение переменной expressions
	//Берем оттуда значения, декодируем и помещаем в w- ответ
	json.NewEncoder(w).Encode(task)
}
func getExpression(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	for _, expression := range expressions {
		if expression.ID == id {
			json.NewEncoder(w).Encode(expression)
			return
		}
	}
	json.NewEncoder(w).Encode(&Expression{})
}
