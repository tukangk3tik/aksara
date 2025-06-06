// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	ClearOffices(ctx context.Context) error
	ClearSchools(ctx context.Context) error
	ClearUserRoles(ctx context.Context) error
	ClearUsers(ctx context.Context) error
	CreateOffice(ctx context.Context, arg *CreateOfficeParams) (Offices, error)
	CreateSchool(ctx context.Context, arg *CreateSchoolParams) (Schools, error)
	CreateUser(ctx context.Context, arg *CreateUserParams) (Users, error)
	CreateUserRole(ctx context.Context, arg *CreateUserRoleParams) (UserRoles, error)
	DeleteOffice(ctx context.Context, id int64) (sql.Result, error)
	DeleteSchool(ctx context.Context, id int64) (sql.Result, error)
	DeleteUser(ctx context.Context, id int64) (sql.Result, error)
	DeleteUserRole(ctx context.Context, id int32) (sql.Result, error)
	GetOfficeById(ctx context.Context, id int64) (GetOfficeByIdRow, error)
	GetSchoolById(ctx context.Context, id int64) (GetSchoolByIdRow, error)
	GetUserByEmail(ctx context.Context, email string) (Users, error)
	GetUserById(ctx context.Context, id int64) (Users, error)
	GetUserRoleById(ctx context.Context, id int32) (UserRoles, error)
	ListAllOffices(ctx context.Context, arg *ListAllOfficesParams) ([]ListAllOfficesRow, error)
	ListAllSchools(ctx context.Context, arg *ListAllSchoolsParams) ([]ListAllSchoolsRow, error)
	ListAllUserRoles(ctx context.Context, arg *ListAllUserRolesParams) ([]UserRoles, error)
	ListAllUsers(ctx context.Context, arg *ListAllUsersParams) ([]Users, error)
	// Optional province filter
	// Optional regency filter
	// Optional district filter
	// Optional name search
	ListOfficesWithFilters(ctx context.Context, arg *ListOfficesWithFiltersParams) ([]ListOfficesWithFiltersRow, error)
	ListSchoolsByDistrict(ctx context.Context, arg *ListSchoolsByDistrictParams) ([]Schools, error)
	ListSchoolsByOffice(ctx context.Context, arg *ListSchoolsByOfficeParams) ([]Schools, error)
	ListSchoolsByProvince(ctx context.Context, arg *ListSchoolsByProvinceParams) ([]Schools, error)
	ListSchoolsByRegency(ctx context.Context, arg *ListSchoolsByRegencyParams) ([]Schools, error)
	LocationDistrictByRegency(ctx context.Context, arg *LocationDistrictByRegencyParams) ([]LocDistricts, error)
	LocationProvince(ctx context.Context, arg *LocationProvinceParams) ([]LocProvinces, error)
	LocationRegencyByProvince(ctx context.Context, arg *LocationRegencyByProvinceParams) ([]LocRegencies, error)
	TotalListAllOffices(ctx context.Context) (int64, error)
	TotalListAllSchools(ctx context.Context) (int64, error)
	TotalListAllUsers(ctx context.Context) (int64, error)
	// Optional province filter
	// Optional regency filter
	// Optional district filter
	// Optional name search
	TotalListOfficesWithFilters(ctx context.Context, arg *TotalListOfficesWithFiltersParams) (int64, error)
	UpdateOffice(ctx context.Context, arg *UpdateOfficeParams) (Offices, error)
	UpdateSchool(ctx context.Context, arg *UpdateSchoolParams) (Schools, error)
	UpdateUser(ctx context.Context, arg *UpdateUserParams) (Users, error)
}

var _ Querier = (*Queries)(nil)
