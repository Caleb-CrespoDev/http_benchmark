# Phase 2 — Server Bootstrap (MANUAL, outside Ansible)

Status: **YOUR RESPONSIBILITY** (not automated)

## Why this isn't in Ansible
Ansible/Semaphore for this project never uses `become`/sudo. All privileged
setup on the target server — Docker install, base packages, firewall,
exporters — is done by hand, once, outside this repo's automation.

## What needs to exist before later phases can run
- [ ] Docker Engine + Compose plugin installed
- [ ] `keileb` added to the `docker` group (so app containers can be run
      without sudo, if needed for manual deploys)
- [ ] Base packages (curl, git, ca-certificates, etc.)
- [ ] Directory layout for persistent data (e.g. `/opt/bench/{postgres,prometheus,grafana,apps}`)
- [ ] Firewall (ufw) — only needed ports open (app ports, Grafana,
      Prometheus if remote scraping, SSH)
- [ ] `node_exporter` running (host metrics)
- [ ] `cAdvisor` running (container metrics)

## Notes
- Nothing here is run by Semaphore. This checklist exists just so you have
  a record of what the target server needs before Phases 3-5's manual
  setup and Phase 6's automated load testing.
