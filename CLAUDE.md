# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`dh-auth` is a Go-based authentication service for DutyHub. The project is in early development stages with a minimal structure focused on database connectivity and configuration management.

## Architecture

### Configuration Layer (`config/`)
- Uses environment-based configuration via `config.Load()`
- Configuration is validated before use via `config.Validate()`
- Database DSN generation: `DatabaseConfig.DSN()` uses correct PostgreSQL format: `host=value port=value user=value password=value dbname=value sslmode=disable`
- All environment variables have defaults defined in `config/config.go`
- **Security Note**: SSL is currently hardcoded to `disable` - should be made configurable for production use

### Repository Layer (`internal/repository/`)
- PostgreSQL integration located in `internal/repository/postgres/`
- Uses `pgxpool` for connection pooling with the following defaults:
  - MaxConns: 5
  - MinConns: 1
  - MaxConnLifetime: 30 minutes
  - MaxConnIdleTime: 10 minutes
  - HealthCheckPeriod: 1 minute
- Database connection details configured via environment variables (see `.env.example`)
- Connection pool is initialized and pinged during `NewPool()` creation

### Entry Point (`cmd/main.go`)
- Currently contains placeholder/template code from GoLand IDE
- Will need to be replaced with actual application initialization

## Environment Configuration

Required environment variables (with defaults):
- `SERVER_PORT` (default: 8080)
- `DB_HOST` (default: localhost)
- `DB_PORT` (default: 5432)
- `DB_USER` (default: postgres)
- `DB_PASS` (default: postgres) - **Required, validated**
- `DB_NAME` (default: postgres in code, but .env.example shows db_dutyhub) - **Required, validated**

Configuration is loaded via `config.Load()` which returns an error if validation fails.

**Note**: There's an inconsistency between the code default ("postgres") and .env.example ("db_dutyhub") for DB_NAME.

## Development Commands

### Running the application
```bash
go run cmd/main.go
```

### Building the application
```bash
go build -o bin/dh-auth cmd/main.go
```

### Running tests
```bash
go test ./...
```

### Running a single test
```bash
go test ./path/to/package -run TestName
```

## Known Issues

### Critical
- **config/config.go:63** - SSL mode hardcoded to `disable`; should be configurable via environment variable
- **internal/repository/postgres/db.go:19** - DSN may be exposed in error messages, potentially leaking credentials
- **cmd/main.go** - Contains only IDE template code; needs proper implementation with graceful shutdown

### Medium Priority
- **config/config.go:34-36** - Validation checks DB_HOST for empty value, but getEnv() always provides "localhost" default, making the check ineffective
- **config/config.go:52-53** - Default value mismatch between code and .env.example for DB_NAME
- **internal/repository/postgres/db.go** - Connection pool settings are hardcoded (MaxConns: 5 may be too small for production)
- No test coverage exists for any package

### Code Quality
- **internal/repository/postgres/db.go:19** - Uses `errors.New(fmt.Sprintf(...))` instead of `fmt.Errorf`
- Error messages don't follow Go conventions (should be lowercase, no punctuation)
- **internal/repository/postgres/db.go:13** - `Database.Pool` field is exported, violating encapsulation

## Important Notes

- The project uses Go 1.25
- No external dependencies are currently declared in go.mod (pgxpool needs to be added to go.mod)
- The main.go file contains IDE template code and needs implementation
- Database configuration validation requires `DB_PASS`, `DB_HOST`, and `DB_NAME` to be set
- For production use, ensure SSL is enabled and connection pool settings are tuned appropriately
