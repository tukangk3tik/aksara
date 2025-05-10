migrate-up:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up

migrate-up1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose up 1

migrate-down:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 

migrate-down1:
	migrate -path db/migrations -database "$(DB_URL)" -verbose down 1

migrate-seed:
	migrate -path db/seed -database "$(DB_URL)" -verbose up 1

import-loc-assets:
	@echo "Importing location assets from CSV files..."
	@if [ -z "$(DB_URL)" ]; then \
		echo "Error: DB_URL is not set. Please set it to your database connection string."; \
		echo "Example: make import-loc-assets DB_URL=\"postgres://postgres:password@localhost:5432/aksara?sslmode=disable\""; \
		exit 1; \
	fi
	@echo "Importing provinces.csv..."
	@echo "CREATE TEMP TABLE temp_provinces (id INTEGER, name TEXT);" > /tmp/import_provinces.sql
	@echo "\COPY temp_provinces FROM './docs/loc_assets/provinces.csv' DELIMITER ';' CSV HEADER;" >> /tmp/import_provinces.sql
	@echo "INSERT INTO loc_provinces (id, name) SELECT id, name FROM temp_provinces ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name;" >> /tmp/import_provinces.sql
	@echo "DROP TABLE temp_provinces;" >> /tmp/import_provinces.sql
	@psql $(DB_URL) -f /tmp/import_provinces.sql
	@rm /tmp/import_provinces.sql
	
	@echo "Importing regencies.csv..."
	@echo "CREATE TEMP TABLE temp_regencies (id INTEGER, province_id INTEGER, name TEXT);" > /tmp/import_regencies.sql
	@echo "\COPY temp_regencies FROM './docs/loc_assets/regencies.csv' DELIMITER ';' CSV HEADER;" >> /tmp/import_regencies.sql
	@echo "INSERT INTO loc_regencies (id, province_id, name) SELECT id, province_id, name FROM temp_regencies ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, province_id = EXCLUDED.province_id;" >> /tmp/import_regencies.sql
	@echo "DROP TABLE temp_regencies;" >> /tmp/import_regencies.sql
	@psql $(DB_URL) -f /tmp/import_regencies.sql
	@rm /tmp/import_regencies.sql
	
	@echo "Importing districts.csv..."
	@echo "CREATE TEMP TABLE temp_districts (id INTEGER, regency_id INTEGER, name TEXT);" > /tmp/import_districts.sql
	@echo "\COPY temp_districts FROM './docs/loc_assets/districts.csv' DELIMITER ';' CSV HEADER;" >> /tmp/import_districts.sql
	@echo "INSERT INTO loc_districts (id, regency_id, name) SELECT id, regency_id, name FROM temp_districts ON CONFLICT (id) DO UPDATE SET name = EXCLUDED.name, regency_id = EXCLUDED.regency_id;" >> /tmp/import_districts.sql
	@echo "DROP TABLE temp_districts;" >> /tmp/import_districts.sql
	@psql $(DB_URL) -f /tmp/import_districts.sql
	@rm /tmp/import_districts.sql
	
	@echo "All location assets imported successfully!"

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/tukangk3tik/aksara/db/sqlc Store

fe-server:
	cd _admin-frontend && npm run dev