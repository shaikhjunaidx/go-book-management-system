# Go Book Management System

This project is a Book Management System implemented in Go, using GORM for ORM and Gorilla Mux for routing. It demonstrates basic CRUD (Create, Read, Update, Delete) operations on a Book model, with data stored in a MySQL database. The project illustrates how to handle HTTP requests, interact with a database, and organize a Go application for scalability and maintainability.

## Overview

**Tech Stack**: Go, GORM, Gorilla Mux  
**Purpose**: Learn Go structs, slices, and basic HTTP handling.  
**Features**:
- Create a new book
- Retrieve all books
- Retrieve a book by ID
- Update a book
- Delete a book

## Project Structure

```
cmd/
└── main/
├── main.exe
└── main.go
pkg/
├── config/
│ └── app.go
├── controllers/
│ └── book-controller.go
├── models/
│ └── book.go
├── routes/
│ └── bookstore-routes.go
└── utils/
└── utils.go
go.mod
go.sum
```

## Getting Started

### Prerequisites

- Go 1.16 or later
- MySQL database

### Installation

1. Clone the repository:
   ``` bash
   git clone https://github.com/yourusername/go-book-management-system.git
   ```
2. Navigate to the project directory:
    ```bash
    cd go-book-management-system
    ```
3. Install dependencies:
    ```bash
    go mod download
    ```

### Configuration
Update the database connection details in pkg/config/app.go if necessary.

### Running the Application
1. Start the MySQL server and ensure the database is accessible.
2. Run the application:
    ```bash
        go run cmd/main/main.go
    ```

## Endpoints

- **Create a Book**
  - `POST /book`
- **Retrieve All Books**
  - `GET /books`
- **Retrieve a Book by ID**
  - `GET /book/{bookId}`
- **Update a Book**
  - `PUT /book/{bookId}`
- **Delete a Book**
  - `DELETE /book/{bookId}`

## Example Book JSON

```json
{
    "name": "1984",
    "author": "George Orwell",
    "publication": "Secker & Warburg"
}