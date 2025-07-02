# Add User
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com"
  }'
```
---
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Smith",
    "email": "jane@example.com"
  }'
```

# List Users
```bash
curl http://localhost:8080/users
```

# Get User by ID
```bash
curl http://localhost:8080/users/3574af65-5c65-41d6-bc5f-6b8c1c228171
```

# List Users by Name
```bash
curl "http://localhost:8080/users?name=john%20doe" | jq
```

# Health Check
```bash
curl http://localhost:8080/health
```

