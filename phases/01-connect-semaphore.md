# Phase 1 — Connect to Semaphore

Status: **NOT STARTED** (this is next)

## Goal
Semaphore is already running on a separate, existing server (no need to
stand up a local instance). This phase just wires that existing Semaphore
up to this project: a Project, the target server's SSH key, an Inventory
entry, the become-password secret, and a Repository pointing at this
project's playbooks.

## Todo
- [ ] Get access details for the existing Semaphore instance (URL,
      admin/login) from the user
- [ ] Create a Semaphore **Project** for this benchmark
- [ ] Add target server SSH key (`simple_http_bench_ed25519` private key)
      to Semaphore's Key Store
- [ ] Add target server `192.168.1.4` to a Semaphore **Inventory**
- [ ] Add the sudo/become password for `keileb` on the target server as a
      Semaphore secret (used for `become` in later playbooks)
- [ ] Point Semaphore at this project's `ansible/` directory as a
      **Repository** (git remote, since Semaphore runs on a different
      server and can't use a local path on this machine)
- [ ] Run a trivial "ping" playbook/task from Semaphore against the target
      server to confirm end-to-end wiring works

## Notes
- Everything after this phase should be launched/observed through
  Semaphore rather than run ad-hoc, so the UI stays the single source of
  truth for what's been executed.
- Because Semaphore is remote, this repo (or at least `ansible/`) needs to
  be pushed to a git remote it can pull from — decide where before this
  phase's Repository step.
