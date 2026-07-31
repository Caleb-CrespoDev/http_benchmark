# Phase 8 — Results & Report

Status: **NOT STARTED**

## Goal
Turn raw run data into a clear Node/Express vs Go comparison.

## Todo
- [ ] Grafana: build (or finalize) a side-by-side comparison dashboard —
      throughput, latency percentiles (p50/p95/p99), error rate, CPU/mem
      per app, DB load, all plotted per load-step
- [ ] Export/aggregate `results/` data into a single comparison table
      (app x load-step -> key metrics)
- [ ] Write up findings: where each stack wins/loses, at what load level
      differences emerge, any surprises worth flagging
- [ ] Decide with user whether this should become an Artifact (shareable
      HTML report/dashboard) or stay as local markdown/CSV
