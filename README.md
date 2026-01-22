# Go MySQL CRUD API

This project is a **simple RESTful API** built using **Go (Golang)**, **Gin framework**, and **MySQL**, following a **clean layered architecture** (Controller → Service → Repository).

---

## Features

- REST API using Gin
- MySQL database connection
- CRUD operations (Create, Read, Update, Delete)
- Clean layered architecture
- JSON-based API
- Tested using Postman

---

## Project Architecture

- Controller → Service → Repository → Database

## Folder Structure
```
go-mysql-crud/

│
├── config/
│ └── db.go
│
├── controller/
│ └── user_controller.go
│
├── service/
│ └── user_service.go
│
├── repository/
│ └── user_repository.go
│
├── model/
│ └── user.go
│
├── routes/
│ └── routes.go
│
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Technologies Used

| Technology          | Purpose             |
|---------------------|---------------------|
| Go                  | Backend language    |
| Gin                 |  framework          |
| MySQL               | Database            |
| go-sql-driver/mysql | MySQL driver        |
| Postman             | API testing         |

---

## Prerequisites

- Go installed (`go version`)
- MySQL installed or running via Docker
- Postman installed

---

## Install & Build

Initialize module (if not already):
```bash
go mod init go-mysql-crud
go mod tidy
```

Get required packages:
```bash
go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
```

Build or run:
- Run directly:
```powershell
go run cmd\main.go
```

Server will start at: http://localhost:8080

## Database Setup

Create database and table:

```sql
CREATE DATABASE go_crud;
USE go_crud;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);
```

---

## API Endpoints

Base: http://localhost:8080

- Create user
  POST /users
  Request JSON:
  ```json
  {
    "name": "Dasun",
    "email": "dasun@gmail.com"
  }
  ```
  Example curl:
  ```bash
  curl -X POST http://localhost:8080/users -H "Content-Type: application/json" -d '{"name":"Dasun","email":"dasun@gmail.com"}'
  ```

- Get all users
  GET /users
  ```bash
  curl http://localhost:8080/users
  ```

- Get user by id
  GET /users/{id}
  ```bash
  curl http://localhost:8080/users/1
  ```

- Update user
  PUT /users/{id}
  Request JSON:
  ```json
  {
    "name": "Updated Name",
    "email": "updated@email.com"
  }
  ```

- Delete user
  DELETE /users/{id}

