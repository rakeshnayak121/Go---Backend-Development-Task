# User Management REST API

A RESTful API built using Go, Fiber, PostgreSQL, SQLC, Validator, and Zap Logger to manage users and dynamically calculate age from Date of Birth.

## Features

* Create User
* Get User By ID
* Get All Users
* Update User
* Delete User
* Dynamic Age Calculation
* PostgreSQL Database
* SQLC Database Access Layer
* Input Validation using go-playground/validator
* Logging using Uber Zap
* Pagination Support
* Unit Testing

---

## Tech Stack

* Go
* Fiber
* PostgreSQL
* SQLC
* Uber Zap
* go-playground/validator

---

## Project Structure

```text
.
├── config
├── database
├── db
│   ├── migrations
│   ├── query
│   └── sqlc
├── handlers
├── logger
├── models
├── routes
├── tests
├── utils
├── .env
├── go.mod
├── go.sum
└── main.go
```

---

## Environment Variables

Create a `.env` file:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=userdb
```

---

## Database Schema

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    dob DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Installation

Clone Repository:

```bash
git clone https://github.com/rakeshnayak121/Go---Backend-Development-Task.git
cd Go---Backend-Development-Task
```

Install Dependencies:

```bash
go mod tidy
```

Run Application:

```bash
go run .
```

---

## API Endpoints

### Health Check

```http
GET /health
```

Response:

```json
{
  "status": "ok"
}
```

---

### Create User

```http
POST /users
```

Request:

```json
{
  "name": "Rakesh",
  "dob": "2002-07-15"
}
```

Response:

```json
{
  "message": "User Created Successfully",
  "id": 1
}
```

---

### Get User By ID

```http
GET /users/:id
```

Response:

```json
{
  "id": 1,
  "name": "Rakesh",
  "dob": "2002-07-15",
  "age": 23
}
```

---

### Get All Users

```http
GET /users
```

Response:

```json
[
  {
    "id": 1,
    "name": "Rakesh",
    "dob": "2002-07-15",
    "age": 23
  }
]
```

---

### Pagination

```http
GET /users?page=1&limit=5
```

---

### Update User

```http
PUT /users/:id
```

Request:

```json
{
  "name": "Updated User",
  "dob": "2001-01-10"
}
```

Response:

```json
{
  "message": "User Updated Successfully"
}
```

---

### Delete User

```http
DELETE /users/:id
```

Response:

```http
204 No Content
```

---

## Run Tests

```bash
go test ./...
```

Expected:

```text
ok user-api/tests
```

---

## SQLC Generate

Whenever queries are modified:

```bash
sqlc generate
```

---

## Author
