# Node/Express vs Go — Stress Test Benchmark Plan

## Goal
Compare Node.js+Express vs Go (net/http or similar) under increasing load
(e.g. 1000 req/10s, 2000 req/10s, ...), both backed by the same shared
PostgreSQL database, which is reset to a clean state before every test run.
Everything (app, DB, host) is observed via Prometheus + Grafana. Test
execution is driven by Ansible playbooks, triggered from Semaphore
(ansible-semaphore) running locally.

## Fixed decisions (confirmed with user)
- Database: **PostgreSQL** (shared by both apps, wiped between test runs)
- Target server: `192.168.1.4`, user `keileb`, Ubuntu 26.04 LTS, x86_64,
  7.2GB RAM, 110GB disk. No passwordless sudo (must handle `become` password
  in Ansible/Semaphore).
- Container runtime on target server: **Docker**
- Semaphore: already running on a separate, existing server — not stood up
  locally for this project. This project just registers itself as a
  Project/Inventory/Repository against that existing instance.
- SSH key: dedicated key `~/.ssh/simple_http_bench_ed25519` (no passphrase,
  used only for this project), already installed in the server's
  `authorized_keys`. Connection verified 2026-07-30.
- Everything on the target server is currently empty — all software must be
  installed from scratch via Ansible.

## Versioning rule
My knowledge of current tool versions may be stale. **Before installing any
piece of software for the first time (Semaphore, Ansible, Node.js, Go,
PostgreSQL, Prometheus, Grafana, Docker, load-test tool, exporters, etc.),
I must ask the user to confirm the version to use** rather than assuming.
Confirmed versions get recorded in `docs/versions.md`.

## Directory layout
```
simple_http_test/
  PLAN.md                 <- this file
  phases/                 <- one file per phase, each with its own TODO checklist
  docs/versions.md         <- confirmed tool versions (filled in as we go)
  docs/decisions.md        <- open questions / decisions log
  ansible/
    inventory/             <- Ansible inventory for the target server
    playbooks/              <- playbooks (bootstrap, deploy, test, reset-db, ...)
    roles/                   <- roles (docker, postgres, node_app, go_app, prometheus, grafana, loadtest)
  apps/
    node-express/           <- Node + Express benchmark app source
    go-http/                <- Go benchmark app source
  results/                 <- exported test results / reports
```

## Phases
Each phase has a dedicated file in `phases/` with a checklist so work can be
resumed at any point. Status legend: `[ ]` todo, `[~]` in progress, `[x]` done.

- [x] **Phase 0 — Prep & Access**: project scaffolding, SSH key, verify
      connection to target server. See `phases/00-prep-access.md`
- [ ] **Phase 1 — Connect to Semaphore**: register this project (Project,
      Inventory, SSH key, Repository) against the existing remote
      Semaphore instance. See `phases/01-connect-semaphore.md`
- [ ] **Phase 2 — Server Bootstrap**: base packages, Docker on target server,
      firewall, users, directories, node exporter + cAdvisor for host/container
      metrics. See `phases/02-server-bootstrap.md`
- [ ] **Phase 3 — Shared Database**: PostgreSQL container, persistent-but-
      resettable data strategy, postgres_exporter, reset playbook.
      See `phases/03-database.md`
- [ ] **Phase 4 — Benchmark Apps**: build minimal but equivalent Node/Express
      and Go apps (same endpoints, same DB schema/queries), instrumented with
      Prometheus client libs, containerized, deployed via Ansible.
      See `phases/04-apps.md`
- [ ] **Phase 5 — Monitoring Stack**: Prometheus + Grafana on target server,
      scrape configs for node/app/db/container exporters, base dashboards.
      See `phases/05-monitoring.md`
- [ ] **Phase 6 — Load Testing**: pick + install load-test tool, define load
      matrix (1000/10s, 2000/10s, ...), Ansible playbook: reset DB -> run
      load stage -> collect results -> repeat per app per load level.
      See `phases/06-load-testing.md`
- [ ] **Phase 7 — Semaphore Orchestration**: wrap playbooks as Semaphore
      templates/tasks so full benchmark suite (both apps, all load levels)
      can be launched and monitored from the UI.
      See `phases/07-semaphore-orchestration.md`
- [ ] **Phase 8 — Results & Report**: Grafana dashboard(s) for side-by-side
      comparison, export results, write up comparison report.
      See `phases/08-results-report.md`

## Next up
Phase 1 — get access to the existing remote Semaphore instance, then
register the target server + SSH key as a Semaphore inventory/key store
entry.
