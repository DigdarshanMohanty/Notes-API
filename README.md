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

# API Endpoints

Auth
Method	Endpoint	Description
POST	/register	Register new user
POST	/login	Login and get JWT token
Notes
Method	Endpoint	Description
GET	/notes	Get all notes (JWT required)
GET	/notes/:id	Get a single note
POST	/notes	Create a new note
PUT	/notes/:id	Update a note
DELETE	/notes/:id	Delete a note