# RUN MIGRATION

migrate -path db/migration -database "mysql://root:12345678@tcp(localhost:3306)/starter-go?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

# Create Migration File

migrate create -ext sql -dir db/migration -seq adding_phone_in_users_table
