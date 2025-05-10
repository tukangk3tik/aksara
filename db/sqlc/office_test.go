package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const districtID = 537104

func createOneOffice(t *testing.T, iterate int) Offices {
	arg := CreateOfficeParams{
		Code:       fmt.Sprintf("test-%d", iterate),
		Name:       fmt.Sprintf("test-%d", iterate),
		ProvinceID: provinceID,
		RegencyID:  regencyID,
		DistrictID: districtID,
		Email:      fmt.Sprintf("test-%d@mail.com", iterate),
		Phone:      sql.NullString{String: "+628123456789", Valid: true},
		Address:    sql.NullString{String: "test address", Valid: true},
		LogoUrl:    sql.NullString{String: "test logo url", Valid: true},
		CreatedBy:  1,
	}

	office, err := testQueries.CreateOffice(context.Background(), &arg)
	require.NoError(t, err)

	require.Equal(t, arg.Code, office.Code)
	require.Equal(t, arg.Name, office.Name)
	require.Equal(t, arg.ProvinceID, office.ProvinceID)
	require.Equal(t, arg.RegencyID, office.RegencyID)
	require.Equal(t, arg.DistrictID, office.DistrictID)
	require.Equal(t, arg.Email, office.Email)
	require.Equal(t, arg.Phone, office.Phone)
	require.Equal(t, arg.Address, office.Address)
	require.Equal(t, arg.LogoUrl, office.LogoUrl)
	require.Equal(t, arg.CreatedBy, office.CreatedBy)

	return office
}

func TestCreateOffice(t *testing.T) {
	createOneOffice(t, 1)
	testQueries.ClearOffices(context.Background())
}

func TestGetOfficeById(t *testing.T) {
	office := createOneOffice(t, 1)

	checkOffice, err := testQueries.GetOfficeById(context.Background(), office.ID)
	require.NoError(t, err)
	require.Equal(t, checkOffice.ID, office.ID)
	require.Equal(t, checkOffice.Code, office.Code)
	require.Equal(t, checkOffice.Name, office.Name)
	require.Equal(t, checkOffice.ProvinceID, office.ProvinceID)
	require.Equal(t, checkOffice.RegencyID, office.RegencyID)
	require.Equal(t, checkOffice.DistrictID, office.DistrictID)
	require.Equal(t, checkOffice.Email, office.Email)
	require.Equal(t, checkOffice.Phone, office.Phone)
	require.Equal(t, checkOffice.Address, office.Address)
	require.Equal(t, checkOffice.LogoUrl, office.LogoUrl)
	require.Equal(t, checkOffice.CreatedBy, office.CreatedBy)

	testQueries.ClearOffices(context.Background())
}

func TestUpdateOffice(t *testing.T) {
	office := createOneOffice(t, 1)
	
	arg, err := testQueries.UpdateOffice(context.Background(), &UpdateOfficeParams{
		Code:       "test update code",
		Name:       "test update name",
		ProvinceID: provinceID,
		RegencyID:  regencyID,
		DistrictID: districtID,
		Email:      "test update email",
		ID:         office.ID,
	})

	office.Code = arg.Code
	office.Name = arg.Name
	office.Email = arg.Email

	require.NoError(t, err)
	require.Equal(t, arg.ID, office.ID)
	require.Equal(t, arg.Code, office.Code)
	require.Equal(t, arg.Name, office.Name)
	require.Equal(t, arg.ProvinceID, office.ProvinceID)
	require.Equal(t, arg.RegencyID, office.RegencyID)
	require.Equal(t, arg.DistrictID, office.DistrictID)
	require.Equal(t, arg.Email, office.Email)

	testQueries.ClearOffices(context.Background())
}

func TestDeleteOffice(t *testing.T) {
	office := createOneOffice(t, 1)
	
	testQueries.DeleteOffice(context.Background(), office.ID)
	
	checkOffice, err := testQueries.GetOfficeById(context.Background(), office.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Equal(t, int(checkOffice.ID), 0)
	require.Equal(t, checkOffice.Name, "")
	require.Equal(t, checkOffice.Code, "")
	require.Equal(t, int(checkOffice.ProvinceID), 0)
	require.Equal(t, int(checkOffice.RegencyID), 0)
	require.Equal(t, int(checkOffice.DistrictID), 0)
	require.Equal(t, checkOffice.Email, "")
	
	testQueries.ClearOffices(context.Background())
}

func TestListAllOffices(t *testing.T) {
	for i := 1; i <= 10; i++ {
		createOneOffice(t, i)
	}
	
	offices, err := testQueries.ListAllOffices(context.Background(), &ListAllOfficesParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, len(offices), 10)
	require.NotEmpty(t, offices)
	
	for _, office := range offices {
		require.NotEmpty(t, office)
	}
	
	testQueries.ClearOffices(context.Background())
}

func TestListOfficesWithFilters(t *testing.T) {
	for i := 1; i <= 10; i++ {
		createOneOffice(t, i)
	}
	
	offices, err := testQueries.ListOfficesWithFilters(context.Background(), &ListOfficesWithFiltersParams{
		ProvinceID: sql.NullInt32{Int32: provinceID, Valid: true},
		Limit:      10,
		Offset:     0,
	})
	require.NoError(t, err)
	require.Equal(t, len(offices), 10)
	require.NotEmpty(t, offices)
	
	for _, office := range offices {
		require.NotEmpty(t, office)
	}
	
	testQueries.ClearOffices(context.Background())
}