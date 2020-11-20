package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/longndb/simple-rest-go/pkg/handler"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/todo", handler.GetAllTodo).Methods(http.MethodGet)
	r.HandleFunc("/api/todo/{id}", handler.GetTodoById).Methods(http.MethodGet)
	r.HandleFunc("/api/todo", handler.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/api/todo/{id}", handler.UpdateTodo).Methods(http.MethodPut)
	r.HandleFunc("/api/todo/{id}", handler.DeleteTodo).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8080", r))
}
