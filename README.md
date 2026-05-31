# URL Shortener API

A RESTful URL Shortener service built with Go, Gin, PostgreSQL, GORM and JWT Authentication.

## Features

* User Registration
* User Login
* JWT Authentication
* URL Shortening
* URL Redirection
* Click Counter
* User Dashboard
* Admin Dashboard
* Role-Based Authorization
* PostgreSQL Database
* Docker Compose Support
* Unit Testing

---

## Tech Stack

* Go
* Gin
* PostgreSQL
* GORM
* JWT
* Docker
* Viper
* bcrypt

---

## Project Structure

```text
.
├── cmd
│   └── server
│       └── main.go
│
├── internal
│   ├── config
│   ├── database
│   ├── handler
│   ├── middleware
│   ├── models
│   ├── repository
│   ├── routes
│   ├── service
│   └── utils
│
├── tests
│
├── .env
├── docker-compose.yml
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

---

## Environment Variables

Create a `.env` file in the project root:

```env
APP_PORT=8080

DB_HOST=localhost
DB_PORT=5433
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=url_shortener
DB_SSLMODE=disable

JWT_SECRET=super-secret-key
```

---

## Run PostgreSQL

```bash
docker compose up -d
```

Check running containers:

```bash
docker ps
```

---

## Run Application

```bash
go run cmd/server/main.go
```

Application will be available at:

```text
http://localhost:8080
```

---

## API Endpoints

### Authentication

#### Register

```http
POST /api/v1/auth/register
```

Request:

```json
{
  "username": "amin",
  "email": "amin@test.com",
  "password": "123456"
}
```

---

#### Login

```http
POST /api/v1/auth/login
```

Request:

```json
{
  "email": "amin@test.com",
  "password": "123456"
}
```

Response:

```json
{
  "token": "JWT_TOKEN"
}
```

---

### URL Management

#### Create Short URL

```http
POST /api/v1/urls
```

Headers:

```text
Authorization: Bearer JWT_TOKEN
```

Request:

```json
{
  "original_url": "https://quera.org"
}
```

---

#### Get User URLs

```http
GET /api/v1/urls/me
```

Headers:

```text
Authorization: Bearer JWT_TOKEN
```

---

### Redirect

```http
GET /{shortCode}
```

Example:

```text
http://localhost:8080/UQhFIAh
```

---

### Admin Endpoints

#### Get All URLs

```http
GET /api/v1/admin/urls
```

#### Delete URL

```http
DELETE /api/v1/admin/urls/{id}
```

Admin endpoints require:

```text
role = admin
```

---

## Running Tests

```bash
go test ./...
```

Example Output:

```text
ok github.com/aminramezanipoor/url-shortener/tests
```

---

## Database Schema

### Users

| Field    | Type   |
| -------- | ------ |
| id       | UUID   |
| username | string |
| email    | string |
| password | string |
| role     | string |

### URLs

| Field        | Type   |
| ------------ | ------ |
| id           | UUID   |
| original_url | string |
| short_code   | string |
| clicks       | int    |
| user_id      | UUID   |

---

## Author

Amin Ramezanipoor

Bootcamp Project - URL Shortener API
