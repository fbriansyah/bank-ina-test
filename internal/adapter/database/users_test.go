package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/fbriansyah/bank-ina-test/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T, email, password string) User {

	hash, err := util.HashPassword(password)
	require.NoError(t, err)

	arg := CreateUserParams{
		Name:     util.RandomName(),
		Email:    email,
		Password: hash,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)

	err = util.CheckPassword(password, user.Password)
	require.NoError(t, err)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t, util.RandomEmail(5), "secret")
}

func TestGetUserByID(t *testing.T) {
	email := util.RandomEmail(5)
	user1 := CreateRandomUser(t, email, "secret")

	user2, err := testQueries.GetUserByID(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, email, user2.Email)
}

func TestUpdateUser(t *testing.T) {
	email := util.RandomEmail(6)

	user1 := CreateRandomUser(t, email, "secret")

	arg := UpdateUserParams{
		Email:    email,
		Name:     fmt.Sprintf("%s-updated", user1.Name),
		Password: user1.Password,
		ID:       user1.ID,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, true, user2.UpdatedAt.Valid) // check if update at not null
}

func TestDeleteUser(t *testing.T) {

	email := util.RandomEmail(7)

	user := CreateRandomUser(t, email, "secret")
	err := testQueries.DeleteUser(context.Background(), user.ID)
	require.NoError(t, err)

	// should return error
	_, err = testQueries.GetUserByID(context.Background(), user.ID)
	require.Error(t, err)
}

func TestGetUserByEmail(t *testing.T) {

	email := util.RandomEmail(8)

	user1 := CreateRandomUser(t, email, "secret")

	user2, err := testQueries.GetUserByEmail(context.Background(), email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
}
