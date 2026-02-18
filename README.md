# Product Service (Clean Architecture)

This project is implemented using Clean Architecture principles and Dependency Injection.

## Project Structure

- `cmd/server/main.go`: Application entry point.
- `internal/domain`: core business logic and entities.
  - `entity/`: GORM models and business entities.
  - `repository/`: Repository interfaces.
- `internal/usecase`: Business logic (Orchestrator).
- `internal/adapter`: Interface adapters.
  - `handler/`: HTTP handlers (Gin).
  - `repository/`: Repository implementations (Postgres).
- `internal/infrastructure`: Frameworks and drivers.
  - `db/`: Database connection.
  - `router/`: Router configuration.
  - `container/`: Dependency Injection container.

## Configuration Setup

The application uses environment variables for configuration. Follow these steps to set it up:

1. Copy the example environment file:
   ```bash
   cp app.env.example app.env
   ```
2. Open `app.env` and update the values:
   - `DB_DRIVER`: The database driver (default: `postgres`)
   - `DB_SOURCE`: The connection string (`postgresql://user:pass@host:port/dbname?sslmode=disable`)
   - `HTTP_SERVER_ADDRESS`: The address and port for the API (default: `0.0.0.0:8080`)

## How to Start Service

### 1. Using Docker Compose (Recommended)
This method automatically handles PostgreSQL setup and configuration.
```bash
docker-compose up -d --build
```

### 2. Running Locally
1. Ensure a PostgreSQL database is running.
2. Ensure your `app.env` is correctly configured to point to your local DB.
3. Run the application:
```bash
go run cmd/server/main.go
```

## API Documentation

### Generate Swagger Docs
If you change the API annotations, regenerate the documentation using:
```bash
swag init -g cmd/server/main.go --parseDependency --parseInternal
```

### Access Swagger UI
Swagger UI is available at:
`http://localhost:8080/api-docs/index.html`

## Testing (4 Levels)

Run all tests:
```bash
go test ./...
```

Run specific levels:
1. **Domain**: `go test ./internal/domain/entity/...`
2. **UseCase**: `go test ./internal/usecase/...`
3. **Repository**: `go test ./internal/adapter/repository/postgres/...`
4. **Component (E2E)**: `go test ./tests/component/...`

## API Features

- **POST /product**: Create a new product.
- **PATCH /product/{id}**: Partial update (supports undefined fields).
- **Dependency Injection**: Full constructor injection.
- **Clean Architecture**: Decoupled domain logic from infrastructure.
