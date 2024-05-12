# GoSQL - CRUD API USING GORM AND POSTGRES

GoSQL is a simple CRUD application built in Go using Gorilla Mux and GORM, which provides a RESTful API for managing books.

## Features

- Retrieve a list of movies
- Retrieve a single movie by ID
- Create a new movie
- Update an existing movie
- Delete a movie by ID

## Prerequisites

Make sure you have Go installed on your machine. You can download and install it from [here](https://golang.org/dl/).

## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux): A powerful HTTP router and URL matcher for building Go web servers with routing capabilities.
- [GORM](https://gorm.io/): An Object Relational Mapping library for Go. It provides a simple and efficient way to work with SQL databases.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Yash-sudo-web/gosql_bookmanagement.git

2. Navigate to the project directory:

   ```bash
   cd cmd/main

3. Build and run the application:

   ```bash
   go build . && ./main

## API Endpoints
- GET /book/{id}: Get details of a book by ID.
- GET /book/: Get details of all books.
- POST /book/: Create a new book.
- PUT /book/{id}: Update details of a book by ID.
- DELETE /book/{id}: Delete a book by ID.

## Usage
You can use tools like cURL or Postman to interact with the API endpoints. Here are some examples:
- Get all books:

  ```bash
  curl http://localhost:8000/book/
  
- Get a book by ID:
  
  ```bash
  curl http://localhost:8000/book/1

- Create a new book:

  ```bash
  curl -X POST -H "Content-Type: application/json" -d '{ "name": "The Great Gatsby", "author": "F. Scott Fitzgerald", "publication": "Scribner" }' http://localhost:8000/book/

- Update an existing movie:

  ```bash
  curl -X PUT -H "Content-Type: application/json" -d '{ "name": "The Great Gatsby", "author": "F. Scott Fitzgerald", "publication": "Scribner" }' http://localhost:8000/book/1

- Delete a movie by ID:

  ```bash
  curl -X DELETE http://localhost:8000/book/1

## Contributing
Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the [MIT License](https://github.com/Yash-sudo-web/gosql_bookmanagement/blob/main/LICENSE).
