# go-healthcare-api

A simple Go-based API for managing patient data.

## Source Code

The source code for this project is hosted on GitHub at:
https://github.com/Falcs1/go-healthcare-api.git

## Getting Started

These instructions will get you a copy of the project up and running on your local machine.

### Prerequisites

* Go (version 1.16 or higher recommended)
* A compatible database (details depending on your implementation - e.g., SQLite, PostgreSQL, MySQL)

### Setup Instructions

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/Falcs1/go-healthcare-api.git](https://github.com/Falcs1/go-healthcare-api.git)
    cd go-healthcare-api
    ```

2.  **Install dependencies:**
    Navigate to the project directory and run:
    ```bash
    go mod download
    ```

3.  **Configure your database:**
    * Ensure your database is running.
    * Update the database connection string or configuration in the application (refer to `[mention configuration file or environment variable]` if applicable - *You might want to add a note here about how database connection is handled, e.g., environment variables, config file, hardcoded for demo*).
    * Run any necessary database migrations (if your project uses them, e.g., `go run main.go migrate`).

4.  **Run the application:**
    Execute the following command from the project's root directory:
    ```bash
    go run .
    ```
    The server should start and listen on the port specified in the application configuration (defaulting to `8081` as per the API documentation).

5.  **Access the API:**
    The API will be available at `http://localhost:8081` (or the configured port).

## API Endpoints

This section provides details on the available API endpoints for interacting with the system. All requests and responses are formatted as JSON.

**Base URL:** `http://localhost:8081`

### `POST /patients`

* **Description:** Creates a new patient record.
* **Method:** `POST`
* **Path:** `/patients`
* **Request Body (JSON):**
    Requires the following fields:
    * `name` (string): The full name of the patient.
    * `email` (string): The email address of the patient.
    * `phone` (string): The phone number of the patient.

    ```json
    {
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "123-456-7890"
    }
    ```
* **Example cURL Request:**

    ```bash
    curl -X POST http://localhost:8081/patients \
    -H "Content-Type: application/json" \
    -d '{"name": "John Doe", "email": "john.doe@example.com", "phone": "123-456-7890"}'
    ```
* **Success Response (`201 Created`):**
    Returns the details of the newly created patient, including its generated ID and timestamps.

    ```json
    {
      "ID": 1,
      "CreatedAt": "2023-10-27T10:00:00Z",
      "UpdatedAt": "2023-10-27T10:00:00Z",
      "DeletedAt": null,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "123-456-7890",
      "Appointments": null
    }
    ```
* **Error Responses:**
    * `400 Bad Request`: If the request body JSON is invalid or missing required fields (`name`, `email`, `phone`).
    * `500 Internal Server Error`: If an unexpected error occurs on the server, such as a database issue.

### `GET /patients`

* **Description:** Retrieves a list of all patient records.
* **Method:** `GET`
* **Path:** `/patients`
* **Example cURL Request:**

    ```bash
    curl http://localhost:8081/patients
    ```
* **Success Response (`200 OK`):**
    Returns an array of patient objects. The structure of each object is the same as the success response for `POST /patients`.

    ```json
    [
      {
        "ID": 1,
        "CreatedAt": "2023-10-27T10:00:00Z",
        "UpdatedAt": "2023-10-27T10:00:00Z",
        "DeletedAt": null,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "phone": "123-456-7890",
        "Appointments": null
      },
      {
        "ID": 2,
        "CreatedAt": "2023-10-27T10:05:00Z",
        "UpdatedAt": "2023-10-27T10:05:00Z",
        "DeletedAt": null,
        "name": "Alice Smith",
        "email": "alice.s@example.com",
        "phone": "987-654-3210",
        "Appointments": null
      }
    ]
    ```
* **Error Responses:**
    * `500 Internal Server Error`: If an unexpected error occurs on the server, such as a database issue.

### `GET /patients/:id`

* **Description:** Retrieves a specific patient record by its unique ID.
* **Method:** `GET`
* **Path:** `/patients/:id`
* **URL Parameters:**
    * `:id` (integer, required): The unique identifier of the patient to retrieve.

* **Example cURL Request (for patient with ID 1):**

    ```bash
    curl http://localhost:8081/patients/1
    ```
* **Success Response (`200 OK`):**
    Returns the details of the requested patient. The structure is the same as a single object in the `GET /patients` list and the `POST /patients` success response.

    ```json
    {
      "ID": 1,
      "CreatedAt": "
