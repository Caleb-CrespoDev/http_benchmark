# Confirmed tool versions

Filled in as each tool is installed. Rule: always ask the user to confirm
the current version before first install (see PLAN.md "Versioning rule") —
do not trust model knowledge for this, as it may be stale.

| Tool | Version | Confirmed date | Notes |
|------|---------|-----------------|-------|
| Ansible (control node) | TBD | | local machine |
| Docker (target server) | TBD | | 192.168.1.4, installed manually |
| PostgreSQL | 18.4 | 2026-07-31 | shared DB container |
| Node.js | 24.18.1 (LTS) | 2026-07-31 | benchmark app |
| Express | ^4.21.2 | 2026-07-31 | benchmark app |
| Go | 1.26.5 | 2026-07-31 | benchmark app |
| Prometheus | 3.13.2 | 2026-07-31 | target server |
| Grafana | 13.1.1 | 2026-07-31 | target server |
| node_exporter | 1.12.1 | 2026-07-31 | host metrics |
| cAdvisor | 0.60.5 | 2026-07-31 | container metrics |
| postgres_exporter | 0.20.1 | 2026-07-31 | DB metrics |
| Load test tool (hey) | 0.1.4 | 2026-07-31 | apt package, 192.168.1.5 |
