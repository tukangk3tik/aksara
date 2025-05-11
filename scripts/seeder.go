package main

import (
	"context"
	"database/sql"
	"log"

	db "github.com/tukangk3tik/aksara/db/sqlc"
	"github.com/tukangk3tik/aksara/utils"
)

var OfficeSeed = []db.Offices{
	{
		ID: 1,
		Code: "KP01",
		Name: "Kantor Pusat",
		ProvinceID: 31,
		RegencyID: 3171,
		DistrictID: 317103,
		Email: "kantor_pusat@aksara.com",
		Phone: sql.NullString{String: "+62", Valid: true},
		Address: sql.NullString{String: "Kantor Pusat", Valid: true},
		LogoUrl: sql.NullString{String: "logo_url", Valid: true},
		CreatedBy: 0,
	},
	{
		ID: 2,
		Code: "KDB01",
		Name: "Kantor Kab. Belu",
		ProvinceID: 53,
		RegencyID: 5304,
		DistrictID: 530412,
		Email: "kd_belu01@aksara.com",
		Phone: sql.NullString{String: "+62", Valid: true},
		Address: sql.NullString{String: "Kantor Kab. Belu", Valid: true},
		LogoUrl: sql.NullString{String: "logo_url", Valid: true},
		CreatedBy: 0,
	},
}

var UsersSeed = []db.Users{
	{
		ID: 1,
		Name: "Felix",
		Fullname: "Felix Seran",
		Email: "felix@aksara.com",
		Password: "$2y$10$agn8hHbQEc9dlNhDAb.X3OFuwkdS.0oaT19/FHXs1CQtYu7WbCmge",
		UserRoleID: 1,
		OfficeID: sql.NullInt64{Int64: 1, Valid: true},
		SchoolID: sql.NullInt64{},
		IsSuperAdmin: sql.NullBool{Bool: true, Valid: true},
	},
	{
		ID: 2,
		Name: "Admin Kab. Belu",
		Fullname: "Mateus Asato",
		Email: "mateus@aksara.com",
		Password: "password1234",
		UserRoleID: 2,
		OfficeID: sql.NullInt64{Int64: 2, Valid: true},
		SchoolID: sql.NullInt64{},
		IsSuperAdmin: sql.NullBool{Bool: false, Valid: true},
	},
	{
		ID: 3,
		Name: "operatorSdnAsulun",
		Fullname: "Operator Sdn Asulun",
		Email: "asulun_sdn@aksara.com",
		Password: "password1234",
		UserRoleID: 3,
		OfficeID: sql.NullInt64{Int64: 2, Valid: true},
		SchoolID: sql.NullInt64{},
	},
	{
		ID: 4,
		Name: "guruAndri",
		Fullname: "Andri Raosa",
		Email: "andri_guru_asulun@aksara.com",
		Password: "password1234",
		UserRoleID: 4,
		OfficeID: sql.NullInt64{Int64: 2, Valid: true},
		SchoolID: sql.NullInt64{},
	},
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}	
	defer conn.Close()

	store := db.NewStore(conn)	
	for _, userRole := range db.UserRolesSeed {
		_, err := store.CreateUserRole(context.Background(), &db.CreateUserRoleParams{
			ID:   userRole.ID,
			Name: userRole.Name,
		})

		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"user_roles_pkey\"" {
				continue
			}
			log.Fatal("cannot create user role:", err)
		}
	}

	for _, office := range OfficeSeed {
		_, err := store.CreateOffice(context.Background(), &db.CreateOfficeParams{
			Code: office.Code,
			Name: office.Name,
			ProvinceID: office.ProvinceID,
			RegencyID: office.RegencyID,
			DistrictID: office.DistrictID,
			Email: office.Email,
			Phone: office.Phone,
			Address: office.Address,
			LogoUrl: office.LogoUrl,
			CreatedBy: office.CreatedBy,
		})

		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"offices_email_key\"" || err.Error() == "pq: duplicate key value violates unique constraint \"offices_code_key\"" {
				continue
			}
			log.Fatal("cannot create office:", err)
		}
	}

	for _, user := range UsersSeed {
		_, err := store.CreateUser(context.Background(), &db.CreateUserParams{
			Name: user.Name,
			Fullname: user.Fullname,
			Email: user.Email,
			Password: user.Password,
			UserRoleID: user.UserRoleID,
			OfficeID: user.OfficeID,
			SchoolID: user.SchoolID,
		})

		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"" || err.Error() == "pq: duplicate key value violates unique constraint \"users_code_key\"" {
				continue
			}
			log.Fatal("cannot create user:", err)
		}
	}
}