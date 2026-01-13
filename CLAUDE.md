# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Swan is a Go web development boilerplate/framework for building web applications with a modern stack. The repository contains:

- **swctl**: CLI tool for scaffolding and code generation
- **bootstrap**: Template application demonstrating the framework's capabilities
- **utl**: Shared utilities

## Key Technologies

- **Web Framework**: Echo v4
- **Dependency Injection**: Uber Fx
- **ORM**: GORM (primary), Ent (optional)
- **Migrations**: Atlas + golang-migrate
- **API Documentation**: Swag (Swagger)
- **Pub/Sub**: Watermill
- **CLI**: Cobra
- **Frontend**: React + Vite + TanStack Router

## Development Commands

### swctl Installation
```bash
# Install swctl CLI tool
go install github.com/ugabiga/swan/swctl@latest

# Or build locally from swctl directory
cd swctl
make install
```

### Project Initialization
```bash
# Create new project
swctl new <project-name>
```

### Code Generation

All generation commands create files in the specified directory and automatically register them in `config/app/container.go`.

```bash
# Generate HTTP handler
swctl make:handler [folder-path] [api-prefix] [endpoint-name]
# Example: swctl make:handler todos /api/v1 todos

# Generate Cobra command
swctl make:command [folder-path]
# Example: swctl make:command todos

# Generate event handler
swctl make:event [folder-path]
# Example: swctl make:event todos

# Generate struct with DI registration
swctl make:struct [folder-path] [struct-name]
# Example: swctl make:struct todos Todo
```

### Bootstrap Application (Development)

```bash
cd bootstrap

# Run with live reload (requires air)
air

# Build and run manually
swag fmt . && swag init -g cmd/app/main.go && cd web && pnpm run build && cd .. && go build -o ./main ./cmd/app/main.go
./main server

# Run server command directly
go run cmd/app/main.go server [--port 8080]

# View all registered routes
go run cmd/app/main.go routes
```

### Database Migrations

```bash
cd bootstrap

# Create new migration (requires atlas CLI)
go run cmd/app/main.go migrate create <migration-name>

# Run migrations
go run cmd/app/main.go migrate up

# Rollback migrations
go run cmd/app/main.go migrate down
```

### Frontend Development

```bash
cd bootstrap/web

# Install dependencies
pnpm install

# Development server
pnpm dev  # runs on port 3001

# Build frontend
pnpm build

# Generate API client from OpenAPI spec
pnpm api  # or: orval
```

## Architecture

### Dependency Injection with Fx

The application uses Uber Fx for dependency injection. All DI configuration is centralized in `bootstrap/internal/app/container.go`:

- **provide()**: Registers providers (constructors) for services
- **invoke()**: Runs initialization functions that set up commands, routes, etc.
- **entry()**: Defines application entry points (command runner, event router)

When adding new services, register them in `container.go` in the appropriate `fx.Provide()` or `fx.Invoke()` block.

### Application Entry Points

Two main entry points exist:

1. **cmd/app/main.go**: Main application executable that starts the web server
2. **cmd/loader/main.go**: Atlas schema loader for GORM models (used during migrations)

### Server Architecture

Echo server setup happens in `bootstrap/internal/app/server/`:
- `server.go`: Server initialization and Cobra commands
- `router.go`: Route registration via `SetRouter()` function
- `openapi.go`: Swagger/OpenAPI endpoint handler
- `static.go`: Static file serving for frontend

Routes are registered in `SetRouter()` which is invoked by Fx. Add new route groups here.

### Database Support

The bootstrap supports multiple database backends through GORM:
- SQLite (default for development)
- PostgreSQL
- MySQL

Configure via `DATABASE_URL` environment variable. Atlas migrations support all three databases.

### ORM Setup

**GORM** (default): Active and configured in `bootstrap/internal/app/database/gormdb/`

**Ent** (optional): Currently commented out. To activate:
1. Uncomment `entdb.NewEntClient` in `container.go`
2. Uncomment code in `entdb/client.go`

### Migration Workflow

Migrations use Atlas CLI to generate golang-migrate files from GORM models:

1. Define models in your code
2. Add models to `cmd/loader/main.go` models slice
3. Run `go run cmd/app/main.go migrate create <name>`
4. Apply with `go run cmd/app/main.go migrate up`

Configuration is in `atlas.hcl` with two environments:
- `gorm`: Uses GORM schema loader (recommended)
- `local`: Uses Ent schema (when using Ent)

### Frontend Integration

The frontend lives in `bootstrap/web/` and uses:
- **Vite** for bundling
- **TanStack Router** for routing
- **TanStack Query** for data fetching
- **Orval** to generate API client from OpenAPI spec
- **Radix UI** for components
- **Tailwind CSS** for styling

The build output is embedded in Go and served via `StaticHandler` at runtime.

### Configuration

Configuration is environment-based using `godotenv`. Copy `.example.env` to `.env`:

- `DATABASE_URL`: Database connection string
- `DEV_DATABASE_URL`: Development/testing database
- OAuth settings (Google OAuth support included)
- `EVENT_ENABLED`: Enable/disable event pub/sub system

### Event System

Watermill-based pub/sub is available but disabled by default (`EVENT_ENABLED=false`). When enabled:
- Events are defined in `bootstrap/internal/app/event/`
- Event router runs in background via `RunEventRouter()` in container entry
- Use `EventEmitter` to publish events

### Testing

Test setup uses `NewTestApp()` which loads `.test.env` and provides access to test dependencies (Logger, DB) via Fx.

## Development Workflow

1. Generate code using `swctl make:*` commands
2. Implement business logic in generated files
3. Update `container.go` if adding services (usually done automatically)
4. Create GORM models and register in `cmd/loader/main.go`
5. Generate and run migrations
6. Update frontend in `web/` directory
7. Run with `air` for hot reload during development

## Important Notes

- When adding new HTTP handlers, they're automatically registered through Fx invocation
- Swagger docs regenerate on each build via `swag init`
- Frontend assets are built and embedded, no separate frontend server needed in production
- The proxy in `.air.toml` runs app on 8080, proxies frontend on 3000 in development
