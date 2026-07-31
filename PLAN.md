# Node/Express vs Go — Stress Test Benchmark Plan

## Goal
Compare Node.js+Express vs Go (net/http or similar) under increasing load
(e.g. 1000 req/10s, 2000 req/10s, ...), both backed by the same shared
PostgreSQL database, which is reset to a clean state before every test run.
Everything (app, DB, host) is observed via Prometheus + Grafana.

Docker install, the DB, the apps, and the monitoring stack are all set up
**manually by the user** — Ansible/Semaphore never uses `become`/sudo and
never runs docker commands. The only thing Ansible automates is the load
test itself: calling each app's `/reset` HTTP endpoint, then running the
stress-test tool, triggered from the existing (remote) Semaphore instance.

## Fixed decisions (confirmed with user)
- Database: **PostgreSQL** (shared by both apps, wiped between test runs
  via an app-exposed HTTP reset endpoint — not an Ansible/DB playbook)
- Target server: `192.168.1.4`, user `keileb`, Ubuntu 26.04 LTS, x86_64,
  7.2GB RAM, 110GB disk.
- Ansible/Semaphore for this project **never uses `become`/sudo and never
  runs docker commands**. All privileged/docker setup (Docker install, DB,
  apps, monitoring) is done manually by the user. Ansible's only job is
  triggering load-test runs (HTTP reset call + stress-test command).
- Container runtime on target server: **Docker** (installed manually)
- Semaphore: already running on a separate, existing server — not stood up
  locally for this project. This project just registers itself as a
  Project/Inventory/Repository against that existing instance.
- SSH key: dedicated key `~/.ssh/simple_http_bench_ed25519` (no passphrase,
  used only for this project), already installed in the server's
  `authorized_keys`. Connection verified 2026-07-30; end-to-end Semaphore
  ping confirmed 2026-07-31.
- Repo `http_benchmark` is public on GitHub (was private, opened up
  2026-07-31) — cloned via plain HTTPS on `noteb`, no deploy key needed.
- Everything on the target server beyond the automated load-test playbooks
  is installed/configured manually by the user.

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
    playbooks/              <- ping.yml, run-loadtest.yml (only phases Ansible drives)
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
- [x] **Phase 1 — Connect to Semaphore**: register this project (Project,
      Inventory, SSH key, Repository) against the existing remote
      Semaphore instance. Connectivity confirmed via ping playbook. See
      `phases/01-connect-semaphore.md`
- [x] **Phase 2 — Server Bootstrap** *(manual, not Ansible)*: base packages,
      Docker on target server, `/opt/bench` layout, node_exporter + cAdvisor
      running. See `phases/02-server-bootstrap.md`
- [x] **Phase 3 — Shared Database** *(manual, not Ansible)*: PostgreSQL 18.4
      + postgres_exporter running on `noteb`. Schema deferred to Phase 4.
      See `phases/03-database.md`
- [x] **Phase 4 — Benchmark Apps** *(manual build/deploy)*: Node 24.18.1
      LTS + Express and Go 1.26.5 apps built, both exposing `/healthz`,
      `/items` (GET/POST), `/items/{id}` (PUT/DELETE), `POST /reset`,
      `/metrics`. See `phases/04-apps.md`
- [x] **Phase 5 — Monitoring Stack** *(manual, not Ansible)*: Prometheus
      3.13.2 + Grafana 13.1.1 running on `noteb`, all targets up, combined
      "Benchmark Overview" dashboard provisioned. See
      `phases/05-monitoring.md`
- [x] **Phase 6 — Load Testing** *(the only automated phase)*: `hey`
      0.1.4, load matrix 1000/2000/5000/10000 req/10s as 4 standalone
      playbooks hitting all 5 endpoints, plus a standalone `reset.yml`.
      Dry-run verified end-to-end on `noteb`. See `phases/06-load-testing.md`
- [ ] **Phase 7 — Semaphore Orchestration**: wrap the Phase 6 playbook as
      Semaphore templates/tasks so load tests can be launched and monitored
      from the UI. See `phases/07-semaphore-orchestration.md`
- [ ] **Phase 8 — Results & Report**: Grafana dashboard(s) for side-by-side
      comparison, export results, write up comparison report.
      See `phases/08-results-report.md`

## Next up
Phase 7 — wrap `reset.yml` and the four `run-loadtest-*.yml` playbooks as
Semaphore templates (with the `app` survey variable) so they're
launchable from the UI.
