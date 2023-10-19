package port

import (
	dmsession "github.com/fbriansyah/bank-ina-test/internal/application/domain/session"
	dmtask "github.com/fbriansyah/bank-ina-test/internal/application/domain/task"
	dmtoken "github.com/fbriansyah/bank-ina-test/internal/application/domain/token"
	dmuser "github.com/fbriansyah/bank-ina-test/internal/application/domain/user"
)

type ServicePort interface {
	Login(email, password string) (dmsession.Session, error)
	CheckToken(token string) (*dmtoken.Payload, error)
	// Register new user
	Register(email, password, name string) (dmuser.User, error)
	// ListUsers get all user
	ListUsers() ([]dmuser.User, error)
	// GetUserByID get data user by user ID
	GetUserByID(id int32) (dmuser.User, error)
	UpdateUser(id int32, user dmuser.User) (dmuser.User, error)
	DeleteUser(id int32) error

	CreateTask(userID int32, title, description string) (dmtask.Task, error)
	ListTasks(userID int32) ([]dmtask.Task, error)
	GetTaskByID(id int32) (dmtask.Task, error)
	UpdateTask(task dmtask.Task) (dmtask.Task, error)
	DeleteTasks(id int32) error
}
