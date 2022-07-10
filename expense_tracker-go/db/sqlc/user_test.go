package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var name = "isaiah"

var ctx = context.Background()

func CreateNewUser() (User, error) {
	user, err := testQueries.CreateUser(ctx, name)
	return user, err
}

func TestCreateUser(t *testing.T) {
	user, err := CreateNewUser()
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, name, user.Name)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)
}

func TestGetUser(t *testing.T) {
	user, _ := CreateNewUser()

	require.NotEmpty(t, user)
	user2, err := testQueries.GetUser(ctx, user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user.Name, user2.Name)
	require.Equal(t, user.CreatedAt, user2.CreatedAt)
	require.Exactly(t, user.ID, user2.ID)
	require.Equal(t, user.UpdatedAt, user2.UpdatedAt)
}

func TestUpdateUser(t *testing.T) {
	user, err := CreateNewUser()

	require.NoError(t, err)
	require.NotEmpty(t, user)
	arg := UpdateUserParams{
		ID:   user.ID,
		Name: "james",
		UpdatedAt: time.Now(),
	}

	err = testQueries.UpdateUser(ctx, arg)

	require.NoError(t, err)
}

func TestDelectUser(t *testing.T) {
	user, err := CreateNewUser()

	require.NoError(t, err)
	require.NotEmpty(t, user)

	err = testQueries.DeleteUser(ctx, user.ID)
	require.NoError(t, err)
}

func TestListUser(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateNewUser()
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(ctx, arg)

	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}
