package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClearUserRoles(t *testing.T) {
	SeedUserRoles(testQueries, context.Background())

	testQueries.ClearUserRoles(context.Background())
	userRoles, err := testQueries.ListAllUserRoles(context.Background(), &ListAllUserRolesParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, len(userRoles), 0)
}