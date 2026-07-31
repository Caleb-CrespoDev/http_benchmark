# Phase 2 — Server Bootstrap

Status: **NOT STARTED**

## Goal
Get the target server (192.168.1.4) from empty Ubuntu 26.04 to a machine
ready to host Docker containers, with base host-level observability.

## Todo
- [ ] Ansible inventory entry + group_vars for target server
- [ ] Playbook: apt update/upgrade, base packages (curl, git, ca-certificates, etc.)
- [ ] Install Docker Engine + Docker Compose plugin
- [ ] Confirm Docker version with user before installing
- [ ] Add `keileb` to the `docker` group so Ansible/Semaphore can manage
      containers without root on every task (still needs `become` for the
      install steps themselves)
- [ ] Set up directory layout on server for persistent data
      (e.g. `/opt/bench/{postgres,prometheus,grafana,apps}`)
- [ ] Configure firewall (ufw) — open only needed ports (app ports, Grafana,
      Prometheus if remote scraping needed, SSH)
- [ ] Install + start `node_exporter` (host metrics) as a systemd service or
      container
- [ ] Install + start `cAdvisor` (container metrics) — Docker-compatible
      setup
- [ ] Verify both exporters respond locally (`curl localhost:9100/metrics`,
      `curl localhost:8080/metrics`)
- [ ] Run this whole phase as a Semaphore task, confirm idempotency
      (re-run playbook, no errors/changes second time)

## Notes
- No passwordless sudo on this server — playbooks need `become: true` and
  Semaphore must supply the become password from its secret store.
