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
  `authorized_keys` and reused as the GitHub deploy key for this
  (currently private) repo. Connection verified 2026-07-30; end-to-end
  Semaphore ping confirmed 2026-07-31.
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
- [ ] **Phase 2 — Server Bootstrap** *(manual, not Ansible)*: base packages,
      Docker on target server, firewall, node exporter + cAdvisor.
      See `phases/02-server-bootstrap.md`
- [ ] **Phase 3 — Shared Database** *(manual, not Ansible)*: PostgreSQL
      container, schema, postgres_exporter. Reset happens via an app-side
      HTTP endpoint, not an Ansible/DB playbook. See `phases/03-database.md`
- [ ] **Phase 4 — Benchmark Apps** *(manual build/deploy)*: equivalent
      Node/Express and Go apps, each exposing a `/reset` endpoint that
      Phase 6's Ansible playbook calls. See `phases/04-apps.md`
- [ ] **Phase 5 — Monitoring Stack** *(manual, not Ansible)*: Prometheus +
      Grafana on target server, scrape configs, base dashboards.
      See `phases/05-monitoring.md`
- [ ] **Phase 6 — Load Testing** *(the only automated phase)*: pick + install
      load-test tool, define load matrix (1000/10s, 2000/10s, ...), Ansible
      playbook: call app's `/reset` endpoint -> run stress-test tool ->
      collect results -> repeat per app per load level. No become, no
      docker commands. See `phases/06-load-testing.md`
- [ ] **Phase 7 — Semaphore Orchestration**: wrap the Phase 6 playbook as
      Semaphore templates/tasks so load tests can be launched and monitored
      from the UI. See `phases/07-semaphore-orchestration.md`
- [ ] **Phase 8 — Results & Report**: Grafana dashboard(s) for side-by-side
      comparison, export results, write up comparison report.
      See `phases/08-results-report.md`

## Next up
Phase 2-5 (server, DB, apps, monitoring) are on you to set up manually.
Once the apps are up with `/reset` endpoints, come back for Phase 6 —
picking the load-test tool and writing `run-loadtest.yml`.
