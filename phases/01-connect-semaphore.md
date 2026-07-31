# Phase 1 — Connect to Semaphore

Status: **DONE** (2026-07-31)

## Goal
Semaphore is already running on a separate, existing server (no need to
stand up a local instance). This phase just wires that existing Semaphore
up to this project: a Project, the target server's SSH key, an Inventory
entry, and a Repository pointing at this project's playbooks.

## Todo
- [x] Get access details for the existing Semaphore instance (URL,
      admin/login) from the user
- [x] Create a Semaphore **Project** for this benchmark
- [x] Add target server SSH key (`simple_http_bench_ed25519` private key,
      reused as the GitHub deploy key) to Semaphore's Key Store
- [x] Add target server `192.168.1.4` to a Semaphore **Inventory**
      (`ansible/inventory/hosts.yml`, host alias `bench-target`)
- [x] Point Semaphore at this project's git repo as a **Repository**,
      authenticated via the reused SSH key as a read-only GitHub deploy key
- [x] Run a trivial "ping" playbook/task from Semaphore against the target
      server to confirm end-to-end wiring works — succeeded 2026-07-31,
      `bench-target` (Ubuntu 26.04, hostname `noteb`) reachable, facts
      gathered, `ok=5 changed=0 unreachable=0 failed=0`

## Notes
- No `become`/sudo secret needed — this project's Ansible never escalates
  privileges or runs docker commands (see PLAN.md). Phases 2-5 are done
  manually by the user; Ansible's only job (Phase 6+) is HTTP calls to
  each app's `/reset` endpoint and running the stress-test tool.
- Everything after this phase should be launched/observed through
  Semaphore rather than run ad-hoc, so the UI stays the single source of
  truth for what's been executed.
