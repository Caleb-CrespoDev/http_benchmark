# Phase 7 — Semaphore Orchestration

Status: **NOT STARTED**

## Goal
Make the entire benchmark suite runnable and observable from the local
Semaphore UI, not just from the command line.

## Todo
- [ ] Semaphore template: "Bootstrap server" -> `ansible/playbooks/bootstrap.yml`
- [ ] Semaphore template: "Deploy DB" -> Phase 3 playbooks
- [ ] Semaphore template: "Deploy monitoring" -> Phase 5 playbooks
- [ ] Semaphore template: "Deploy app" (survey variable: node|go) -> Phase 4
      playbooks
- [ ] Semaphore template: "Reset DB" -> `reset-db.yml`, runnable standalone
- [ ] Semaphore template: "Run load test" (survey variables: app, load-step)
      -> `run-loadtest.yml`
- [ ] Semaphore template: "Run full suite" — chains reset+deploy+test across
      every (app, load-step) combination, e.g. via a wrapping playbook or
      Semaphore's task chaining if supported by the confirmed version
- [ ] Confirm task output/logs in Semaphore are sufficient to see pass/fail
      and key metrics per run without needing to SSH in manually
- [ ] Set up schedule or manual-trigger convention — confirm with user
      whether runs should be manually triggered per session
