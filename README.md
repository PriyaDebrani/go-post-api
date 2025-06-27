# Go Post API

A RESTful API built with Go (Golang) using the Gin framework for user management with custom validation rules.

## 🚀 Features

- **User Management**: Create users with validation
- **Custom Validation**: PAN (Permanent Account Number) and mobile number validation
- **Middleware**: Request latency logging
- **Clean Architecture**: Separation of concerns with repository, service, and transport layers
- **Comprehensive Testing**: Unit tests for all layers

## 📋 Prerequisites

- Go 1.22.4 or higher
- Git

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-post-api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## 🏗️ Project Structure

```
go-post-api/
├── main.go              # Application entry point
├── go.mod               # Go module file
├── go.sum               # Dependency checksums
├── model/
│   └── user.go          # User data model
├── middleware/
│   └── latency.go       # Request latency logging middleware
├── validator/
│   └── custom.go        # Custom validation rules
├── transport.go         # HTTP transport layer
├── service.go           # Business logic layer
├── repo.go              # Data access layer
└── *_test.go            # Test files
```

## 📚 API Documentation

### Create User

**Endpoint:** `POST /users`

**Request Body:**
```json
{
  "name": "John Doe",
  "pan": "ABCDE1234F",
  "mobile": "9876543210",
  "email": "john.doe@example.com"
}
```

**Response (Success - 200):**
```json
{
  "message": "User created successfully"
}
```

**Response (Validation Error - 400):**
```json
{
  "errors": {
    "Name": "Name is required",
    "Email": "Invalid email format",
    "PAN": "Invalid PAN format",
    "Mobile": "Mobile must be a 10 digit number"
  }
}
```

## 🔍 Validation Rules

### User Model Validation
- **Name**: Required field
- **PAN**: Required, must match pattern `^[A-Z]{5}[0-9]{4}[A-Z]$`
- **Mobile**: Required, must be exactly 10 digits
- **Email**: Required, must be valid email format

### PAN Format
- 5 uppercase letters + 4 digits + 1 uppercase letter
- Example: `ABCDE1234F`

### Mobile Format
- Exactly 10 digits
- Example: `9876543210`

## 🧪 Testing

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run specific test files:
```bash
go test transport_test.go transport.go
go test service_test.go service.go
go test repo_test.go repo.go
```

## 🔧 Configuration

The application runs on port `8080` by default. You can modify this in `main.go`:

```go
r.Run(":8080") // Change port here
```

## 📊 Middleware

### Latency Logger
Automatically logs the duration of each HTTP request to help with performance monitoring.

## 🏛️ Architecture

This project follows a clean architecture pattern with three main layers:

1. **Transport Layer** (`transport.go`): Handles HTTP requests/responses
2. **Service Layer** (`service.go`): Contains business logic
3. **Repository Layer** (`repo.go`): Manages data access (currently in-memory)

## 🛡️ Error Handling

The API provides comprehensive error handling:
- JSON validation errors
- Field validation errors with detailed messages
- Custom validation for PAN and mobile numbers

## 📦 Dependencies

- **Gin**: HTTP web framework
- **Validator**: Data validation library
- **Testify**: Testing utilities
