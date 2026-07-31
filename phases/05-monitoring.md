# Phase 5 — Monitoring Stack (Prometheus + Grafana)

Status: **NOT STARTED**

## Goal
Full observability of host, containers, database, and both apps during
every test run.

## Todo
- [ ] Confirm Prometheus version with user
- [ ] Confirm Grafana version with user
- [ ] Ansible role: run Prometheus as a Docker container, config scraping:
      - node_exporter (host)
      - cAdvisor (containers)
      - postgres_exporter (DB)
      - node-express app `/metrics`
      - go-http app `/metrics`
- [ ] Ansible role: run Grafana as a Docker container, provision Prometheus
      as a datasource automatically (not manual UI click-through)
- [ ] Build base dashboards: host resources, container resources, DB
      (connections/queries/locks), app (request rate, latency histogram,
      error rate)
- [ ] Confirm retention/storage settings are enough for the full test
      matrix without needing manual pruning mid-session
- [ ] Verify all targets show as "up" in Prometheus before any load test
      runs
