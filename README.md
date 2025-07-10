# Go Echo Starter

A starter project using Echo framework and Swagger documentation.

## Project Structure

```
go-echo-starter/
├── cmd/
│   └── server/
│       └── main.go
├── configs/
├── internal/
│   ├── app/
│   │   ├── handlers/     # HTTP handlers
│   │   ├── repositories/ # Data repositories
│   │   └── services/     # Business logic
│   ├── domain/          # Domain models and DTOs
│   └── infra/           # Infrastructure code
├── shared/              # Shared utilities
└── docs/               # Swagger documentation
```

## Getting Started

### Prerequisites
- Go 1.23.5 or higher
- Make

### Installation

1. Clone the repository:
```bash
git clone https://github.com/vnurhaqiqi/go-echo-starter.git
cd go-echo-starter
```

2. Install dependencies:
```bash
go mod tidy
```

### Running the Project

1. Start the server:
```bash
make run
```

2. Generate Wire dependencies:
```bash
make wire
```

3. Generate Swagger documentation:
```bash
make swagger
```

## API Documentation

Access the Swagger UI at: `http://localhost:8080/swagger`

### Available Endpoints

#### Customers API

- **GET /customers**
  - Get all customers
  - Response: `[]CustomerResponse`

- **GET /customers/{id}**
  - Get customer by ID
  - Parameters:
    - `id`: Customer ID (UUID)
  - Response: `CustomerResponse`

- **POST /customers**
  - Create new customer
  - Request body: `CustomerRequest`
  - Response: `CustomerResponse`

- **PUT /customers/{id}**
  - Update customer
  - Parameters:
    - `id`: Customer ID (UUID)
  - Request body: `CustomerRequest`
  - Response: `CustomerResponse`


## Project Details

- **Framework**: Echo v4
- **Documentation**: Swagger/OpenAPI
- **Validation**: go-playground/validator
- **Logging**: zerolog
- **Dependency Injection**: Wire

## Development

### Adding New Endpoints

1. Create new handler in `internal/app/handlers/`
2. Define request/response DTOs in `internal/domain/dto/`
3. Implement business logic in `internal/app/services/`
4. Add Swagger documentation tags
5. Run `make swagger` to update documentation

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
