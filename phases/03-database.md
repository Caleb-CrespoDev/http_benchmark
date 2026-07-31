# Phase 3 — Shared Database (MANUAL, outside Ansible)

Status: **DONE** (2026-07-31) — schema design rolled into Phase 4

## Why this isn't in Ansible
Standing up and configuring Postgres requires docker-level operations that
Ansible/Semaphore doesn't perform for this project. You set this up by
hand.

## What needs to exist before Phase 6 can run
- [x] PostgreSQL 18.4 running via `docker/postgres/docker-compose.yml`,
      data at `/opt/bench/postgres` (note: PG 18 image changed its VOLUME
      convention to `/var/lib/postgresql`, not `.../data` — compose file
      already accounts for this)
- [ ] Schema for the benchmark endpoints in place — deferred to Phase 4
      (designed alongside the actual app endpoints)
- [x] `postgres_exporter` v0.20.1 wired to Postgres via `DATA_SOURCE_NAME`,
      `docker/postgres/docker-compose.yml`, `127.0.0.1:9187`
- [x] Reset mechanism confirmed as app-side (Phase 4 exposes `/reset`,
      Phase 6's Ansible playbook just calls it over HTTP) — no
      docker/sudo involved

## Notes
- The reset strategy question from `docs/decisions.md` is answered: reset
  happens app-side (the endpoint handles it against the DB internally),
  not via a standalone Ansible/Semaphore DB playbook.
