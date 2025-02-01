# 🚀 Go CRUD API

A RESTful API developed in **Go (Golang)**, designed with a focus on performance, simplicity, and best development practices. The project implements CRUD operations for user management using a clean and modular architecture.

## 📦 Technologies Used

- **Go (Golang):** The main language, chosen for its efficiency and performance.
- **[Chi](https://github.com/go-chi/chi):** A lightweight and powerful routing framework.
- **In-Memory Database:** Simple and fast for temporary operations.
- **Middleware:** Robust error handling, logging, and request tracking.

## 🔑 Features

- ✅ **Full User CRUD:** Create, read, update, and delete operations.
- 📊 **Structured Logs:** Using `log/slog` for easy debugging.
- ⚡ **Optimized Performance:** HTTP server timeouts configuration.
- 🔐 **Error Handling:** Middleware for elegant error capturing and responses.

## 🗂️ Project Structure

```
├── main.go            # Application entry point
├── api/               # API logic (routes, handlers, middlewares)
│   ├── models/        # Data model definitions
│   ├── dtos/          # Data Transfer Objects
│   └── helpers/       # Helper functions
├── database/          # In-memory database configuration
└── go.mod             # Dependency management
```

## 🚀 How to Run the Project

```bash
# Clone the repository
git clone https://github.com/walker007/go-crud
cd go-crud

# Install dependencies
go mod tidy

# Run the application
go run main.go
```

The API will be available at `http://localhost:8080`.

## 📈 Available Endpoints

- `GET /api/users` - List all users
- `POST /api/users` - Create a new user
- `GET /api/users/{id}` - Get user by ID
- `PUT /api/users/{id}` - Update user information
- `DELETE /api/users/{id}` - Delete a user

## 🚀 Technical Challenges Overcome

- **Efficient Error Management:** Implementation of middlewares for graceful error handling.
- **Clean Architecture:** Clear separation between application layers, making the code easy to maintain and understand.
- **Simplicity and Efficiency:** Use of an in-memory database for fast and temporary operations.

---

Made with ❤️ in Go.

