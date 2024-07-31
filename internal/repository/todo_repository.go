package repository

import (
	"errors"

	"github.com/Skapar/to-do-list/internal/model"
)

var tasks = []model.Task{}

func CreateTask(task model.Task) {
    tasks = append(tasks, task)
}

func UpdateTask(id string, updatedTask model.Task) error {
    for index, item := range tasks {
        if item.ID == id {
            tasks[index] = updatedTask
            return nil
        }
    }
    return errors.New("task not found")
}

func DeleteTask(id string) error {
    for index, item := range tasks {
        if item.ID == id {
            tasks = append(tasks[:index], tasks[index+1:]...)
            return nil
        }
    }
    return errors.New("task not found")
}

func MarkTaskDone(id string) error {
    for index, item := range tasks {
        if item.ID == id {
            tasks[index].Status = "done"
            return nil
        }
    }
    return errors.New("task not found")
}


func ListTasks(status string) []model.Task {
    var filteredTasks []model.Task
    for _, task := range tasks {
        if task.Status == status {
            filteredTasks = append(filteredTasks, task)
        }
    }
    return filteredTasks
}
