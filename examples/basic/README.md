# Basic Skeleton Plugin Example

This example demonstrates the minimal setup required to use the GoREST Skeleton Plugin along with authentication.

## Setup

1. Copy the environment file:
   ```bash
   cp .env.example .env
   ```

2. Edit `.env` and configure your database connection and JWT secret:
   ```env
   DATABASE_URL="postgres://user:password@localhost:5432/gorest_skeleton_example?sslmode=disable"
   JWT_SECRET="your-super-secret-jwt-key-at-least-32-characters-long"
   ```

3. Install dependencies:
   ```bash
   go mod download
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8000` with the following endpoints:

- **Auth endpoints:**
  - `POST /auth/register` - Register a new user
  - `POST /auth/login` - Login and get JWT token
  - `POST /auth/refresh` - Refresh JWT token

- **Skeleton endpoints (require authentication):**
  - `POST /api/skeleton` - Create an item
  - `GET /api/skeleton/:id` - Get item by ID
  - `GET /api/skeleton` - List items
  - `PUT /api/skeleton/:id` - Update an item
  - `DELETE /api/skeleton/:id` - Delete an item

## Testing the Endpoints

### 1. Register a new user

```bash
curl -X POST http://localhost:8000/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123",
    "name": "Test User"
  }'
```

Save the `token` from the response.

### 2. Create an item

```bash
TOKEN="your-jwt-token-here"

curl -X POST http://localhost:8000/api/skeleton \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "name": "My First Item",
    "description": "This is a test item",
    "active": true
  }'
```

### 3. List items

```bash
curl -H "Authorization: Bearer $TOKEN" \
  http://localhost:8000/api/skeleton
```

### 4. Get item by ID

```bash
ITEM_ID="your-item-id-here"

curl -H "Authorization: Bearer $TOKEN" \
  http://localhost:8000/api/skeleton/$ITEM_ID
```

### 5. Update an item

```bash
curl -X PUT http://localhost:8000/api/skeleton/$ITEM_ID \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "name": "Updated Item Name",
    "active": false
  }'
```

### 6. Delete an item

```bash
curl -X DELETE http://localhost:8000/api/skeleton/$ITEM_ID \
  -H "Authorization: Bearer $TOKEN"
```

## Database Setup

The application will automatically run migrations on startup if `migrations.auto_migrate: true` is set in `gorest.yaml`.

To manually run migrations:

```bash
gorest migrate up
```

## Troubleshooting

### Database connection errors

Make sure PostgreSQL is running and the connection string in `.env` is correct.

### "user not authenticated" errors

Ensure you're including the JWT token in the `Authorization` header:
```
Authorization: Bearer your-token-here
```

### Token expired

If your token expires (default 15 minutes), use the refresh endpoint or login again.

## Next Steps

- Explore the source code in the parent directory
- Modify the skeleton plugin to fit your needs
- Add custom business logic to handlers
- Create your own plugin based on this template
