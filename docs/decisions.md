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

## Open questions (to resolve when the relevant phase starts)
- Load-test tool: k6 vs hey vs wrk2 vs Locust vs autocannon? (Phase 6)
- Exact load matrix beyond the two examples given (1000/10s, 2000/10s) —
  how many steps, max load, ramp vs fixed-rate? (Phase 6)
- Reset endpoint contract: exact path/method (e.g. `GET /reset` vs
  `POST /reset`) and expected response, so both apps and the Ansible
  playbook agree. (Phase 4/6)
- Whether apps run one-at-a-time (isolated resource envelope) or need to be
  strictly identical container resource limits (cpu/mem) for a fair
  comparison — recommend explicit `--cpus`/`--memory` limits on both.
  (Phase 4)
- Grafana dashboard scope: one dashboard per app vs single side-by-side
  comparison dashboard. (Phase 5/8)
