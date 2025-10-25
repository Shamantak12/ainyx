# User API

A simple REST API built with Go for managing users and calculating their age from their date of birth.

## What it does

- REST API using Go Fiber
- PostgreSQL database 
- Automatically calculates user age
- Input validation
- Structured logging with Zap
- Docker support
- Clean code structure

## 📊 API Endpoints

### Create User
```http
POST /api/v1/users
Content-Type: application/json

{
  "name": "Alice",
  "dob": "1990-05-10"
}
```

### Get User by ID
```http
GET /api/v1/users/:id
```

### List All Users
```http
GET /api/v1/users
```

### Update User
```http
PUT /api/v1/users/:id
Content-Type: application/json

{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}
```

### Delete User
```http
DELETE /api/v1/users/:id
```

## 🛠️ Tech Stack

- **Go 1.21** - Programming language
- **GoFiber v2** - Web framework
- **PostgreSQL** - Database
- **SQLC** - SQL code generation
- **Uber Zap** - Structured logging
- **go-playground/validator** - Input validation
- **Docker** - Containerization

## 📁 Project Structure

```
/
├── cmd/server/main.go          # Application entry point
├── config/                    # Configuration files
├── db/
│   ├── migrations/            # Database migrations
│   └── sqlc/                  # SQLC configuration and queries
├── internal/
│   ├── handler/               # HTTP handlers
│   ├── repository/            # Data access layer
│   ├── service/               # Business logic layer
│   ├── routes/                # Route definitions
│   ├── middleware/            # HTTP middleware
│   ├── models/                # Data models
│   └── logger/                # Logging configuration
├── Dockerfile                 # Docker configuration
├── docker-compose.yml         # Docker Compose setup
└── README.md                  # This file
```

## Getting Started

### What you need

- Go 1.21+
- PostgreSQL 12+
- Docker (optional but easier)

### Easy way with Docker

1. Clone this repo
   ```bash
   git clone <repository-url>
   cd user-api
   ```

2. Start everything with Docker
   ```bash
   docker-compose up -d
   ```

3. API will be running at `http://localhost:8080`

### Manual setup (if you want to run it yourself)

1. Install dependencies
   ```bash
   go mod download
   ```

2. Set up PostgreSQL
   ```sql
   CREATE DATABASE userdb;
   CREATE USER postgres WITH PASSWORD 'password';
   GRANT ALL PRIVILEGES ON DATABASE userdb TO postgres;
   ```

3. Create the users table (run the SQL in `db/migrations/001_create_users_table.sql`)

4. Set up environment variables
   ```bash
   cp env.example .env
   # Edit .env with your database details
   ```

5. Run the app
   ```bash
   go run cmd/server/main.go
   ```

## 🔧 Configuration

Environment variables can be set in a `.env` file or as system environment variables:

```env
# Server Configuration
SERVER_HOST=0.0.0.0
SERVER_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=userdb
DB_SSLMODE=disable

# Environment
ENV=development
```

## 📝 API Examples

### Create a User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "dob": "1990-05-15"
  }'
```

### Get User with Age
```bash
curl http://localhost:8080/api/v1/users/1
```

Response:
```json
{
  "id": 1,
  "name": "John Doe",
  "dob": "1990-05-15",
  "age": 34
}
```

### List All Users
```bash
curl http://localhost:8080/api/v1/users
```

### Update User
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Smith",
    "dob": "1991-03-20"
  }'
```

### Delete User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## 🧪 Testing

### Health Check
```bash
curl http://localhost:8080/health
```

## 🐳 Docker Commands

### Build and run with Docker Compose
```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down

# Rebuild and start
docker-compose up --build -d
```

### Build Docker image manually
```bash
# Build image
docker build -t user-api .

# Run container
docker run -p 8080:8080 user-api
```

## 📊 Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## 🔍 Features Implemented

### ✅ Core Requirements
- [x] RESTful API with Go Fiber
- [x] PostgreSQL database with SQLC
- [x] Dynamic age calculation
- [x] Input validation with go-playground/validator
- [x] Structured logging with Uber Zap
- [x] Clean HTTP status codes and error handling

### ✅ Bonus Features
- [x] Docker support with docker-compose
- [x] Request ID middleware
- [x] Request duration logging
- [x] CORS middleware
- [x] Health check endpoint
- [x] Structured error responses

## 🏗️ Architecture

The application follows clean architecture principles:

- **Handler Layer**: HTTP request/response handling
- **Service Layer**: Business logic and age calculation
- **Repository Layer**: Data access and database operations
- **Model Layer**: Data structures and validation

## 📈 Performance

- **Fiber framework** for high-performance HTTP handling
- **SQLC** for type-safe, optimized database queries
- **Structured logging** for better observability
- **Connection pooling** for database efficiency

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For issues and questions, please create an issue in the repository.



