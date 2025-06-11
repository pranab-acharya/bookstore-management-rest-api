# Bookstore Management REST API

A CRUD-based REST API for managing a bookstore, built with Go.  
Uses `gorilla/mux` for routing, `jinzhu/gorm` as ORM, and MySQL as the database via `jinzhu/gorm/dialects/mysql`.

## Features

- Create, Read, Update, Delete (CRUD) operations for books
- RESTful API endpoints
- MySQL database integration

## Dependencies

- [gorilla/mux](https://github.com/gorilla/mux)
- [jinzhu/gorm](https://github.com/jinzhu/gorm)
- [jinzhu/gorm/dialects/mysql](https://github.com/jinzhu/gorm/tree/master/dialects/mysql)

## Getting Started

### Prerequisites

- Go 1.18+
- MySQL server

### Installation

```bash
git clone https://github.com/pranab-acharya/bookstore-management-rest-api.git
cd bookstore-management-rest-api
go mod tidy
```

### Configuration

Update your MySQL credentials in the code or use environment variables.

### Running the API

```bash
go run main.go
```

## API Endpoints

| Method | Endpoint        | Description           |
|--------|----------------|----------------------|
| GET    | /books         | List all books       |
| GET    | /books/{id}    | Get book by ID       |
| POST   | /books         | Create new book      |
| PUT    | /books/{id}    | Update book by ID    |
| DELETE | /books/{id}    | Delete book by ID    |

## Example Book Model

```go
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}
```

## License
MIT
