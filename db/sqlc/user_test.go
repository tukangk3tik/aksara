package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T, i int) Users {
	arg := CreateUserParams{
		Name:       fmt.Sprintf("test-%d", i),
		Fullname:   fmt.Sprintf("test-%d", i),
		Email:      fmt.Sprintf("test-%d@mail.com", i),
		Password:   "test",
		UserRoleID: 1,
		OfficeID:   sql.NullInt64{},
		SchoolID:   sql.NullInt64{},
	}

	user, err := testQueries.CreateUser(context.Background(), &arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Fullname, user.Fullname)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.UserRoleID, user.UserRoleID)
	require.Equal(t, arg.OfficeID, user.OfficeID)
	require.Equal(t, arg.SchoolID, user.SchoolID)
	require.NotZero(t, user.ID)

	return user
}

func TestCreateUser(t *testing.T) {
	SeedUserRoles(testQueries, context.Background())

	createRandomUser(t, 1)

	testQueries.ClearUsers(context.Background())
	testQueries.ClearUserRoles(context.Background())
}

func TestGetUserById(t *testing.T) {
	SeedUserRoles(testQueries, context.Background())

	user := createRandomUser(t, 1)
	
	checkUser, err := testQueries.GetUserById(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, checkUser.ID, user.ID)
	require.Equal(t, checkUser.Name, user.Name)
	require.Equal(t, checkUser.Fullname, user.Fullname)
	require.Equal(t, checkUser.Email, user.Email)
	require.Equal(t, checkUser.UserRoleID, user.UserRoleID)
	require.Equal(t, checkUser.OfficeID, user.OfficeID)
	require.Equal(t, checkUser.SchoolID, user.SchoolID)

	testQueries.ClearUsers(context.Background())
	testQueries.ClearUserRoles(context.Background())
}

func TestUpdateUser(t *testing.T) {
	SeedUserRoles(testQueries, context.Background())

	user := createRandomUser(t, 2)

	user.Name = "test update name"
	user.Fullname = "test update fullname"
	_, err := testQueries.UpdateUser(context.Background(), &UpdateUserParams{
		Name:       user.Name,
		Fullname:   user.Fullname,
		Email:      user.Email,
		UserRoleID: user.UserRoleID,
		OfficeID:   user.OfficeID,
		SchoolID:   user.SchoolID,
		ID:         user.ID,
	})
	require.NoError(t, err)
	
	checkUser, err := testQueries.GetUserById(context.Background(), user.ID)
	require.NoError(t, err)
	require.Equal(t, checkUser.ID, user.ID)
	require.Equal(t, checkUser.Name, user.Name)
	require.Equal(t, checkUser.Fullname, user.Fullname)
	
	testQueries.ClearUsers(context.Background())
	testQueries.ClearUserRoles(context.Background())
}

func TestDeleteUser(t *testing.T) {
	SeedUserRoles(testQueries, context.Background())

	user := createRandomUser(t, 3)
	
	testQueries.DeleteUser(context.Background(), user.ID)
	
	checkUser, err := testQueries.GetUserById(context.Background(), user.ID)
	require.Error(t, err)
	require.Equal(t, checkUser, Users{})

	testQueries.ClearUsers(context.Background())
	testQueries.ClearUserRoles(context.Background())
}

func TestTotalListAllUsers(t *testing.T) {
	SeedUserRoles(testQueries, context.Background())

	for i := range 10 {
		createRandomUser(t, i)
	}

	total, err := testQueries.TotalListAllUsers(context.Background())
	require.NoError(t, err)
	require.Equal(t, total, int64(10))

	testQueries.ClearUsers(context.Background())
	testQueries.ClearUserRoles(context.Background())
}

func TestListAllUsers(t *testing.T) {
	SeedUserRoles(testQueries, context.Background())

	for i := range 10 {
		createRandomUser(t, i)
	}

	users, err := testQueries.ListAllUsers(context.Background(), &ListAllUsersParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, len(users), 10)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
	}

	testQueries.ClearUsers(context.Background())
	testQueries.ClearUserRoles(context.Background())
}