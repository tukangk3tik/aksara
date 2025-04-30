
### Make a new table:

migrate create -ext sql -dir db/migrations -seq create_new_table

### Run the table migration:

make migrate-up
