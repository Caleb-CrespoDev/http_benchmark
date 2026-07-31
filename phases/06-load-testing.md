# Phase 6 — Load Testing

Status: **NOT STARTED**

## Goal
This is the only phase Ansible/Semaphore actually automates. Everything
upstream (Docker, DB, apps, monitoring — Phases 2-5) is set up manually by
you. Ansible's job here is narrow and non-privileged: hit an HTTP reset
endpoint, then run the stress-test tool. No `become`, no docker commands.

## Todo
- [ ] Decide load-test tool with user (options: k6, hey, wrk2, autocannon,
      Locust — confirm with user + confirm version). Tool is installed by
      you on the target server (or run from the control node against the
      target's IP, if that's simpler); Ansible just invokes it.
- [ ] Define the load matrix precisely: list of (request count, duration)
      or (rate, duration) steps — e.g. 1000 req/10s, 2000 req/10s,
      5000 req/10s, ... — confirm max step and step count with user
- [ ] Decide which endpoint(s) from Phase 4 get load-tested (all of them?
      weighted mix?)
- [ ] Write Ansible playbook `run-loadtest.yml`:
      1. `uri` module: `GET`/`POST` the app's `/reset` endpoint (Phase 4
         contract) — the only "state-changing" step, and it's just an
         HTTP call, no docker/sudo
      2. `command`/`shell` module: run the load-test tool for the given
         step against the app's URL
      3. collect tool output (latency percentiles, error count, throughput)
      4. tag/store results per (app, load-step, timestamp) in `results/`
- [ ] Parametrize playbook so Semaphore can pass in app + load-step as task
      variables (survey variables)
- [ ] Dry run one load step manually end-to-end (`ansible-playbook ...`)
      before wiring into Semaphore templates
