# Go REST API with Gin

This project is a simple REST API built with Go and the Gin framework. It demonstrates how to structure a Go application with modular components like models, services, handlers, and routing. Additionally, it includes database migration support using SQL and the `migrate` tool.

## Project Structure
go-restapi-gin/
├── cmd/                     # Main application entry point 
│   └── main.go              # The main file to run the application
├── config/                  # Application configuration (environment, database, etc.)
│   └── config.go
├── internal/                # Internal application logic
│   ├── models/              
│   │   └── user.go
│   ├── services/            
│   │   └── user.go
│   ├── handlers/            
│   │   └── user.go
│   ├── routes/              
│   │   └── routes.go
├── migrations/              # Database migration files
│   └── 202412120001_create_users_table.sql
├── scripts/                 # (Optional) scripts for setup, run migration, deployment, etc.
│   └── setup.sh
├── .env                     # Environment variables for runtime configuration
├── go.mod                   # File modul Go
└── go.sum                   # Go checksum file for dependencies


## Setting Up the Project

1. **Clone the repository:**

   ```bash
   git clone https://github.com/dikiamust/gin-restapi
   cd go-restapi-gin
   ```

2. **Install Dependencies:**

   ```bash
   go mod tidy
   ```

3. **Set up environment variables:**

Copy .env.example to .env and fill in the necessary environment variables like your database credentials.

4. **Running Migrations:**

   ```bash
   migrate -database "postgres://username:password@localhost:5432/go_restapi_gin?sslmode=disable" -path ./migrations up
   ```
note:
- if you need to rollback the migration, run this command:

   ```bash
   migrate -database "postgres://username:password@localhost:5432/go_restapi_gin?sslmode=disable" -path ./migrations down
   ```
- To create a new migration file, use the following command (create role migration):
   ```bash
   migrate create -ext sql -dir migrations -seq create_roles_table
   ```

5. **Running the Application:**
   ```bash
   go run cmd/main.go
   ```