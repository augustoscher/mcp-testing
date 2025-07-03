# Project Overview

This repository contains two main components:

## 1. `user-api` (Go)
A simple REST API built in Go for managing users with in-memory storage.

**Features:**
- Add, list, and retrieve users by ID or name
- In-memory data storage (data is lost on restart)
- JSON API responses

**Endpoints:**
- `POST /users` — Add a new user
- `GET /users` — List all users (optionally filter by name)
- `GET /users/{id}` — Get user by ID
- `GET /health` — Health check

**To run:**
```bash
cd user-api
make dev
```
The API will be available at [http://localhost:8080](http://localhost:8080).

## 2. `mcp-server` (Node.js/TypeScript)
An MCP (Model Context Protocol) server that exposes the user API as MCP tools, allowing clients to interact with the user API via the MCP protocol.

**Exposed MCP Tools:**
- `getUsers` — Retrieve all users
- `getUsersByName` — Retrieve users by name
- `getUserById` — Retrieve user by ID
- `registerUser` — Register a new user

**How it works:**
The MCP server acts as a bridge, forwarding tool calls to the REST API running at `http://localhost:8080`.

---

## How to Run and Interact with the MCP Server

### Prerequisites
- Node.js (see `.tool-versions` for recommended version)
- The `user-api` server must be running on `localhost:8080`

### Steps

1. **Start the user-api server:**
   ```bash
   cd user-api
   make dev
   # or: go run main.go
   ```

2. **Install dependencies and start the MCP server:**
   ```bash
   cd mcp-server
   yarn install
   yarn start
   # or: node src/index.ts
   ```

   The MCP server will start and listen for MCP protocol connections via stdio.

3. **Interact with the MCP server:**
   - Use an MCP-compatible client to connect to the server (e.g., via stdio).
   - Invoke the available tools (`getUsers`, `getUsersByName`, `getUserById`, `registerUser`) as defined in the MCP protocol.

**Example tool call**
```json
Prompt: Liste os usuarios na user api

Aqui está a lista de usuários atualmente cadastrados na user-api:

Jane Smith — jane@example.com (ID: 565b21a7-56ed-433b-9cd4-c14025efdd04)
John Doe — johndoe@example.com (ID: bd9cf73e-7958-40ee-8688-d0db87f403e8)
Jane Smith — jane@example.com (ID: 83b20f4e-1624-4a23-b686-7b9894fda3d5)
John Doe — johndoe@example.com (ID: 997482dc-43c3-4d86-8e6a-b92aae556770)

Se precisar de mais detalhes ou quiser filtrar por nome ou ID, é só pedir!
```

**Note:**
- The MCP server expects the user-api to be running and accessible at `http://localhost:8080`.
- All user data is stored in memory and will be lost when the user-api server restarts.
