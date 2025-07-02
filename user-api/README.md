# User API

A simple REST API built in Go that manages users with in-memory storage.

## Features

- Add new users
- List all users
- Get user by ID
- In-memory data storage
- JSON API responses

## User Structure

```json
{
  "id": "string",
  "name": "string", 
  "email": "string"
}
```

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/users` | Add a new user |
| GET | `/users` | List all users |
| GET | `/users/{id}` | Get user by ID |
| GET | `/health` | Health check |

## Prerequisites

Before running the API, make sure you have the following installed:

- **Go** (version 1.21 or higher)
  - Download from: https://golang.org/dl/
  - Verify installation: `go version`

## Setup and Running

### Option 1: Using Makefile (Recommended)
The easiest way to get started is using the provided Makefile:

```bash
# Navigate to the project directory
cd user-api

# Setup and run in development mode
make dev
```

This single command will:
- Install all dependencies
- Start the server on port 8080
- Show helpful information about available endpoints

### Option 2: Manual Setup
1. **Navigate to the project directory:**
   ```bash
   cd user-api
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the server:**
   ```bash
   go run main.go
   ```

### Option 2: Build and Run
1. **Build the executable:**
   ```bash
   go build -o user-api main.go
   ```

2. **Run the executable:**
   ```bash
   ./user-api
   ```

### Option 3: Development Mode with Auto-reload
If you have `air` installed for hot reloading:
```bash
# Install air (optional)
go install github.com/cosmtrek/air@latest

# Run with auto-reload
air
```

## Available Makefile Commands

The project includes a comprehensive Makefile with several useful commands:

| Command | Description |
|---------|-------------|
| `make dev` | Setup dependencies and run the application |
| `make build` | Build the application executable |
| `make run` | Run the built executable |
| `make clean` | Clean build artifacts and module cache |
| `make test` | Run tests |
| `make deps` | Install dependencies only |
| `make install-air` | Install air for hot reloading |
| `make dev-air` | Run with air for hot reloading |
| `make check` | Check if Go is installed |
| `make setup` | Complete setup (check Go, install deps) |
| `make help` | Show all available commands |

### Examples:
```bash
# Quick development start
make dev

# Build and run production version
make build
make run

# Setup with hot reloading
make install-air
make dev-air

# Clean everything
make clean

# See all available commands
make help
```

## Server Information

- **Port:** 8080
- **URL:** http://localhost:8080
- **Health Check:** http://localhost:8080/health

When the server starts successfully, you should see output like:
```
Server starting on port 8080...
Available endpoints:
  POST   /users     - Add a new user
  GET    /users     - List all users
  GET    /users/{id} - Get user by ID
  GET    /health    - Health check
```

## Troubleshooting

### Port Already in Use
If port 8080 is already in use, you can change it by modifying the `port` variable in `main.go`:
```go
port := 8081  // Change to any available port
```

### Dependencies Issues
If you encounter dependency issues:
```bash
# Clean module cache
go clean -modcache

# Reinstall dependencies
go mod tidy
```

### Permission Issues (Linux/Mac)
If you get permission errors when running the executable:
```bash
chmod +x user-api
```

## Testing the API

### 1. Add a User
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```

### 2. List All Users
```bash
curl http://localhost:8080/users
```

### 3. Get User by ID
```bash
curl http://localhost:8080/users/{user-id}
```

### 4. Health Check
```bash
curl http://localhost:8080/health
```

## Example Responses

### Add User Response
```json
{
  "success": true,
  "message": "User added successfully",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

### List Users Response
```json
{
  "success": true,
  "message": "Found 2 users",
  "data": [
    {
      "id": "550e8400-e29b-41d4-a716-446655440000",
      "name": "John Doe",
      "email": "john@example.com"
    },
    {
      "id": "550e8400-e29b-41d4-a716-446655440001",
      "name": "Jane Smith",
      "email": "jane@example.com"
    }
  ]
}
```

### Get User by ID Response
```json
{
  "success": true,
  "message": "User found",
  "data": {
    "id": "550e8400-e29b-41d4-a716-446655440000",
    "name": "John Doe",
    "email": "john@example.com"
  }
}
```

## Notes

- Data is stored in memory and will be lost when the server restarts
- User IDs are automatically generated using UUID
- Name and email are required fields when adding a user
- The API returns consistent JSON responses with success status and messages 