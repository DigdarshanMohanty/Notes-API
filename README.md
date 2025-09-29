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

##  Tech Stack
- **Go** (Gin framework)
- **GORM** (ORM for PostgreSQL)
- **PostgreSQL** (Database)
- **JWT** (Authentication)
- **Bcrypt** (Password hashing)
- **godotenv** (Environment variable management)

## 📡 API Endpoints

### 🔑 Authentication

| Method | Endpoint   | Description          | Auth Required |
|--------|-----------|----------------------|---------------|
| POST   | `/register` | Register a new user   | ❌ No |
| POST   | `/login`    | Login and get JWT     | ❌ No |

---

### USers (JWT required)

| Method | Endpoint        | Description            | Auth Required |
|--------|----------------|------------------------|---------------|
| POST   | `/user/:id`       | Display user details      | ✅ Yes |

---

### Notes (JWT required)

| Method | Endpoint        | Description            | Auth Required |
|--------|----------------|------------------------|---------------|
| POST   | `/notes`       | Create a new note      | ✅ Yes |
| GET    | `/notes`       | Get all notes (user)   | ✅ Yes |
| PATCH    | `/notes/:id`   | Update a note          | ✅ Yes |
| DELETE | `/notes/:id`   | Delete a note          | ✅ Yes |
