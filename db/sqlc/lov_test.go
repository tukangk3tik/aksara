package db

import (
	"context"
	"database/sql"
	"testing"

	"fmt"

	"github.com/stretchr/testify/require"
)

func createOneLov(t *testing.T, iterate int) Lovs {
	arg := CreateLovParams{
		GroupKey:         "TEST_GROUP_KEY_1",
		ParamKey:         fmt.Sprintf("TEST_PARAM_KEY_%d", iterate),
		ParamDescription: fmt.Sprintf("TEST_PARAM_DESCRIPTION_%d", iterate),
		ParentID:         sql.NullInt64{Int64: 1, Valid: true},
	}

	lov, err := testQueries.CreateLov(context.Background(), &arg)
	require.NoError(t, err)

	return lov
}

func TestCreateLov(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	createOneLov(t, 1)
	testQueries.ClearLOVs(context.Background())
}

func TestGetLovByParamKey(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	lov := createOneLov(t, 1)

	checkLov, err := testQueries.GetLovByParamKey(context.Background(), lov.ParamKey)
	require.NoError(t, err)
	require.Equal(t, checkLov.ID, lov.ID)
	require.Equal(t, checkLov.GroupKey, lov.GroupKey)
	require.Equal(t, checkLov.ParamKey, lov.ParamKey)
	require.Equal(t, checkLov.ParamDescription, lov.ParamDescription)
	require.Equal(t, checkLov.ParentID, lov.ParentID)

	testQueries.ClearLOVs(context.Background())
}

func TestGetLovById(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	lov := createOneLov(t, 1)
	checkLov, err := testQueries.GetLovById(context.Background(), lov.ID)
	require.NoError(t, err)
	require.Equal(t, checkLov.ID, lov.ID)
	require.Equal(t, checkLov.GroupKey, lov.GroupKey)
	require.Equal(t, checkLov.ParamKey, lov.ParamKey)
	require.Equal(t, checkLov.ParamDescription, lov.ParamDescription)
	require.Equal(t, checkLov.ParentID, lov.ParentID)

	testQueries.ClearLOVs(context.Background())
}

func TestUpdateLov(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	lov := createOneLov(t, 1)
	arg, err := testQueries.UpdateLov(context.Background(), &UpdateLovParams{
		GroupKey:         "TEST_GROUP_KEY_1",
		ParamKey:         "TEST_PARAM_KEY_99",
		ParamDescription: "TEST_PARAM_DESCRIPTION_99",
		ParentID:         sql.NullInt64{Int64: 1, Valid: true},
		ID:               lov.ID,
	})
	require.NoError(t, err)

	lov.GroupKey = arg.GroupKey
	lov.ParamKey = arg.ParamKey
	lov.ParamDescription = arg.ParamDescription
	lov.ParentID = arg.ParentID

	require.Equal(t, arg.ID, lov.ID)
	require.Equal(t, arg.GroupKey, lov.GroupKey)
	require.Equal(t, arg.ParamKey, lov.ParamKey)
	require.Equal(t, arg.ParamDescription, lov.ParamDescription)
	require.Equal(t, arg.ParentID, lov.ParentID)

	testQueries.ClearLOVs(context.Background())
}

func TestDeleteLov(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	lov := createOneLov(t, 1)
	testQueries.DeleteLov(context.Background(), lov.ID)

	checkLov, err := testQueries.GetLovById(context.Background(), lov.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Equal(t, int(checkLov.ID), 0)
	require.Equal(t, checkLov.GroupKey, "")
	require.Equal(t, checkLov.ParamKey, "")
	require.Equal(t, checkLov.ParamDescription, "")
	require.Equal(t, int(checkLov.ParentID.Int64), 0)

	testQueries.ClearLOVs(context.Background())
}

func TestListAllLovs(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	for i := 1; i <= 10; i++ {
		createOneLov(t, i)
	}

	lovs, err := testQueries.ListAllLovs(context.Background(), &ListAllLovsParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, len(lovs), 10)
	require.NotEmpty(t, lovs)

	for _, lov := range lovs {
		require.NotEmpty(t, lov)
	}

	testQueries.ClearLOVs(context.Background())
}

func TestListLovByGroupKey(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	for i := 1; i <= 10; i++ {
		createOneLov(t, i)
	}

	lovs, err := testQueries.ListLovByGroupKey(context.Background(), &ListLovByGroupKeyParams{
		ParamKey: sql.NullString{
			String: fmt.Sprintf("%%%s%%", "KEY_1"),
			Valid:  true,
		},
		ParamDescription: sql.NullString{
			String: fmt.Sprintf("%%%s%%", "DESCRIPTION_1"),
			Valid:  true,
		},
		GroupKey: "TEST_GROUP_KEY_1",
		Limit:    10,
		Offset:   0,
	})
	require.NoError(t, err)
	require.Equal(t, len(lovs), 2)
	require.NotEmpty(t, lovs)

	for _, lov := range lovs {
		require.NotEmpty(t, lov)
	}

	testQueries.ClearLOVs(context.Background())
}

func TestTotalListAllLovs(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	for i := 1; i <= 10; i++ {
		createOneLov(t, i)
	}

	total, err := testQueries.TotalListAllLovs(context.Background())
	require.NoError(t, err)
	require.Equal(t, total, int64(10))

	testQueries.ClearLOVs(context.Background())
}

func TestTotalListLovByGroupKey(t *testing.T) {
	testQueries.ClearLOVs(context.Background())

	for i := 1; i <= 10; i++ {
		createOneLov(t, i)
	}

	total, err := testQueries.TotalListLovByGroupKey(context.Background(), &TotalListLovByGroupKeyParams{
		GroupKey: "TEST_GROUP_KEY_1",
		ParamKey: sql.NullString{
			String: fmt.Sprintf("%%%s%%", "KEY_1"),
			Valid:  true,
		},
		ParamDescription: sql.NullString{
			String: fmt.Sprintf("%%%s%%", "DESCRIPTION_1"),
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.Equal(t, total, int64(2))

	testQueries.ClearLOVs(context.Background())
}

