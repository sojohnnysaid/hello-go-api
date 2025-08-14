# Hello Go API

A bare-bones Go API service with automatic OpenAPI/Swagger generation using [Huma](https://github.com/danielgtaylor/huma).

## Features

- Automatic OpenAPI 3.1 specification generation
- Interactive Swagger UI documentation
- Type-safe request/response handling
- Built-in validation using struct tags
- Example values and documentation in the OpenAPI spec

## Requirements

- Go 1.23 or newer

## Installation

```bash
go mod tidy
```

## Running the API

```bash
go run main.go
```

The server will start on `http://127.0.0.1:8888`

## API Endpoints

### Greeting Endpoint

**GET** `/greeting/{name}`

Get a personalized greeting message.

Example:
```bash
curl http://127.0.0.1:8888/greeting/world
```

Response:
```json
{
  "message": "Hello, world!"
}
```

## Documentation

Once the server is running, you can access:

- **Interactive API Documentation**: http://127.0.0.1:8888/docs
- **OpenAPI JSON**: http://127.0.0.1:8888/openapi.json
- **OpenAPI YAML**: http://127.0.0.1:8888/openapi.yaml

## Project Structure

```
hello-go-api/
├── go.mod          # Go module file
├── go.sum          # Go dependencies lock file
├── main.go         # Main application file
└── README.md       # This file
```

## How It Works

The API uses Huma framework which:
1. Automatically generates OpenAPI documentation from Go structs
2. Provides type-safe handlers with automatic validation
3. Serves interactive Swagger UI at `/docs`
4. Supports multiple content types (JSON, CBOR, etc.)

## Extending the API

To add new endpoints:

1. Define input/output structs with appropriate tags
2. Register the operation with `huma.Register()`
3. Implement the handler function

Example tags:
- `json:"fieldname"` - JSON field name
- `path:"name"` - Path parameter
- `query:"name"` - Query parameter
- `header:"name"` - Header parameter
- `example:"value"` - Example value for docs
- `doc:"description"` - Field description
- `maxLength:"30"` - Validation constraint