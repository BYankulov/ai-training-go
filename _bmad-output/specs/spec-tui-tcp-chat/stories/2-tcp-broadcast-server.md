---
title: 'TCP broadcast server: goroutine-per-connection + channel broadcaster'
type: 'feature'
created: '2026-09-16'
status: 'done'
route: 'oneshot'
review_loop_iteration: 0
context: ['_bmad-output/specs/spec-tui-tcp-chat/SPEC.md', '_bmad-output/specs/spec-tui-tcp-chat/conventions.md']
---

<frozen-after-approval reason="human-owned intent — do not modify unless human renegotiates">

## Intent

**Problem:** `apps/tui_tcp_chat/main.go` is still the story-1 placeholder — it prints a startup line and blocks on a signal, but does not listen for connections or relay anything. CAP-1 (2+ concurrent TCP clients exchanging broadcast messages) is unimplemented.

**Approach:** Replace the placeholder with a real TCP server using Go's stdlib `net` package directly: `net.Listen` accepts connections; each connection gets its own goroutine reading newline-delimited text lines; every read is forwarded over a channel to one central broadcaster goroutine, which fans each message out to every *other* connected client's per-client send channel (sender does not get its own message echoed back — the client's own TTY already shows what was typed). No mutex anywhere — the broadcaster goroutine is the sole owner of the client set. Listen on port `9000`; add a `ports: ["9000:9000"]` mapping to the `tui_tcp_chat` service in `compose.yaml` so the server is reachable from the host for `nc`/`telnet` testing before the TUI client (story 3) exists. A dropped connection simply removes that client and closes its goroutine — no reconnect logic (non-goal, per SPEC.md).

</frozen-after-approval>

## Implementation Notes

- `apps/tui_tcp_chat/main.go` rewritten: `net.Listen("tcp", ":9000")` accept loop, one goroutine per connection (`handleConn`), a single `runBroadcaster` goroutine owning `map[*client]bool` — reachable only via three channels (`register`, `unregister`, `broadcast`), so no mutex is needed anywhere.
- Per-client outbound channel is buffered (`make(chan string, 16)`); the broadcaster's fan-out uses `select { case c.send <- msg: default: }` so one slow/dead client's full buffer never blocks delivery to the rest — matches "no reconnect/resilience" non-goal by simply dropping for that client rather than adding retry/backoff logic.
- Each connection's reader uses `bufio.Scanner` (newline-delimited text) and forwards non-empty lines to the broadcaster; the broadcaster skips the message's own sender, so a client never sees its own text echoed back (its own TTY already shows what was typed via `nc`/`telnet`).
- Disconnect: reader loop ends on EOF/error → sends on `unregister` → broadcaster deletes the client and closes its `send` channel → the per-connection writer goroutine's `range c.send` exits → connection is closed. Verified clean in server logs (`client connected` / `client disconnected` pairs, no panics) across two concurrent `nc` sessions.
- Switched startup/logging from `fmt.Println` to `log` (story 1 deferred this pending real server logic, which lands here) — logs listen address plus every connect/disconnect with remote address.
- `compose.yaml`: added a `ports` mapping for the `tui_tcp_chat` service so the server is reachable from the host for `nc`/`telnet` testing ahead of the TUI client (story 3).
- **Surprise:** host port `9000` was already bound by an unrelated local container (`portainer`) on this dev machine — mapped host `9001` → container `9000` instead (`ports: ["9001:9000"]`); the server's own internal listen address is unaffected (still `:9000`), so story 3's client only needs to know the container/network-internal port unless it also runs from the host, in which case use `9001`.
- Verified end-to-end via Docker: `docker compose build`, `docker compose up -d`, two concurrent `nc localhost 9001` sessions exchanging messages (cross-delivery confirmed, no self-echo, correct ordering), then `docker compose down` and `docker compose up --build -d` again — all clean, no manual steps outside Docker.
- Post-review fixes: accept loop now distinguishes an intentional shutdown (`errors.Is(err, net.ErrClosed)`) from a transient `Accept` error (logs and keeps looping instead of killing the whole server); the writer goroutine now closes the connection on a write error so the reader side notices and unregisters promptly; `scanner.Err()` is now logged after the read loop ends, distinguishing a real read error from a normal disconnect; `compose.yaml`'s port mapping now has an inline comment explaining the 9001→9000 host/container mismatch; README's Apps section now documents connecting with `nc`/`telnet`. Re-verified with the same Docker + two-`nc`-client test after applying these — no regressions.

## Review Triage Log

- `medium` — Accept loop treated every `ln.Accept()` error as shutdown, killing the whole server (and every connected client) on a transient error, not just an intentional close. Fixed: distinguish `net.ErrClosed` from other errors; log and keep accepting on transient errors.
- `low` — Writer goroutine returned on a `Fprintln` error without closing the connection, so the reader side could stay registered indefinitely once that client's send buffer filled. Fixed: writer now closes the connection on write failure.
- `low` — `scanner.Err()` was never checked after the read loop, so a real read error was indistinguishable in the logs from a normal client disconnect. Fixed: log it when non-nil.
- `low` — `compose.yaml`'s `9001:9000` port mapping had no explanation for the mismatch. Fixed: added an inline comment.
- `low` — README's Apps section didn't say how to actually connect and exchange messages now that CAP-1 works. Fixed: added an `nc`/`telnet` line.
- `low`, rejected — No read/write deadlines on client connections, so a half-open connection (peer stops reading without closing) could leak a goroutine and a phantom registered client indefinitely. Real, but rare to hit via manual `nc`/`telnet` testing, and a correct fix needs an arbitrary idle-timeout policy not specified anywhere in `SPEC.md` — more than a simple correction. Left as-is; a candidate for the same non-goal territory as "no reconnect/resilience handling."
- `defer` — No unit/integration tests for the broadcaster's concurrency logic (self-echo suppression, drop-on-full-buffer, register/unregister). Real coverage gap for stories 3-5 to build on; recorded in `_bmad-output/implementation-artifacts/deferred-work.md`.
- `false` — Listen port is a hardcoded constant with no env/flag override, framed as a gap exposed by the story's own port-conflict workaround. Disproven: Docker Compose's `ports` mapping already decouples host port from container port — the actual conflict (host 9000 taken by an unrelated `portainer` container) was fully resolved via `ports: ["9001:9000"]` with zero code changes; an app-level override would solve nothing the compose mapping doesn't already solve.
- `false` — No graceful shutdown draining or notifying connected clients on SIGINT/SIGTERM. Disproven: `SPEC.md`'s non-goals already state "no reconnect/resilience handling — a dropped connection simply ends that client's session," which covers server-initiated shutdown ending connections without ceremony, not just client-side drops.

