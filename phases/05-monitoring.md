# Phase 5 — Monitoring Stack (MANUAL, outside Ansible)

Status: **DONE** (2026-07-31)

## Why this isn't in Ansible
Running Prometheus/Grafana as containers and provisioning them requires
docker-level operations that Ansible/Semaphore doesn't perform for this
project. You set this up by hand.

## What needs to exist before Phase 6 can run
- [x] Prometheus 3.13.2 running (`docker/monitoring/docker-compose.yml`,
      `network_mode: host`), scraping:
      - node_exporter (host) — up
      - cAdvisor (containers) — up
      - postgres_exporter (DB) — up
      - bench-app `/metrics` — down when no app profile is running
        (expected, apps are one-at-a-time)
- [x] Grafana 13.1.1 running, Prometheus datasource auto-provisioned
      (`uid: prometheus_ds`)
- [x] Base dashboard "Benchmark Overview" (single combined dashboard, per
      user preference): host CPU %, host memory %, per-container CPU,
      per-container memory, app request rate, app latency p95, app status
      codes, DB active connections, DB transaction rate —
      `docker/monitoring/grafana/provisioning/dashboards/json/overview.json`
- [x] All targets showing "up" in Prometheus (confirmed via
      `/api/v1/targets`)

## Notes
- Ansible's only observability-adjacent job (Phase 6) is triggering test
  runs — it doesn't deploy or configure this stack.
- Two real bugs found and fixed while standing this up (see
  docs/decisions.md): host-owned bind mounts needed chown to each image's
  internal UID (Prometheus `65534:65534`, Grafana `472:0`), and Grafana
  initially used bridge networking while its datasource pointed at
  `localhost:9090` (only reachable because Prometheus uses host
  networking) — fixed by switching Grafana to `network_mode: host` too.
