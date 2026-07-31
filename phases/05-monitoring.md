# Phase 5 — Monitoring Stack (MANUAL, outside Ansible)

Status: **YOUR RESPONSIBILITY** (not automated)

## Why this isn't in Ansible
Running Prometheus/Grafana as containers and provisioning them requires
docker-level operations that Ansible/Semaphore doesn't perform for this
project. You set this up by hand.

## What needs to exist before Phase 6 can run
- [ ] Prometheus running, scraping:
      - node_exporter (host)
      - cAdvisor (containers)
      - postgres_exporter (DB), if set up
      - node-express app `/metrics`
      - go-http app `/metrics`
- [ ] Grafana running, Prometheus datasource configured
- [ ] Base dashboards: host resources, container resources, DB, app
      (request rate, latency histogram, error rate)
- [ ] All targets showing "up" in Prometheus before any load test runs

## Notes
- Ansible's only observability-adjacent job (Phase 6) is triggering test
  runs — it doesn't deploy or configure this stack.
