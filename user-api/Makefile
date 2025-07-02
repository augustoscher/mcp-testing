# User API Makefile

# Variables
BINARY_NAME=user-api
MAIN_FILE=main.go
PORT=8080

# Default target
.PHONY: help
help: ## Show this help message
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: dev
dev: ## Setup dependencies and run the application in development mode
	@echo "🚀 Setting up and running User API in development mode..."
	@echo "📦 Installing dependencies..."
	@go mod tidy
	@echo "🔧 Starting server on port $(PORT)..."
	@echo "📍 API will be available at: http://localhost:$(PORT)"
	@echo "🔍 Health check: http://localhost:$(PORT)/health"
	@echo ""
	@go run $(MAIN_FILE)

.PHONY: build
build: ## Build the application
	@echo "🔨 Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Build complete: ./$(BINARY_NAME)"

.PHONY: run
run: ## Run the built application
	@echo "🚀 Running $(BINARY_NAME)..."
	@./$(BINARY_NAME)

.PHONY: clean
clean: ## Clean build artifacts
	@echo "🧹 Cleaning build artifacts..."
	@rm -f $(BINARY_NAME)
	@go clean -modcache
	@echo "✅ Clean complete"

.PHONY: test
test: ## Run tests
	@echo "🧪 Running tests..."
	@go test -v ./...

.PHONY: deps
deps: ## Install dependencies
	@echo "📦 Installing dependencies..."
	@go mod tidy
	@echo "✅ Dependencies installed"

.PHONY: install-air
install-air: ## Install air for hot reloading
	@echo "📦 Installing air for hot reloading..."
	@go install github.com/cosmtrek/air@latest
	@echo "✅ Air installed"

.PHONY: dev-air
dev-air: deps ## Run with air for hot reloading (requires air to be installed)
	@echo "🔥 Starting with hot reloading..."
	@echo "📍 API will be available at: http://localhost:$(PORT)"
	@air

.PHONY: check
check: ## Check if Go is installed
	@echo "🔍 Checking Go installation..."
	@go version || (echo "❌ Go is not installed. Please install Go from https://golang.org/dl/" && exit 1)
	@echo "✅ Go is installed"

.PHONY: setup
setup: check deps ## Complete setup (check Go, install deps)
	@echo "✅ Setup complete! Run 'make dev' to start the server" 