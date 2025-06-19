# Authentication Backend in Go

A minimalist authentication backend in Go providing **user sign-up** and **login** functionality with secure password handling and JWT tokens.

## 🛠️ Features

- **User Registration** endpoint (`/signup`)
- **User Login** endpoint (`/login`)
- Password hashing using **bcrypt**
- JWT generation for authenticated sessions
- Basic request validation and error handling

## 🚀 Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Aadesh998/Authentication-backend.git
   cd Authentication-backend
   
2. **Install dependencies:**
  ```bash
  go mod download
```

3. **Configure environment variables**
Create a .env file (or export env vars) with settings like:
  ```bash
  JWT_SECRET=your_jwt_secret_key
  DB_DSN=your_database_connection_string  # if using a DB
  ```

4. **Run the App**
   ```bash
   go run cmd/main.go
   ```

5. **Test endpoints with curl or HTTP client:**
   * Sign up a new user:
    ```bash
    curl -X POST http://localhost:8080/signup \
      -H "Content-Type: application/json" \
      -d '{"email":"you@example.com","password":"Secret123"}'
    ```
   * Login:
    ```bash
    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"email":"you@example.com","password":"Secret123"}'
    ```
   * Response:
    ```bash
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpX..."
    }
    ```
   * Access protected routes using header:
    ```bash
    Authorization: Bearer <token>
    ```
