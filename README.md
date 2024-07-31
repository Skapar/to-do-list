# Todo List Microservice

This is a simple Todo List microservice.

## Requirements

- Docker
- Docker Compose
- Go 1.21.6

## Installation

1. Clone the repository:

    ```sh
    git clone <github.com/Skapar/to-do-list>
    cd todo-list
    ```

2. Build and run the Docker containers:

    ```sh
    make build
    make run
    ```

## Usage

### Create a New Task

- **Endpoint:** `POST /api/todo-list/tasks`
- **Request Body:**

    ```json
    {
      "title": "Buy a book",
      "activeAt": "2023-08-04"
    }
    ```

### Update a Task

- **Endpoint:** `PUT /api/todo-list/tasks/{id}`
- **Request Body:**

    ```json
    {
      "title": "Buy a high-load application book",
      "activeAt": "2023-08-05"
    }
    ```

### Delete a Task

- **Endpoint:** `DELETE /api/todo-list/tasks/{id}`

### Mark Task as Done

- **Endpoint:** `PUT /api/todo-list/tasks/{id}/done`

### List Tasks by Status

- **Endpoint:** `GET /api/todo-list/tasks?status=active` or `GET /api/todo-list/tasks?status=done`
