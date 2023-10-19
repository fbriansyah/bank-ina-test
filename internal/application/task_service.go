package application

import (
	"context"

	db "github.com/fbriansyah/bank-ina-test/internal/adapter/database"
	dmtask "github.com/fbriansyah/bank-ina-test/internal/application/domain/task"
)

func (s *Service) CreateTask(userID int32, title, description string) (dmtask.Task, error) {
	task, err := s.db.CreateTask(context.Background(), db.CreateTaskParams{
		UserID:      userID,
		Title:       title,
		Description: description,
	})

	if err != nil {
		return dmtask.Task{}, err
	}

	return dmtask.Task{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
	}, nil
}
func (s *Service) ListTasks(userID int32) ([]dmtask.Task, error) {
	tasks, err := s.db.GetAllTaskByUser(context.Background(), userID)
	if err != nil {
		return []dmtask.Task{}, err
	}

	var listTask []dmtask.Task
	for _, task := range tasks {
		listTask = append(listTask, dmtask.Task{
			ID:          task.ID,
			UserID:      task.UserID,
			Title:       task.Title,
			Description: task.Description,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt,
		})
	}

	return listTask, nil
}
func (s *Service) GetTaskByID(id int32) (dmtask.Task, error) {
	task, err := s.db.GetTaskByID(context.Background(), id)
	if err != nil {
		return dmtask.Task{}, err
	}
	return dmtask.Task{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
	}, nil
}
func (s *Service) UpdateTask(tsk dmtask.Task) (dmtask.Task, error) {
	task, err := s.db.UpdateTask(context.Background(), db.UpdateTaskParams{
		UserID:      tsk.UserID,
		Title:       tsk.Title,
		Description: tsk.Description,
		Status:      tsk.Status,
		ID:          tsk.ID,
	})

	if err != nil {
		return dmtask.Task{}, err
	}

	return dmtask.Task{
		ID:          task.ID,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
	}, nil
}
func (s *Service) DeleteTasks(id int32) error {
	return s.db.DeleteTask(context.Background(), id)
}
