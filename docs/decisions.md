# Decisions log

## Confirmed
- 2026-07-30: Database = PostgreSQL.
- 2026-07-30: Target server = 192.168.1.4, user `keileb`, Ubuntu 26.04 LTS.
- 2026-07-30: Container runtime = Docker (target server).
- 2026-07-31: Semaphore is already running on a separate, existing server
  — no local Semaphore instance is stood up for this project. Phase 1
  just registers this project against it (Project/Inventory/SSH
  key/Repository).
- 2026-07-30: Dedicated SSH key `~/.ssh/simple_http_bench_ed25519` created,
  installed on target server, connection verified; reused as the GitHub
  deploy key for this repo (2026-07-31).
- 2026-07-31: Ansible/Semaphore for this project **never uses `become`/
  sudo and never runs docker commands**. All privileged/docker setup
  (Docker install, DB, apps, monitoring — Phases 2-5) is done manually by
  the user. Ansible's only automated job (Phase 6+) is calling each app's
  `/reset` HTTP endpoint and running the stress-test tool.
- 2026-07-31: End-to-end Semaphore -> target server connectivity confirmed
  via `ansible/playbooks/ping.yml` (ok=5, no failures).
- 2026-07-31: `http_benchmark` GitHub repo made public — Semaphore/target
  server clone it over plain HTTPS, no deploy key needed.
- 2026-07-31: App server (`noteb`, 192.168.1.4) hosts Docker, Postgres,
  the apps, and monitoring. Semaphore (192.168.1.5) is where load-test
  runs are triggered from — Ansible's Phase 6 playbook makes HTTP calls
  from there to `noteb`'s app port, no SSH/become needed for that step.
  `ufw` on `noteb` is intentionally left inactive.
- 2026-07-31: Reset endpoint contract = `POST /reset` on each app,
  `TRUNCATE TABLE items RESTART IDENTITY`.
- 2026-07-31: Apps run one-at-a-time, both on port 4000, both limited to
  2 vCPU / 4GiB (`docker/apps/docker-compose.yml`, Compose profiles
  `node`/`go` sharing container name `bench-app`).
- 2026-07-31: Shared schema = single `items` table (`id`, `value`,
  `created_at`), `docker/postgres/init/001-schema.sql`.
- 2026-07-31: Both apps smoke-tested end-to-end on `noteb` via
  `docker/apps/smoke-test.sh` (healthz, reset, write, read-back, metrics)
  — caught and fixed a real bug: `docker/apps/.env` must use the key
  `PGPASSWORD` (what the apps read), not `POSTGRES_PASSWORD` (Postgres'
  own env var name) — easy to mix up since both `.env` files hold the
  same password value.
- 2026-07-31: Prometheus + Grafana stood up on `noteb`, all targets
  confirmed "up" (node, cadvisor, postgres; bench-app down when idle,
  expected). Found and fixed two bugs: bind-mounted data dirs needed
  chown to each image's internal UID (Prometheus `65534:65534`, Grafana
  `472:0`), and Grafana needed `network_mode: host` to reach Prometheus at
  `localhost:9090` (was on bridge networking, couldn't reach it at all).
- 2026-07-31: Grafana dashboard scope = single combined "Benchmark
  Overview" dashboard (not split per-app/per-resource), covering host,
  container, DB, and app metrics together.

## Open questions (to resolve when the relevant phase starts)
- Load-test tool: k6 vs hey vs wrk2 vs Locust vs autocannon? (Phase 6)
- Exact load matrix beyond the two examples given (1000/10s, 2000/10s) —
  how many steps, max load, ramp vs fixed-rate? (Phase 6)
