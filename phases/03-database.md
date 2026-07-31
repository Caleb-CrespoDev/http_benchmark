# Phase 3 — Shared Database (PostgreSQL)

Status: **NOT STARTED**

## Goal
One PostgreSQL instance, shared by both apps, that can be reset to an
identical clean state before every test run.

## Todo
- [ ] Confirm PostgreSQL version with user before installing
- [ ] Decide + record reset strategy (see docs/decisions.md open question):
      recommend a `reset` playbook that drops & recreates the schema and
      re-applies migrations/seed data — faster and simpler than recreating
      the whole container each run
- [ ] Define schema needed for the benchmark endpoints (kept intentionally
      simple — e.g. a single table both apps read/write to)
- [ ] Ansible role: run PostgreSQL as a Docker container with a named
      volume, expose only to the server's internal network (not public)
- [ ] Init/migration SQL committed to `ansible/roles/postgres/files/`
- [ ] `postgres_exporter` container wired to Prometheus
- [ ] Ansible playbook: `reset-db.yml` — drop/recreate schema + reseed,
      idempotent, safe to run before every single test
- [ ] Verify: connect from local machine (or via ssh tunnel) with `psql`,
      confirm schema present
- [ ] Wrap as a Semaphore template ("Reset DB") that can be run standalone
      or as a pre-step in the load-test template
