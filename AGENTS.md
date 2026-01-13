# AGENTS.md

Guidance for future coding work in this repository.

## Scope
- Repository: Swan (Go web development boilerplate/framework)
- Key components: `swctl`, `bootstrap`, `utl`

## Principles
- Keep changes minimal and consistent with existing architecture.
- Prefer clear, explicit registration in Fx containers.
- Favor generated code via `swctl make:*` when appropriate.

## Key Paths
- CLI tooling: `swctl/`
- Template app: `bootstrap/`
- Shared utilities: `utl/`

## Go Architecture Notes
- Dependency Injection: Uber Fx, configured in `bootstrap/internal/app/container.go`
- Web server: Echo v4
- ORM: GORM
- Migrations: Atlas + golang-migrate
- API Docs: Swag (Swagger)
- Events: Watermill (guarded by `EVENT_ENABLED`)

## Frontend Notes
- Location: `bootstrap/web/`
- Stack: React + Vite + TanStack Router/Query
- API client: Orval (`pnpm api`)
- Build artifacts are embedded and served by Go in production.

## Common Commands
```bash
# build/install CLI
cd swctl && make install

# run bootstrap app (dev)
cd bootstrap
air

# run server directly
go run cmd/app/main.go server

# list routes
go run cmd/app/main.go routes

# migrate
go run cmd/app/main.go migrate create <name>
go run cmd/app/main.go migrate up
```

## Code Generation
```bash
swctl make:handler [folder-path] [api-prefix] [endpoint-name]
swctl make:command [folder-path]
swctl make:event [folder-path]
swctl make:struct [folder-path] [struct-name]
```

## Conventions
- Register new services and handlers in `container.go` (often auto-registered).
- Add new routes in `bootstrap/internal/app/server/router.go`.
- Add new GORM models to `cmd/loader/main.go` for migrations.
- Keep environment configs in `.env` (see `.example.env`).

## Testing
- Prefer existing test helpers via `NewTestApp()` (loads `.test.env`).
- Add tests near feature code and keep DB setup explicit.

## When Unsure
- Inspect `CLAUDE.md` for more detailed project guidance.
