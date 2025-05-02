package db

import (
	"context"
	"log"
)

func SeedUserRoles(testQueries *Queries, ctx context.Context) {
	var userRoles = []UserRoles{
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

	for _, userRole := range userRoles {
		_, err := testQueries.CreateUserRole(ctx, userRole.Name)
		if err != nil {
			log.Fatal("cannot create user role:", err)
		}
	}
}

func SeedOffice() {
	
}

func SeedSchool() {
	
}

func SeedUser() {
	
}
