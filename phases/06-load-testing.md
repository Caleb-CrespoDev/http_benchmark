# Phase 6 — Load Testing

Status: **NOT STARTED**

## Goal
Define and automate the actual stress test matrix against both apps, with
a clean DB before every single run.

## Todo
- [ ] Decide load-test tool with user (options: k6, hey, wrk2, autocannon,
      Locust — k6 recommended for good Prometheus/metrics output and
      scriptable scenarios, but confirm with user + confirm version)
- [ ] Define the load matrix precisely: list of (request count, duration)
      or (rate, duration) steps — e.g. 1000 req/10s, 2000 req/10s,
      5000 req/10s, ... — confirm max step and step count with user
- [ ] Decide which endpoint(s) from Phase 4 get load-tested (all of them?
      weighted mix?)
- [ ] Write Ansible playbook `run-loadtest.yml`:
      1. call `reset-db.yml` (Phase 3)
      2. ensure only the target app is running (Phase 4 decision)
      3. run load-test tool for the given step against the app
      4. collect tool output (latency percentiles, error count, throughput)
      5. snapshot/export relevant Prometheus metrics for that run's time
         window
      6. tag/store results per (app, load-step, timestamp) in `results/`
- [ ] Parametrize playbook so Semaphore can pass in app + load-step as task
      variables
- [ ] Dry run one load step manually end-to-end before wiring into
      Semaphore templates
