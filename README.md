The Credit Card Validator System is a web application that validates credit card numbers using the Luhn algorithm. The system consists of a backend service implemented in Go and a frontend user interface using HTML, CSS, and JavaScript. This README provides an overview of the project's components, setup instructions, and how to run the application.

Tech Stack
Frontend: HTML, CSS, JavaScript
Backend: Go (Golang)
Algorithm: Luhn Algorithm for credit card validation
Features
Credit Card Validation: Validates credit card numbers using the Luhn algorithm.
Responsive Design: Mobile-friendly interface with warm background colors and smooth transitions.
User Interface: Two pages with transitions, validation messages, and formatted input for credit card numbers.

Backend Implementation
The backend of the Credit Card Validator System is implemented in Go and performs credit card validation using the Luhn algorithm. It handles HTTP GET requests with JSON payloads, processes the credit card number, and returns validation results.

Implementation Steps
Luhn Algorithm: Validates credit card numbers to ensure they meet the Luhn checksum requirement.

HTTP Server: An HTTP server is created using Go’s net/http package.

Request Handling: The server is configured to handle GET requests with JSON payloads.

JSON Payload Handling:

Validation: Accepts valid JSON requests; rejects invalid ones with a 400 status code.
Extraction: Extracts the credit card number from the JSON payload.
Validation and Response:

Processing: Runs the Luhn algorithm on the extracted credit card number.
Response: Wraps the result in a JSON response payload and returns it to the user.
