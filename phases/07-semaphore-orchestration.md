# Phase 7 — Semaphore Orchestration

Status: **NOT STARTED**

## Goal
Make the load-test suite (Phase 6 — the only automated phase) runnable and
observable from the existing Semaphore UI.

## Todo
- [ ] Semaphore template: "Ping" -> `ansible/playbooks/ping.yml` (done,
      Phase 1)
- [ ] Semaphore template: "Run load test" (survey variables: app,
      load-step) -> `run-loadtest.yml`
- [ ] Semaphore template: "Run full suite" — chains the reset+test loop
      across every (app, load-step) combination, e.g. via a wrapping
      playbook or Semaphore's task chaining if supported by the confirmed
      version
- [ ] Confirm task output/logs in Semaphore are sufficient to see pass/fail
      and key metrics per run without needing to SSH in manually
- [ ] Set up schedule or manual-trigger convention — confirm with user
      whether runs should be manually triggered per session

## Notes
- No "bootstrap/deploy DB/deploy monitoring/deploy app" templates —
  those are all manual, outside Ansible (Phases 2-5).
