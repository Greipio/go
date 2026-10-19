# Changelog

## v2.0.0 — 2026-03-19

### Breaking

- Module import path is now **`github.com/greipio/go/v2`** (Go module major version).
- **Greip API v2** base URL, Bearer auth, and `success` / `data` response handling.
- REST paths updated for lookup, IP score (threats), bulk lookup, country, ASN, email, phone, IBAN, profanity, and payments.
- Response structs aligned with v2 JSON (`ip`, `geo`, `security`, snake_case validation fields, payment `score` as float).

### Added

- `cmd/livecheck` — live smoke test binary (requires `GREIP_TOKEN`).
