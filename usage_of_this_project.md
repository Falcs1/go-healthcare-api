API Endpoints

The API is hosted at http://localhost:8081. All requests and responses use JSON format.
Patient Endpoints

POST /patients

    Description: Creates a new patient.
    Method: POST
    Path: /patients
    Request Body (JSON):
        Requires: name (string), email (string), phone (string) <!-- end list -->
    JSON

{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "123-456-7890"
}

Example cURL Request:
Bash

curl -X POST http://localhost:8081/patients \
-H "Content-Type: application/json" \
-d '{"name": "John Doe", "email": "john.doe@example.com", "phone": "123-456-7890"}'

Success Response (201 Created):
JSON

    {
      "ID": 1,
      "CreatedAt": "...",
      "UpdatedAt": "...",
      "DeletedAt": null,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "phone": "123-456-7890",
      "Appointments": null
    }

    Error Responses:
        400 Bad Request: If JSON is invalid or required fields are missing.
        500 Internal Server Error: If a database error occurs.

GET /patients

    Description: Retrieves a list of all patients.
    Method: GET
    Path: /patients
    Example cURL Request:
    Bash

curl http://localhost:8081/patients

Success Response (200 OK):
JSON

    [
      {
        "ID": 1,
        "CreatedAt": "...",
        "UpdatedAt": "...",
        "DeletedAt": null,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "phone": "123-456-7890",
        "Appointments": null
      },
      {
        "ID": 2,
        "CreatedAt": "...",
        "UpdatedAt": "...",
        "DeletedAt": null,
        "name": "Alice Smith",
        "email": "alice.s@example.com",
        "phone": "987-654-3210",
        "Appointments": null
      }
    ]

    Error Responses:
        500 Internal Server Error: If a database error occurs.

GET /patients/:id

    Description: Retrieves a specific patient by ID.
    Method: GET
    Path: /patients/:id
    URL Parameters:
        :id (integer): The ID of the patient to retrieve.
    Example cURL Request (for patient ID 1):
    Bash

curl http://localhost:8081/patients/1

Success Response (200 OK): (Same as the object structure in the GET /patients list)
JSON

{
  "ID": 1,
  "CreatedAt": "...",
  "UpdatedAt": "...",
  "DeletedAt": null,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "phone": "123-456-7890",
  "Appointments": null
}

Error Responses:

    400 Bad Request: If the ID format is invalid (e.g., /patients/abc).
    404 Not Found: If a patient with the given ID does not exist.
    500 Internal Server Error: If a database error occurs during lookup.
