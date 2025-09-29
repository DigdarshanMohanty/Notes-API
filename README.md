# Notes API

A simple RESTful API built with **Go (Gin + GORM)** that provides CRUD operations for managing notes.  
The project also includes **JWT-based authentication** and **bcrypt password hashing** for secure user management.

## Features :-
- User Authentication
  - Register with hashed passwords (bcrypt)
  - Login with JWT token generation
- Notes CRUD
  - Create, Read, Update, Delete notes
- Secure endpoints (JWT-protected routes)
- PostgreSQL integration with GORM ORM
- Configurable using `.env` file

## 🛠️ Tech Stack
- **Go** (Gin framework)
- **GORM** (ORM for PostgreSQL)
- **PostgreSQL** (Database)
- **JWT** (Authentication)
- **Bcrypt** (Password hashing)
- **godotenv** (Environment variable management)

## 📂 Project Structure

notes-api/
│── cmd/                     
│── internal/
│   ├── db/                  
│   │   └── db.go
│   │
│   ├── middleware/
│   │   └── jwt.go
│   │
│   ├── models/
│   │   ├── user.go
│   │   └── note.go
│   │
│   └── handlers/
│       ├── notes/           # Notes-related routes
│       │   └── notes.go
│       │
│       └── users/           # User-related routes
│           └── users.go
│
│── go.mod
│── go.sum
│── README.md
│── main.go                  # App entrypoint (router + wiring)
