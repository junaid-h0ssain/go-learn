# go-learn

A Go-based API server for managing user coin balances with authentication. This project demonstrates building a RESTful API using Go, including routing, middleware, and database interactions.

## Tech Stack

- **Language**: Go 1.25.6
- **Web Framework**: Chi (lightweight router)
- **Logging**: Logrus
- **Database**: Mock database implementation (for demonstration purposes)
- **Dependencies**:
  - github.com/go-chi/chi v1.5.5
  - github.com/sirupsen/logrus v1.9.4
  - rsc.io/quote v1.5.2

## Features

- User authentication via authorization tokens
- Get user coin balance
- Transfer coins between users (endpoint defined, implementation pending)
- Middleware for request authorization
- JSON API responses
- Error handling with appropriate HTTP status codes

## Project Structure

```
go-learn/
├── main.go                 # Application entry point
├── go.mod                  # Go module definition
├── api/
│   └── api.go              # API types and error handlers
├── handlers/
│   ├── handler.go          # Route definitions
│   └── get_coin_balance.go # Coin balance handler
├── middleware/
│   └── authorization.go    # Authentication middleware
└── tools/
    ├── database.go         # Database interface
    └── mockdb.go           # Mock database implementation
```

## Installation

1. Ensure you have Go 1.25.6 or later installed.
2. Clone or download the project.
3. Navigate to the project directory.
4. Run `go mod tidy` to download dependencies.

## Usage

1. Start the server:
   ```
   go run main.go
   ```

2. The server will start on `localhost:8080`.

### API Endpoints

#### Get Coin Balance
- **URL**: `/account/`
- **Method**: GET
- **Query Parameters**:
  - `username`: The username to get balance for
- **Headers**:
  - `Authorization`: User's auth token
- **Response**:
  ```json
  {
    "Code": 200,
    "Balance": 100.0
  }
  ```

#### Transfer Coins
- **URL**: `/account/transfer`
- **Method**: POST
- **Headers**:
  - `Authorization`: User's auth token
- **Note**: Implementation is pending (marked as TODO)

### Authentication

All `/account` routes require authentication:
- Provide `username` as a query parameter
- Include `Authorization` header with the user's token

Mock users and tokens (for testing):
- Username: alex, Token: 123ABC
- Username: jason, Token: 456DEF
- Username: marie, Token: 789GHI

## Development

This project includes additional Go learning examples in separate files:
- `hello.go`: Basic Go syntax and types
- `channel.go`: Goroutine and channel examples
- `map.go`: Map data structure usage

## Credit

Inspired by [This Video](https://youtu.be/8uiZC0l4Ajw)