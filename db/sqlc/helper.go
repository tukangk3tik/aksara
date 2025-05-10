package db

import (
	"context"
	"database/sql"
	"log"
)

var UserRolesSeed = []UserRoles{
	{
		Name: "Super Admin",
	},
	{
		Name: "Admin Regional",
	},
	{
		Name: "Operator Sekolah",
	},
	{
		Name: "Guru",
	},
}

func SeedUserRoles(testQueries *Queries, ctx context.Context) {
	for _, userRole := range UserRolesSeed {
		_, err := testQueries.CreateUserRole(ctx, userRole.Name)
		if err != nil {
			log.Fatal("cannot create user role:", err)
		}
	}
}

func SeedOffice(testQueries *Queries, ctx context.Context) {
	provinceID := int32(53)
	regencyID := int32(5371)
	districtID := int32(537104)

	var offices = []CreateOfficeParams{
		{
			Code:       "test",
			Name:       "test",
			ProvinceID: provinceID,
			RegencyID:  regencyID,
			DistrictID: districtID,
			Email:      "test@mail.com",
			Phone:      sql.NullString{String: "+628123456789", Valid: true},
			Address:    sql.NullString{String: "test address", Valid: true},
			LogoUrl:    sql.NullString{String: "test logo url", Valid: true},
			CreatedBy:  1,
		},
	}

	for _, office := range offices {
		_, err := testQueries.CreateOffice(ctx, &office)
		if err != nil {
			log.Fatal("cannot create office:", err)
		}
	}
}
