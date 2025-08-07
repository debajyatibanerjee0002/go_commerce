# go_commerce

# phase 1 (api/v1) -> version_v1

- Product add -> POST version_v1/product
- Product read -> GET version_v1/products/:id
- Product read all -> GET version_v1/products
- Product edit -> PATCH version_v1/products/:id
- Product delete -> DELETE version_v1/products/:id

# folder structure

```
go-commerce/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
│
├── internal/
│   ├── api/
│   │   └── v1/
│   │       ├── handlers/
│   │       │   ├── product.go      # Product HTTP handlers
│   │       │   └── health.go       # Health check handler
│   │       ├── middleware/
│   │       │   ├── auth.go         # Authentication middleware
│   │       │   ├── cors.go         # CORS middleware
│   │       │   └── logging.go      # Request logging middleware
│   │       └── routes/
│   │           └── routes.go       # Route definitions
│   │
│   ├── domain/
│   │   ├── product/
│   │   │   ├── model.go           # Product domain models
│   │   │   ├── repository.go      # Product repository interface
│   │   │   └── service.go         # Product business logic
│   │   └── errors/
│   │       └── errors.go          # Custom error types
│   │
│   ├── infrastructure/
│   │   ├── database/
│   │   │   ├── postgres/
│   │   │   │   ├── connection.go  # Database connection
│   │   │   │   └── product_repo.go # Product repository implementation
│   │   │   └── migrations/
│   │   │       └── 001_create_products.sql
│   │   ├── config/
│   │   │   └── config.go          # Configuration management
│   │   └── logger/
│   │       └── logger.go          # Logging setup
│   │
│   └── pkg/
│       ├── utils/
│       │   ├── response.go        # Standard API response helpers
│       │   └── validation.go      # Input validation helpers
│       └── constants/
│           └── constants.go       # Application constants
│
├── api/
│   └── openapi/
│       └── swagger.yaml           # API documentation
│
├── scripts/
│   ├── migrate.sh                 # Database migration script
│   └── seed.sh                    # Database seeding script
│
├── deployments/
│   ├── docker/
│   │   └── Dockerfile
│   └── k8s/
│       ├── deployment.yaml
│       └── service.yaml
│
├── tests/
│   ├── integration/
│   │   └── product_test.go        # Integration tests
│   └── fixtures/
│       └── products.json          # Test data
│
├── .env.example                   # Environment variables template
├── .gitignore
├── docker-compose.yml             # Local development setup
├── go.mod
├── go.sum
├── Makefile                       # Build and development commands
└── README.md
```

The structure uses dependency injection and interfaces, making it easy to swap implementations (e.g., switching from PostgreSQL to MongoDB) without changing business logic. This approach is widely adopted in Go microservices and follows patterns used by companies like Uber, Netflix, and other tech companies using Go for e-commerce platforms.

# Initialize the project

go mod init go-commerce

# Install common dependencies

go get github.com/gin-gonic/gin
go get github.com/lib/pq # PostgreSQL driver
go get github.com/golang-migrate/migrate/v4
go get github.com/joho/godotenv # Environment variables
go get github.com/google/uuid # UUID generation

# Run the application

make run

# Run tests

make test

# Database migration

make migrate-up
