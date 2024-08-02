# Go Backend Boilerplate

This is a boilerplate for a Go backend application. It includes a basic structure for a Go application and a Dockerfile.

The purpose of this boilerplate is to provide a starting point for a Go backend application that mostly handle HTTP requests and responses, some background jobs - very usual backend app usecase.

## Directory Structure

```
.
├── Dockerfile
├── README.md
├── auth -> domain layer
│   ├── auth.go
│   └── auth_test.go
├── bin -> binary files
│   └── web -> web binary, for web server
|       └── main.go 
│   └── worker -> worker binary, for background job
|       └── main.go
├── config -> configuration files
│   ├── config.go
├── docs -> documentation
├── entity -> entity layer
│   ├── user.go
│   └── error.go
├── http -> delivery layer
│   ├── handler -> handler layer
│   │   ├── admin -> admin handler, for admin API
│   │   │   ├── handler.go
│   │   │   ├── route.go
│   │   ├── user -> user handler, for user API
│   │   │   ├── handler.go
│   │   │   ├── route.go
│   ├── middleware -> middleware layer
│   │   ├── auth.go
│   │   ├── cors.go
│   │   ├── logger.go
│   │   ├── recover.go
├── internal -> internal package
|   ├── db
├── migration -> migration files
├── role -> domain layer


