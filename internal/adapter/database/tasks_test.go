package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/fbriansyah/bank-ina-test/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomTask(t *testing.T, userID int32) Task {
	arg := CreateTaskParams{
		UserID:      int32(userID),
		Title:       util.RandomString(8),
		Description: util.RandomString(10),
	}

	task, err := testQueries.CreateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)

	require.Equal(t, arg.UserID, task.UserID)
	require.Equal(t, arg.Title, task.Title)
	require.Equal(t, arg.Description, task.Description)

	return task
}

func TestCreateTask(t *testing.T) {
	user := CreateRandomUser(t, util.RandomEmail(6), "test")
	CreateRandomTask(t, user.ID)
}

func TestGetTaskByID(t *testing.T) {
	user := CreateRandomUser(t, util.RandomEmail(7), "test")
	task1 := CreateRandomTask(t, user.ID)

	task2, err := testQueries.GetTaskByID(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.Title, task2.Title)
	require.Equal(t, task1.Description, task2.Description)
}

func TestUpdateTask(t *testing.T) {
	user := CreateRandomUser(t, util.RandomEmail(8), "test")
	task1 := CreateRandomTask(t, user.ID)

	arg := UpdateTaskParams{
		ID:          task1.ID,
		UserID:      task1.UserID,
		Title:       fmt.Sprintf("%s-updated", task1.Title),
		Description: fmt.Sprintf("%s-updated", task1.Description),
		Status:      "done",
	}

	task2, err := testQueries.UpdateTask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, arg.UserID, task2.UserID)
	require.Equal(t, arg.Title, task2.Title)
	require.Equal(t, arg.Description, task2.Description)
	require.Equal(t, arg.Status, task2.Status)
}

func TestDeleteTask(t *testing.T) {

	user := CreateRandomUser(t, util.RandomEmail(9), "test")
	task1 := CreateRandomTask(t, user.ID)

	err := testQueries.DeleteTask(context.Background(), task1.ID)
	require.NoError(t, err)

	// should return error
	_, err = testQueries.GetTaskByID(context.Background(), task1.ID)
	require.Error(t, err)
}

func TestGetTaskByUser(t *testing.T) {
	user := CreateRandomUser(t, util.RandomEmail(9), "test")

	for i := 0; i < 5; i++ {
		CreateRandomTask(t, user.ID)
	}

	tasks, err := testQueries.GetAllTaskByUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, tasks)

	require.Equal(t, 5, len(tasks))
}
