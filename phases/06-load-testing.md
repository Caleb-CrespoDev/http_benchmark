# Phase 6 — Load Testing

Status: **DONE** (2026-07-31)

## Goal
This is the only phase Ansible/Semaphore actually automates. Everything
upstream (Docker, DB, apps, monitoring — Phases 2-5) is set up manually by
you. Ansible's job here is narrow and non-privileged: run `hey` against the
app, from the Semaphore host (192.168.1.5) over the network — no SSH to
`noteb`, no `become`, no docker commands. Reset is a **separate** playbook,
not chained into the load-test playbooks.

## Todo
- [x] Load-test tool: **`hey`** — single static binary, no scripting,
      `apt install hey` (0.1.4, Ubuntu package) on 192.168.1.5
- [x] Load matrix: **1000, 2000, 5000, 10000 req/10s** — one playbook per
      level (not one parametrized-by-step playbook), `app` (node|go) is a
      survey variable within each, used only to label results
- [x] Endpoints: **all five** — `GET /healthz`, `GET /items`, `POST /items`,
      `PUT /items/{id}`, `DELETE /items/{id}` — split 20/40/25/10/5% of
      total rate, run concurrently via 5 separate `hey` processes per step
      (Phase 4 gained `PUT`/`DELETE` endpoints on both apps for this)
- [x] Playbooks:
      - `ansible/playbooks/reset.yml` — standalone, `POST /reset`, triggered
        manually whenever a clean DB is wanted, not auto-chained
      - `ansible/playbooks/run-loadtest-{1000,2000,5000,10000}.yml` — each:
        seed one item (for PUT/DELETE targets) -> launch 5 `hey` processes
        via `async`/`poll: 0` -> wait via `async_status` -> write each
        endpoint's output to `~/bench-results/<app>-<level>-<timestamp>/`
        on 192.168.1.5
- [x] Dry run: all 5 playbooks syntax-checked; `reset.yml` and
      `run-loadtest-1000.yml` executed end-to-end against the live Go app
      on `noteb` — all endpoints behaved as designed (200s for
      healthz/GET, DELETE = 1x204 then 404s for the rest, matching the
      known single-ID limitation)

## Notes — bugs found and fixed during dry runs
- Go app's `/healthz` and `/reset` handlers didn't set
  `Content-Type: application/json`, unlike the other handlers — broke
  Ansible's `uri` module JSON auto-parsing. Fixed in `apps/go-http/main.go`.
- `results_dir` originally used `lookup('pipe', 'date ...')` inline in
  `vars:` — Ansible re-evaluates that on **every reference**, not once, so
  each parallel `hey` task got a different timestamp and wrote to a
  directory that didn't exist (only the first-created one did). Fixed by
  computing the timestamp once via `set_fact` at the top of each playbook.
