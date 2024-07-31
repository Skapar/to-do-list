package app

import (
	"github.com/Skapar/to-do-list/internal/handler"
	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
    r := mux.NewRouter()
    handler.RegisterTodoRoutes(r)
    return r
}