package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func beforeTest() {
	SeedOffice(testQueries, context.Background())
}

func cleanUpTest() {
	testQueries.ClearSchools(context.Background())
	testQueries.ClearOffices(context.Background())
}

func createOneSchool(t *testing.T) Schools {
	beforeTest()
	
	arg := CreateSchoolParams{
		ID:         1,
		Code:       "testschool",
		Name:       "test scholl",
		OfficeID:   sql.NullInt64{Int64: 1, Valid: true},
		ProvinceID: provinceID,
		RegencyID:  regencyID,
		DistrictID: districtID,
		Email:      sql.NullString{String: "testschool@mail.com", Valid: true},
		Phone:      sql.NullString{String: "+628123456789", Valid: true},
		Address:    sql.NullString{String: "test address", Valid: true},
		LogoUrl:    sql.NullString{String: "test logo url", Valid: true},
		CreatedBy:  1,
	}

	school, err := testQueries.CreateSchool(context.Background(), &arg)
	require.NoError(t, err)

	require.Equal(t, arg.ID, school.ID)
	require.Equal(t, arg.Code, school.Code)
	require.Equal(t, arg.Name, school.Name)
	require.Equal(t, arg.OfficeID, school.OfficeID)
	require.Equal(t, arg.ProvinceID, school.ProvinceID)
	require.Equal(t, arg.RegencyID, school.RegencyID)
	require.Equal(t, arg.DistrictID, school.DistrictID)
	require.Equal(t, arg.Email, school.Email)
	require.Equal(t, arg.Phone, school.Phone)
	require.Equal(t, arg.Address, school.Address)
	require.Equal(t, arg.LogoUrl, school.LogoUrl)
	require.Equal(t, arg.CreatedBy, school.CreatedBy)

	return school
}

func TestCreateSchool(t *testing.T) {
	createOneSchool(t)
	cleanUpTest()
}

func TestGetSchoolById(t *testing.T) {
	school := createOneSchool(t)

	checkSchool, err := testQueries.GetSchoolById(context.Background(), school.ID)
	require.NoError(t, err)
	require.Equal(t, checkSchool.ID, school.ID)
	require.Equal(t, checkSchool.Code, school.Code)
	require.Equal(t, checkSchool.Name, school.Name)
	require.Equal(t, checkSchool.OfficeID, school.OfficeID)
	require.Equal(t, checkSchool.ProvinceID, school.ProvinceID)
	require.Equal(t, checkSchool.RegencyID, school.RegencyID)
	require.Equal(t, checkSchool.DistrictID, school.DistrictID)
	require.Equal(t, checkSchool.Email, school.Email)
	require.Equal(t, checkSchool.Phone, school.Phone)
	require.Equal(t, checkSchool.Address, school.Address)
	require.Equal(t, checkSchool.LogoUrl, school.LogoUrl)
	require.Equal(t, checkSchool.CreatedBy, school.CreatedBy)

	cleanUpTest()
}