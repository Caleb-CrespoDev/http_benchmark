#!/usr/bin/env bash
# Exercises the full endpoint contract against whichever app is currently
# running on :4000 (bring one up first: `docker compose --profile node|go up -d`).
set -euo pipefail

HOST="${1:-localhost:4000}"
fail=0

check() {
  local desc="$1" expected_status="$2"; shift 2
  local status body
  body=$(curl -sS -o /tmp/smoke-body -w "%{http_code}" "$@") || { echo "FAIL  $desc (curl error)"; fail=1; return; }
  status="$body"
  if [ "$status" != "$expected_status" ]; then
    echo "FAIL  $desc (expected $expected_status, got $status)"
    cat /tmp/smoke-body
    fail=1
  else
    echo "OK    $desc ($status)"
  fi
}

echo "== healthz =="
check "GET /healthz" 200 "http://$HOST/healthz"

echo "== reset (clean slate) =="
check "POST /reset" 200 -X POST "http://$HOST/reset"

echo "== items: empty after reset =="
check "GET /items (empty)" 200 "http://$HOST/items"
count=$(python3 -c "import json,sys; print(len(json.load(open('/tmp/smoke-body'))))")
if [ "$count" != "0" ]; then
  echo "FAIL  expected 0 items after reset, got $count"
  fail=1
else
  echo "OK    0 items after reset"
fi

echo "== items: write =="
check "POST /items" 201 -X POST -H 'Content-Type: application/json' -d '{"value":"smoke-test"}' "http://$HOST/items"
item_id=$(python3 -c "import json; print(json.load(open('/tmp/smoke-body'))['id'])")

echo "== items: read back =="
check "GET /items (1 row)" 200 "http://$HOST/items"
count=$(python3 -c "import json,sys; print(len(json.load(open('/tmp/smoke-body'))))")
if [ "$count" != "1" ]; then
  echo "FAIL  expected 1 item after write, got $count"
  fail=1
else
  echo "OK    1 item after write"
fi

echo "== items: update =="
check "PUT /items/$item_id" 200 -X PUT -H 'Content-Type: application/json' -d '{"value":"smoke-test-updated"}' "http://$HOST/items/$item_id"
value=$(python3 -c "import json; print(json.load(open('/tmp/smoke-body'))['value'])")
if [ "$value" != "smoke-test-updated" ]; then
  echo "FAIL  expected updated value, got $value"
  fail=1
else
  echo "OK    value updated"
fi

echo "== items: update missing id =="
check "PUT /items/999999999" 404 -X PUT -H 'Content-Type: application/json' -d '{"value":"x"}' "http://$HOST/items/999999999"

echo "== items: delete =="
check "DELETE /items/$item_id" 204 -X DELETE "http://$HOST/items/$item_id"

echo "== items: delete missing id =="
check "DELETE /items/$item_id" 404 -X DELETE "http://$HOST/items/$item_id"

echo "== items: empty after delete =="
check "GET /items (empty again)" 200 "http://$HOST/items"
count=$(python3 -c "import json,sys; print(len(json.load(open('/tmp/smoke-body'))))")
if [ "$count" != "0" ]; then
  echo "FAIL  expected 0 items after delete, got $count"
  fail=1
else
  echo "OK    0 items after delete"
fi

echo "== metrics =="
check "GET /metrics" 200 "http://$HOST/metrics"

echo "== reset again (cleanup) =="
check "POST /reset" 200 -X POST "http://$HOST/reset"

rm -f /tmp/smoke-body

if [ "$fail" -ne 0 ]; then
  echo
  echo "SMOKE TEST FAILED"
  exit 1
fi

echo
echo "SMOKE TEST PASSED"
