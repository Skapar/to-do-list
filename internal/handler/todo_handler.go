package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Skapar/to-do-list/internal/model"
	"github.com/Skapar/to-do-list/internal/repository"
	"github.com/Skapar/to-do-list/internal/service"
)

func RegisterTodoRoutes(r *mux.Router) {
    r.HandleFunc("/api/todo-list/tasks", CreateTask).Methods("POST")
    r.HandleFunc("/api/todo-list/tasks/{id}", UpdateTask).Methods("PUT")
    r.HandleFunc("/api/todo-list/tasks/{id}", DeleteTask).Methods("DELETE")
    r.HandleFunc("/api/todo-list/tasks/{id}/done", MarkTaskDone).Methods("PUT")
    r.HandleFunc("/api/todo-list/tasks", ListTasks).Methods("GET")
}

// CreateTask godoc
// @Summary Create a new task
// @Description Create a new task with the input payload
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body model.Task true "Create Task"
// @Success 201 {object} model.Task
// @Failure 400 {object} map[string]string
// @Router /api/todo-list/tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
    var task model.Task
    _ = json.NewDecoder(r.Body).Decode(&task)
    task.ID = service.GenerateID()
    task.Status = "active"
    repository.CreateTask(task)
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(task)
}

// UpdateTask godoc
// @Summary Update an existing task
// @Description Update an existing task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID"
// @Param task body model.Task true "Update Task"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /api/todo-list/tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var updatedTask model.Task
    _ = json.NewDecoder(r.Body).Decode(&updatedTask)
    updatedTask.ID = params["id"]
    if err := repository.UpdateTask(params["id"], updatedTask); err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags tasks
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /api/todo-list/tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    if err := repository.DeleteTask(params["id"]); err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// MarkTaskDone godoc
// @Summary Mark a task as done
// @Description Mark a task as done by ID
// @Tags tasks
// @Param id path string true "Task ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Router /api/todo-list/tasks/{id}/done [put]
func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    if err := repository.MarkTaskDone(params["id"]); err != nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.WriteHeader(http.StatusNoContent)
}

// ListTasks godoc
// @Summary List tasks by status
// @Description List tasks by status (active or done)
// @Tags tasks
// @Param status query string false "Task Status"
// @Success 200 {array} model.Task
// @Router /api/todo-list/tasks [get]
func ListTasks(w http.ResponseWriter, r *http.Request) {
    status := r.URL.Query().Get("status")
    if status == "" {
        status = "active"
    }
    tasks := repository.ListTasks(status)
    json.NewEncoder(w).Encode(tasks)
}
