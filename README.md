Credit Card Validator System
Overview
This project implements a Credit Card Validator System with a frontend interface built using HTML, CSS, and JavaScript, and a backend server written in Go. The system validates credit card numbers using the Luhn algorithm and provides immediate feedback on the card's validity.

Technology Stack
Frontend: HTML, CSS, JavaScript
Backend: Go
Validation Algorithm: Luhn Algorithm
Data Format: JSON
HTTP Server: Go’s net/http package
Features
Credit Card Validation: Uses the Luhn algorithm to validate credit card numbers.
Responsive UI: User-friendly interface with a warm gradient background and smooth transitions.
Dynamic Validation: Provides real-time feedback on card validity.
Backend Implementation
Luhn Algorithm: Validates credit card numbers to ensure they meet the Luhn checksum requirement.

HTTP Server: Created using Go’s net/http package.

Request Handling: Configured to handle GET requests with JSON payloads.

JSON Payload Handling:

Validation: Accepts valid JSON requests; rejects invalid ones with an HTTP 400 status code.
Extraction: Extracts the credit card number from the JSON payload.
Validation and Response:

Processing: Runs the Luhn algorithm on the extracted credit card number.
Response: Wraps the result in a JSON response payload and returns it to the user.
