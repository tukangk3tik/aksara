package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUserRole(t *testing.T) {
	role, err := testQueries.CreateUserRole(context.Background(), "Admin")
	require.NoError(t, err)
	require.NotEmpty(t, role)

	require.Equal(t, role.Name, "Admin")

	testQueries.ClearUserRoles(context.Background())
}

func TestGetUserRoleById(t *testing.T) {
	role, err := testQueries.CreateUserRole(context.Background(), "Admin")
	require.NoError(t, err)
	require.NotEmpty(t, role)

	checkRole, err := testQueries.GetUserRoleById(context.Background(), role.ID)
	require.NoError(t, err)
	require.Equal(t, checkRole.ID, role.ID)
	require.Equal(t, checkRole.Name, role.Name)

	testQueries.ClearUserRoles(context.Background())
}

func TestDeleteUserRole(t *testing.T) {
	role, err := testQueries.CreateUserRole(context.Background(), "Admin")
	require.NoError(t, err)
	require.NotEmpty(t, role)

	testQueries.DeleteUserRole(context.Background(), role.ID)

	checkRole, err := testQueries.GetUserRoleById(context.Background(), role.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Equal(t, checkRole, UserRoles{})

	testQueries.ClearUserRoles(context.Background())
}

func TestListAllUserRoles(t *testing.T) {
	SeedUserRoles(testQueries, context.Background())

	userRoles, err := testQueries.ListAllUserRoles(context.Background(), &ListAllUserRolesParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, len(userRoles), len(UserRolesSeed))
	require.NotEmpty(t, userRoles)

	for _, userRole := range userRoles {
		require.NotEmpty(t, userRole)
	}

	testQueries.ClearUserRoles(context.Background())
}