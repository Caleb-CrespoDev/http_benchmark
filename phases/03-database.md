# Phase 3 — Shared Database (MANUAL, outside Ansible)

Status: **YOUR RESPONSIBILITY** (not automated)

## Why this isn't in Ansible
Standing up and configuring Postgres requires docker-level operations that
Ansible/Semaphore doesn't perform for this project. You set this up by
hand.

## What needs to exist before Phase 6 can run
- [ ] PostgreSQL running (version confirmed by you), shared by both apps
- [ ] Schema for the benchmark endpoints in place
- [ ] `postgres_exporter` wired to Prometheus, if you want DB metrics
- [ ] A reset mechanism that does **not** require Ansible to touch Docker
      or the DB directly — see Phase 4: the reset is exposed as an HTTP
      endpoint on each app, and Phase 6's Ansible playbook resets state by
      calling that endpoint (e.g. `GET /reset`), nothing more.

## Notes
- The reset strategy question from `docs/decisions.md` is answered: reset
  happens app-side (the endpoint handles it against the DB internally),
  not via a standalone Ansible/Semaphore DB playbook.
