# Phase 2 — Server Bootstrap (MANUAL, outside Ansible)

Status: **DONE** (2026-07-31)

## Why this isn't in Ansible
Ansible/Semaphore for this project never uses `become`/sudo. All privileged
setup on the target server — Docker install, base packages, firewall,
exporters — is done by hand, once, outside this repo's automation.

## What needs to exist before later phases can run
- [x] Docker Engine + Compose plugin installed
- [x] `keileb` added to the `docker` group (so app containers can be run
      without sudo, if needed for manual deploys)
- [x] Base packages (curl, git, ca-certificates, etc.)
- [x] Directory layout for persistent data: `/opt/bench/{postgres,prometheus,grafana,apps}`,
      owned by `keileb`
- [x] Firewall (ufw) — left **inactive** (intentional choice)
- [x] `node_exporter` v1.12.1 running via `docker/monitoring/docker-compose.yml`
      (host metrics, `:9100`)
- [x] `cAdvisor` v0.60.5 running via `docker/monitoring/docker-compose.yml`
      (container metrics, `127.0.0.1:8080`)

## Notes
- Nothing here is run by Semaphore. This checklist exists just so you have
  a record of what the target server needs before Phases 3-5's manual
  setup and Phase 6's automated load testing.
