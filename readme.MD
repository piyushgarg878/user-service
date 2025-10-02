# User Service

A simple Go-based microservice for managing user registration, authentication, and profiles. This service is built using Gin for routing and GORM for database interactions.

## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Service](#running-the-service)
- [API Endpoints](#api-endpoints)
- [Project Structure](#project-structure)

## Features

-   User registration with password hashing.
-   User login (logic implemented, requires JWT for full functionality).
-   Create, retrieve, and update user profiles.
-   Layered architecture (Handler, Service, Repository) for clean separation of concerns.

## Prerequisites

-   [Go](https://golang.org/doc/install) (version 1.18 or higher)
-   A running PostgreSQL instance (or another SQL database compatible with GORM).

## Installation

1.  **Clone the repository:**
    ```sh
    git clone <your-repository-url>
    cd user-service
    ```

2.  **Install dependencies:**
    This command will download all the necessary Go modules defined in `go.mod`.
    ```sh
    go mod tidy
    ```

3.  **Configure Environment:**
    The application is configured to connect to a database via a connection string in `cmd/main.go`. For production use, it's recommended to use environment variables.

    Update the DSN (Data Source Name) in `cmd/main.go` to point to your database:
    ```go
    // cmd/main.go
    dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
    ```

## Running the Service

Once the installation and configuration are complete, you can run the service with the following command:

```sh
go run cmd/main.go
```

The service will start and listen on port `8080` by default.

## API Endpoints

The following are the available API endpoints for interacting with the service.

---

### 1. Register a New User

-   **Endpoint:** `POST /signup`
-   **Description:** Creates a new user account. The password is automatically hashed before being stored.
-   **Request Body:**

    ```json
    {
        "name": "John Doe",
        "email": "john.doe@example.com",
        "password": "a_strong_password"
    }
    ```

-   **Example `curl`:**

    ```sh
    curl -X POST http://localhost:8080/signup \
    -H "Content-Type: application/json" \
    -d '{
        "name": "John Doe",
        "email": "john.doe@example.com",
        "password": "a_strong_password"
    }'
    ```

-   **Success Response (201 Created):**

    ```json
    {
        "message": "user created",
        "user_id": 1
    }
    ```

---

### 2. Get User by Name

-   **Endpoint:** `GET /users/user/:name`
-   **Description:** Retrieves the details of a single user by their name.
-   **Example `curl`:**

    ```sh
    curl -X GET http://localhost:8080/users/user/John%20Doe
    ```

-   **Success Response (200 OK):**

    ```json
    {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "profile": {
            "id": 1,
            "user_id": 1,
            "bio": "",
            "avatar": ""
        },
        "created_at": "2025-10-02T10:00:00Z",
        "updated_at": "2025-10-02T10:00:00Z"
    }
    ```

---

### 3. Get User Profile

-   **Endpoint:** `GET /users/profile/:id`
-   **Description:** Retrieves the profile information for a specific user ID.
-   **Example `curl`:**

    ```sh
    curl -X GET http://localhost:8080/users/profile/1
    ```

-   **Success Response (200 OK):**

    ```json
    {
        "id": 1,
        "user_id": 1,
        "bio": "Software developer.",
        "avatar": "http://example.com/avatar.png"
    }
    ```

---

### 4. Update User Profile

-   **Endpoint:** `PUT /users/updateprofile/:id`
-   **Description:** Updates the profile information (bio, avatar) for a specific user.
-   **Request Body:**

    ```json
    {
        "bio": "Senior Software Engineer",
        "avatar": "http://example.com/new_avatar.png"
    }
    ```

-   **Example `curl`:**

    ```sh
    curl -X PUT http://localhost:8080/users/updateprofile/1 \
    -H "Content-Type: application/json" \
    -d '{
        "bio": "Senior Software Engineer",
        "avatar": "http://example.com/new_avatar.png"
    }'
    ```

-   **Success Response (200 OK):**

    ```json
    {
        "message": "profile updated"
    }
    ```

## Project Structure

The project follows a standard layered architecture to separate concerns:

-   `cmd/`: Main application entry point.
-   `internal/`: Contains all the core application logic.
    -   `handler/`: Handles HTTP requests and responses.
    -   `service/`: Contains the business logic.
    -   `repository/`: Manages data access and database interactions.
    -   `domain/`: Defines the core data structures (models).
    -   `router/`: Configures the API routes.
-   `pkg/`: Shared libraries and utilities (e.g., password hashing).
