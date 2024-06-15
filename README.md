# FealtyX - GoLang Assignment

## Installation

1. **Clone the repository**:
    ```sh
    git clone https://github.com/ashuthe1/Golang-with-Integration.git
    cd Golang-with-Integration
    ```

2. **Initialize a Go module**:
    ```sh
    go mod init FealtyX
    go mod tidy
    ```

3. **Run the application**:
    ```sh
    go run main.go
    ```

## Project Structure

```
.
├── handlers
│   └── handlers.go
├── models
│   └── models.go
├── ollama
│   └── ollama.go
├── main.go
└── README.md
```

## Features

- **CRUD Operations**: 
  - Create a new student
  - Get all students
  - Get a student by ID
  - Update a student by ID
  - Delete a student by ID

- **AI Integration**:
  - Generate a summary of a student by ID using the LLAMA3 and Google PaLM 2 API.


- **Concurrency**: 
  - Ensures the API can handle concurrent requests safely.

- **Error Handling**: 
  - Appropriately handles errors such as invalid input data and non-existing students.

- **Input Validation**: 
  - Ensures input data for creating and updating students is valid.

## API Endpoints

### Create a New Student

- **Endpoint**: `POST /students`
- **Description**: Creates a new student.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "age": 20,
    "email": "john.doe@example.com"
  }
  ```
- **cURL Command**:
  ```sh
  curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "age": 20, "email": "john.doe@example.com"}'
  ```

### Get All Students

- **Endpoint**: `GET /students`
- **Description**: Retrieves a list of all students.
- **cURL Command**:
  ```sh
  curl -X GET http://localhost:8080/students
  ```

### Get a Student by ID

- **Endpoint**: `GET /students/{id}`
- **Description**: Retrieves a student by their ID.
- **cURL Command**:
  ```sh
  curl -X GET http://localhost:8080/students/1
  ```

### Update a Student by ID

- **Endpoint**: `PUT /students/{id}`
- **Description**: Updates a student by their ID.
- **Request Body**:
  ```json
  {
    "name": "Jane Doe",
    "age": 22,
    "email": "jane.doe@example.com"
  }
  ```
- **cURL Command**:
  ```sh
  curl -X PUT http://localhost:8080/students/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "age": 22, "email": "jane.doe@example.com"}'
  ```

### Delete a Student by ID

- **Endpoint**: `DELETE /students/{id}`
- **Description**: Deletes a student by their ID.
- **cURL Command**:
  ```sh
  curl -X DELETE http://localhost:8080/students/1
  ```

### Generate a Summary of a Student by ID

- **Endpoint**: `GET /students/{id}/summary`
- **Description**: Generates a summary of a student profile by ID using the Google PaLM 2 API.
- **cURL Command**:
  ```sh
  curl -X GET http://localhost:8080/students/1/summary
  ```

## Concurrency

The API uses a mutex (`sync.Mutex`) to handle concurrent requests safely. This ensures data integrity when multiple requests are made simultaneously.

## Error Handling

Errors are appropriately handled, such as:
- Returning a 400 status code for invalid input data.
- Returning a 404 status code if a student with a specified ID does not exist.
- Returning a 500 status code for internal server errors.

## Input Validation

Input data for creating and updating students is validated to ensure it meets the required format and constraints.

---

This completes the setup and provides the necessary details to use the FealtyX API. Happy coding!