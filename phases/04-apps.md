# Phase 4 — Benchmark Apps (Node/Express vs Go)

Status: **DONE** (2026-07-31)

## Goal
Two functionally-identical HTTP apps — same endpoints, same DB
queries/schema, same response payloads — one in Node.js+Express, one in Go
(stdlib `net/http`), so the only real variable is the runtime/framework.

## What Ansible needs from these apps (the only contract that matters)
- [x] Each app exposes a **reset endpoint**: `POST /reset` — `TRUNCATE
      TABLE items RESTART IDENTITY` — called by Ansible via the `uri`
      module before every load-test run, no docker/sudo involved
- [x] Each app exposes the endpoint(s) the load-test tool will hit:
      `GET /items`, `POST /items` (Phase 6)

## Everything else
- [x] Node.js version confirmed: **24.18.1 LTS**
- [x] Go version confirmed: **1.26.5**
- [x] Endpoint set: `GET /healthz` (no DB), `GET /items` (SELECT),
      `POST /items` (INSERT), `POST /reset` (TRUNCATE), `GET /metrics`
- [x] Shared schema: single `items` table (`id`, `value`, `created_at`) —
      `docker/postgres/init/001-schema.sql`
- [x] `apps/node-express`: Express 4, `pg`, `prom-client`
      (`http_request_duration_seconds` histogram + default metrics)
- [x] `apps/go-http`: stdlib `net/http`, `pgx/v5`, `prometheus/client_golang`
- [x] Both apps expose `/metrics` and `/healthz`
- [x] Dockerfiles for both (multi-stage, minimal final images —
      `node:24.18.1-slim` and `gcr.io/distroless/static-debian12`)
- [x] Resource limits: **2 vCPU / 4GiB** for both, set in
      `docker/apps/docker-compose.yml`
- [x] Decision: one-at-a-time — both apps share container name `bench-app`
      and Compose profiles (`node` / `go`) so only one can run at a time,
      both on port 4000
- [x] Smoke tested locally (curl `/healthz`, `/metrics` for both); built
      successfully on `noteb`
