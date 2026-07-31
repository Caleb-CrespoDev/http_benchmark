# Phase 4 — Benchmark Apps (Node/Express vs Go)

Status: **MOSTLY YOUR RESPONSIBILITY** (build/deploy manual; must expose
two endpoints Ansible calls in Phase 6)

## Goal
Two functionally-identical HTTP apps — same endpoints, same DB
queries/schema, same response payloads — one in Node.js+Express, one in Go
(stdlib `net/http` or a comparably minimal router), so the only real
variable is the runtime/framework.

## What Ansible needs from these apps (the only contract that matters)
- [ ] Each app exposes a **reset endpoint** (e.g. `GET /reset` or
      `POST /reset`) that clears/reseeds its DB state — called by Ansible
      via the `uri` module before every load-test run, no docker/sudo
      involved
- [ ] Each app exposes whatever endpoint(s) the load-test tool will hit
      (Phase 6)

## Everything else (manual, outside Ansible)
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
      images) — built/deployed by you, not Ansible
- [ ] Explicit, matching resource limits for both containers when run
      (`--cpus`, `--memory`) — fairness requirement
- [ ] Decide: both apps running simultaneously on different ports, or one
      at a time during tests — recommend one-at-a-time to avoid resource
      contention skewing results (record decision in docs/decisions.md)
- [ ] Smoke test both apps manually (curl) before wiring into load tests
