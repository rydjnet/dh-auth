# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`dh-auth` is a Go-based authentication service for DutyHub. The project is in early development stages with a minimal structure focused on database connectivity and configuration management.

## Architecture

### Configuration Layer (`config/`)
- Uses environment-based configuration via `config.Load()`
- Configuration is validated before use via `config.Validate()`
- Database DSN generation: `DatabaseConfig.DSN()` - Note: Current format uses `:` separator for host/port (format: `host:localhost port:5432...`)
- All environment variables have defaults defined in `config/config.go`

### Repository Layer (`internal/repository/`)
- PostgreSQL integration located in `internal/repository/postgres/`
- Database connection details configured via environment variables (see `.env.example`)

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
- `DB_NAME` (default: db_dutyhub) - **Required, validated**

Configuration is loaded via `config.Load()` which returns an error if validation fails.

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

## Important Notes

- The project uses Go 1.25
- No external dependencies are currently declared in go.mod
- The main.go file contains IDE template code and needs implementation
- Database configuration validation requires `DB_PASS`, `DB_HOST`, and `DB_NAME` to be set
