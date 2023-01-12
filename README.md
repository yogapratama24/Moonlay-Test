# MOONLAT TECHNICAL TEST API

## Instructions

1. `go mod tidy` to install all module and package dependencies
2. `cp env.example .env` to make file .env
3. fill in the DB_NAME and DB_PASSWORD data sections in the .env file
4. in the .env file change data migrate to true, for database migration. after the migration is successful, please change it back to false
5. `go run main.go` to run this application and migrate to database.

note: please make sure that you have golang installed inside your PC/notebook
