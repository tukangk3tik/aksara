
migrate-up:
	migrate -path db/migrations -database "postgres://postgres:aksara2025@localhost:5432/aksara?sslmode=disable" -verbose up

migrate-up1:
	migrate -path db/migrations -database "postgres://postgres:aksara2025@localhost:5432/aksara?sslmode=disable" -verbose up 1

migrate-down:
	migrate -path db/migrations -database "postgres://postgres:aksara2025@localhost:5432/aksara?sslmode=disable" -verbose down 

migrate-down1:
	migrate -path db/migrations -database "postgres://postgres:aksara2025@localhost:5432/aksara?sslmode=disable" -verbose down 1

migrate-seed:
	migrate -path db/seed -database "postgres://postgres:aksara2025@localhost:5432/aksara?sslmode=disable" -verbose up 1

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