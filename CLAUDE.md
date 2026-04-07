# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This is a Pulumi native provider demo. It has two main components:
- **`/provider`** — A Go-based Pulumi provider that exposes Organisation/Team/User as Pulumi resources
- **`/service`** — A TypeScript/Bun backend API the provider talks to, backed by SQLite

The provider is a client to the service: it translates Pulumi CRUD operations into HTTP calls to the service API.

## Commands

### Provider (Go + Task)

Run from the `/provider` directory:

```bash
task build_provider       # Compile Go binary → ./bin/pulumi-resource-nativeProvider
task ensure               # go mod tidy
task get_schema           # Generate schema.json from the provider
task generate_sdks        # Regenerate Node.js, .NET, and Python SDKs
task lint                 # Run golangci-lint
task watch                # Rebuild on file changes
task clean                # Remove bin/, .task/, schema.json
```

### Service (TypeScript/Bun)

Run from the `/service` directory:

```bash
bun install               # Install dependencies
bun run dev               # Start API server in watch mode (port 3000)
bun run db:generate       # Generate Drizzle migrations from schema files
bun run db:migrate        # Run pending migrations
```

Swagger UI is available at `http://localhost:3000/ui` when the service is running.

## Architecture

### Provider (`/provider`)

- **`main.go`** — Registers the provider, sets metadata, wires resources and functions
- **`pkg/config.go`** — Reads `apiToken` (secret) and `baseUrl` config; creates the HTTP client
- **`pkg/organisation.go`** — The only resource currently implemented; shows the full CRUD + Diff + Check pattern
- **`internal/client.go`** — Low-level HTTP client (Bearer auth, JSON marshaling); default base URL is `http://localhost:3000/api`
- **`internal/{organisation,team,user}.go`** — API model types and per-resource HTTP operations

SDKs in `provider/sdk/` are generated — do not edit them by hand; use `task generate_sdks` instead.

### Service (`/service`)

Monorepo with two packages:

- **`packages/core`** — Business logic and Drizzle ORM models. Each resource (organisation, team, user) has an `index.ts` (CRUD functions) and a `*.sql.ts` (Drizzle table schema).
- **`packages/api`** — Hono web server wiring routes to core functions. Uses `@hono/zod-openapi` for schema-validated routes. Auth is a hardcoded Bearer token (`hunter2`).

### Resource hierarchy

Organisations → Teams → Users. Resources are independent in the API but logically nested.

## Key Conventions

- Provider config (`apiToken`, `baseUrl`) is read in `pkg/config.go` and injected via `infer.ComponentResource` context — new resources should follow the pattern in `pkg/organisation.go`.
- Service route handlers validate input with Zod and return a `Result` wrapper (see `packages/api/src/common.ts`).
- Database schema lives in `*.sql.ts` files; after changing them, run `db:generate` then `db:migrate`.
