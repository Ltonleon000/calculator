# Calculator Service

A simple and efficient HTTP service that evaluates arithmetic expressions. The service accepts mathematical expressions through a REST API endpoint and returns the calculated result.

## Features

- Evaluates basic arithmetic expressions
- RESTful API interface
- Supports addition (+) and subtraction (-) operations
- Handles negative numbers
- Comprehensive error handling

## API Specification

### Calculate Endpoint

```
POST /api/v1/calculate
Content-Type: application/json
```

Request body:
```json
{
    "expression": "string"
}
```

Response body (success):
```json
{
    "result": number
}
```

Response body (error):
```json
{
    "error": "string"
}
```

## Getting Started

### Prerequisites

- Go 1.x or higher

### Running the Service

To start the service, run the following command from the project root:

```bash
go run main.go
```

The service will start on port 8080.

## Usage Examples

Here are examples demonstrating different scenarios:

### 1. Successful Calculation

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "10-5+3"
}'
```

Response:
```json
{
    "result": 8
}
```

### 2. Invalid Expression (422 Error)

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+a"
}'
```

Response:
```json
{
    "error": "Expression is not valid"
}
```

### 3. Server Error (500)

```bash
curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": ""
}'
```

Response:
```json
{
    "error": "Internal server error"
}
```

## Error Handling

The service handles various types of errors:

- 422 Unprocessable Entity: Invalid expression format, invalid characters, or syntax errors
- 500 Internal Server Error: Unexpected server-side errors

## Testing

To run the tests:

```bash
go test ./...
```

## Project Structure

```
calc_service/
├── main.go
├── internal/
│   ├── calculator/
│   │   ├── calculator.go
│   │   └── calculator_test.go
│   └── handler/
│       ├── handler.go
│       └── handler_test.go
└── README.md
