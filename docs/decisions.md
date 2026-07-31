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
  installed on target server, connection verified.
- 2026-07-30: Target server has no passwordless sudo — Ansible `become`
  will need a password (stored as a Semaphore secret, not committed to
  files in this repo).

## Open questions (to resolve when the relevant phase starts)
- Access details (URL, login) for the existing remote Semaphore instance,
  and where this repo/`ansible/` should be pushed so that instance can
  pull it as a Repository. (Phase 1)
- Load-test tool: k6 vs hey vs wrk2 vs Locust vs autocannon? (Phase 6)
- Exact load matrix beyond the two examples given (1000/10s, 2000/10s) —
  how many steps, max load, ramp vs fixed-rate? (Phase 6)
- DB reset strategy: drop/recreate schema vs restore from a snapshot/seed
  vs ephemeral container recreation per run? (Phase 3)
- Whether apps run one-at-a-time (isolated resource envelope) or need to be
  strictly identical container resource limits (cpu/mem) for a fair
  comparison — recommend explicit `--cpus`/`--memory` limits on both.
  (Phase 4)
- Grafana dashboard scope: one dashboard per app vs single side-by-side
  comparison dashboard. (Phase 5/8)
