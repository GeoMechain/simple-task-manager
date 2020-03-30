package main

import (
	"errors"
	"fmt"
	"strings"
)

type Task struct {
	Id               string      `json:"id"`
	ProcessPoints    int         `json:"processPoints"`
	MaxProcessPoints int         `json:"maxProcessPoints"`
	Geometry         [][]float64 `json:"geometry"`
	AssignedUser     string      `json:"assignedUser"`
}

var (
	tasks []*Task
)

func InitTasks() {
	startY := 53.5484
	startX := 9.9714

	tasks = make([]*Task, 0)
	for i := 0; i < 5; i++ {
		geom := make([][]float64, 0)
		geom = append(geom, []float64{startX, startY})
		geom = append(geom, []float64{startX + 0.01, startY})
		geom = append(geom, []float64{startX + 0.01, startY + 0.01})
		geom = append(geom, []float64{startX, startY + 0.01})
		geom = append(geom, []float64{startX, startY})

		startX += 0.01

		tasks = append(tasks, &Task{
			Id:               "t-" + GetId(),
			ProcessPoints:    0,
			MaxProcessPoints: 100,
			Geometry:         geom,
		})
	}

	tasks[0].AssignedUser = "Peter"
	tasks[4].AssignedUser = "Maria"
}

func GetTasks(taskIds []string) []*Task {
	result := make([]*Task, 0)
	for _, t := range tasks {
		for _, i := range taskIds {
			if t.Id == i {
				result = append(result, t)
			}
		}
	}

	return result
}

func GetTask(id string) (*Task, error) {
	for _, t := range tasks {
		if t.Id == id {
			return t, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Task with id '%s' not found", id))
}

// AddTasks sets the ID of the tasks and adds them to the storage.
func AddTasks(newTasks []Task) []*Task {
	result := make([]*Task, 0)

	for _, t := range newTasks {
		t.Id = "t-" + GetId()
		result = append(result, &t)
	}

	tasks = append(tasks, result...)

	return result
}

func AssignUser(id, user string) (*Task, error) {
	task, err := GetTask(id)
	if err == nil {
		if strings.TrimSpace(task.AssignedUser) == "" {
			task.AssignedUser = user
		} else {
			err = errors.New(fmt.Sprintf("User '%s' already assigned, cannot overwrite", task.AssignedUser))
			task = nil
		}
	}

	return task, err
}
