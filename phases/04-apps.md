# Phase 4 — Benchmark Apps (Node/Express vs Go)

Status: **NOT STARTED**

## Goal
Two functionally-identical HTTP apps — same endpoints, same DB
queries/schema, same response payloads — one in Node.js+Express, one in Go
(stdlib `net/http` or a comparably minimal router), so the only real
variable is the runtime/framework.

## Todo
- [ ] Confirm Node.js version (LTS) with user before pinning
- [ ] Confirm Go version with user before pinning
- [ ] Design the minimal endpoint set to exercise (recommend at least: one
      read endpoint hitting Postgres, one write endpoint hitting Postgres,
      one trivial no-DB endpoint as a baseline)
- [ ] Write `apps/node-express` app: same routes, pg client, Prometheus
      metrics middleware (e.g. `prom-client`), structured logging
- [ ] Write `apps/go-http` app: same routes, `database/sql` + pg driver,
      Prometheus metrics (`prometheus/client_golang`), structured logging
- [ ] Both apps expose `/metrics` for Prometheus scraping
- [ ] Both apps expose `/healthz` for readiness checks
- [ ] Containerfiles for both apps (multi-stage builds, minimal final
      images)
- [ ] Explicit, matching resource limits for both containers when run
      (`--cpus`, `--memory`) — fairness requirement
- [ ] Ansible role/playbook to build + deploy both app containers to the
      target server via Docker
- [ ] Decide: both apps running simultaneously on different ports, or one
      at a time during tests — recommend one-at-a-time to avoid resource
      contention skewing results (record decision in docs/decisions.md)
- [ ] Smoke test both apps manually (curl) before wiring into load tests
