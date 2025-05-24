package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func beforeTest() {
	cleanUpTest()
	SeedOffice(testQueries, context.Background())
}

func cleanUpTest() {
	testQueries.ClearSchools(context.Background())
	testQueries.ClearOffices(context.Background())
}

func createOneSchool(t *testing.T, iterate int) Schools {
	arg := CreateSchoolParams{
		Code:       fmt.Sprintf("test-%d", iterate),
		Name:       fmt.Sprintf("test-%d", iterate),
		OfficeID:   sql.NullInt64{Int64: 1, Valid: true},
		ProvinceID: provinceID,
		RegencyID:  regencyID,
		DistrictID: districtID,
		Email:      sql.NullString{String: fmt.Sprintf("test-%d@mail.com", iterate), Valid: true},
		Phone:      sql.NullString{String: "+628123456789", Valid: true},
		Address:    sql.NullString{String: "test address", Valid: true},
		LogoUrl:    sql.NullString{String: "test logo url", Valid: true},
		CreatedBy:  1,
	}

	school, err := testQueries.CreateSchool(context.Background(), &arg)
	require.NoError(t, err)

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
	beforeTest()
	createOneSchool(t, 1)
	cleanUpTest()
}

func TestGetSchoolById(t *testing.T) {
	beforeTest()
	school := createOneSchool(t, 1)

	checkSchool, err := testQueries.GetSchoolById(context.Background(), school.ID)
	require.NoError(t, err)
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

func TestUpdateSchool(t *testing.T) {
	beforeTest()
	school := createOneSchool(t, 1)
	
	arg, err := testQueries.UpdateSchool(context.Background(), &UpdateSchoolParams{
		Name:       "test update name",
		OfficeID:   sql.NullInt64{Int64: 1, Valid: true},
		ProvinceID: provinceID,
		RegencyID:  regencyID,
		DistrictID: districtID,
		Email:      sql.NullString{String: "testupdate@email", Valid: true},
		Phone:      sql.NullString{String: "+628123456789", Valid: true},
		Address:    sql.NullString{String: "testupdate address", Valid: true},
		LogoUrl:    sql.NullString{String: "testupdate logo url", Valid: true},
		ID:         school.ID,
	})

	school.Code = arg.Code
	school.Name = arg.Name
	school.Email = arg.Email
	school.OfficeID = arg.OfficeID
	school.ProvinceID = arg.ProvinceID
	school.RegencyID = arg.RegencyID
	school.DistrictID = arg.DistrictID
	school.Phone = arg.Phone
	school.Address = arg.Address
	school.LogoUrl = arg.LogoUrl	

	require.NoError(t, err)
	require.Equal(t, arg.ID, school.ID)
	require.Equal(t, arg.Code, school.Code)
	require.Equal(t, arg.Name, school.Name)
	require.Equal(t, arg.OfficeID, school.OfficeID)
	require.Equal(t, arg.ProvinceID, school.ProvinceID)
	require.Equal(t, arg.RegencyID, school.RegencyID)
	require.Equal(t, arg.DistrictID, school.DistrictID)
	require.Equal(t, arg.Email.String, school.Email.String)
	require.Equal(t, arg.Phone.String, school.Phone.String)
	require.Equal(t, arg.Address.String, school.Address.String)
	require.Equal(t, arg.LogoUrl.String, school.LogoUrl.String)
	require.Equal(t, arg.CreatedBy, school.CreatedBy)

	cleanUpTest()
}

func TestDeleteSchool(t *testing.T) {
	beforeTest()
	school := createOneSchool(t, 1)
	
	testQueries.DeleteSchool(context.Background(), school.ID)
	
	checkSchool, err := testQueries.GetSchoolById(context.Background(), school.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows)
	require.Equal(t, int(checkSchool.ID), 0)
	require.Equal(t, checkSchool.Name, "")
	require.Equal(t, checkSchool.Code, "")
	require.Equal(t, int(checkSchool.OfficeID.Int64), 0)
	require.Equal(t, int(checkSchool.ProvinceID), 0)
	require.Equal(t, int(checkSchool.RegencyID), 0)
	require.Equal(t, int(checkSchool.DistrictID), 0)
	require.Equal(t, checkSchool.Email.String, "")
	require.Equal(t, checkSchool.Phone.String, "")
	require.Equal(t, checkSchool.Address.String, "")
	require.Equal(t, checkSchool.LogoUrl.String, "")
	require.Equal(t, int(checkSchool.CreatedBy), 0)
	
	cleanUpTest()
}

func TestListAllSchools(t *testing.T) {
	beforeTest()

	for i := 1; i <= 10; i++ {
		createOneSchool(t, i)
	}

	schools, err := testQueries.ListAllSchools(context.Background(), &ListAllSchoolsParams{
		Limit:  10,
		Offset: 0,
	})
	require.NoError(t, err)
	require.Equal(t, len(schools), 10)
	require.NotEmpty(t, schools)

	for _, school := range schools {
		require.NotEmpty(t, school)
	}

	cleanUpTest()
}
