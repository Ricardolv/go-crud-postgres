package main

import (
	"context"
	"github.com/Ricardolv/go-crud-postgres/database"
	"github.com/Ricardolv/go-crud-postgres/handlers"
	"github.com/Ricardolv/go-crud-postgres/repository"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Conectar ao PostgreSQL
	conn, err := database.NewPostgresConnection("postgres://user:password@localhost:5432/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// Inicializar repositório e handlers
	userRepo := repository.NewUserRepository(conn)
	userHandler := handlers.NewUserHandlers(userRepo)

	// Configurar roteador
	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", userHandler.ListUsers).Methods("GET")

	// Iniciar servidor
	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
