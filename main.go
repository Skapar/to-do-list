package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/Skapar/to-do-list/docs"
	"github.com/Skapar/to-do-list/internal/handler"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Todo List API
// @version 1.0
// @description This is a simple Todo List microservice.
// @host localhost:8000
// @BasePath /
func main() {
    r := mux.NewRouter()

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

    handler.RegisterTodoRoutes(r)

	r.HandleFunc("/health", HealthCheck).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", r))
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Router /healthz [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}
