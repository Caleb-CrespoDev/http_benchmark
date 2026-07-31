# Phase 0 — Prep & Access

Status: **DONE** (2026-07-30)

## Todo
- [x] Create project directory skeleton (`ansible/`, `apps/`, `docs/`,
      `phases/`, `results/`)
- [x] Generate dedicated SSH keypair for this project
      (`~/.ssh/simple_http_bench_ed25519`, no passphrase)
- [x] Install public key on target server `192.168.1.4` (user `keileb`)
      — done manually by user via `authorized_keys`
- [x] Verify SSH connection from local machine to target server
- [x] Record target server facts: Ubuntu 26.04 LTS, x86_64, 7.2GB RAM,
      110GB disk, no passwordless sudo, nothing installed yet

## Notes
- Server has no docker/docker/ansible installed — greenfield, everything
  comes from Phase 2 onward.
- No passwordless sudo: every phase that touches the target server via
  Ansible needs `--ask-become-pass` or a become password supplied through
  Semaphore's secret storage. Do not store this password in plaintext in
  this repo.
