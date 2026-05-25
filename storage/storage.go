package storage

import (
	"encoding/json"
	"go_stuff/models"
	"os"
)

const dataFile = "data/todos.json"

func LoadTasks() ([]models.Task, error) {
	data, err := os.ReadFile(dataFile)

	if err != nil {
		if os.IsNotExist(err) {
			return []models.Task{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []models.Task{}, nil
	}

	var tasks []models.Task

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []models.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(dataFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
