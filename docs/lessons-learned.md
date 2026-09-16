# Lessons Learned — ai-training-go

Durable, non-obvious lessons only — not a session recap. Newest first.

## 2026-09-16
- **Lesson:** Host port `9000` is already bound by an unrelated local `portainer` container on this dev machine — don't map any Compose service's host port to `9000`.
  **Context:** `compose.yaml` originally mapped `9000:9000` for `tui_tcp_chat`; `docker compose up` failed with "port is already allocated". Remapped to `9001:9000` (host:container) — the app's internal listen port is still `9000`, only the host-side mapping changed.
- **Lesson:** The Go toolchain is not installed on this dev machine — `go build`/`go run`/`go test` all fail with "command not found". Build and test only through Docker (`docker compose build`, then run the container).
  **Context:** Tried `go version` while planning how to test Story 2 quickly; it wasn't found. `docker` was available and used for every build/test cycle instead.
