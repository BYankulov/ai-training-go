---
id: SPEC-tui-tcp-chat
companions: [conventions.md]
sources: [../../planning-artifacts/briefs/brief-ai-training-go-2026-09-14/brief.md]
---

> **Canonical contract.** This SPEC and the files in `companions:` are the complete, preservation-validated contract for what to build, test, and validate. Source documents listed in frontmatter are for traceability — consult them only if you need narrative rationale or prose color this contract intentionally omits.

# TUI TCP Chat (v1)

## Why

This is a vision to realize: a hands-on vehicle for learning Go by building something real rather than working through isolated tutorials. It is the first app in the `ai-training-go` monorepo, a personal Docker-based learning playground. In priority order, it exists to build understanding of goroutines and channels, the `net` package, and project structure/modules — in that order, learned step by step rather than all at once.

## Capabilities

- **CAP-1**
  - **intent:** Server can accept multiple concurrent TCP client connections and broadcast text messages among them.
  - **success:** 2+ clients connected simultaneously; a message sent from one client is received by all others.

- **CAP-2**
  - **intent:** User can run a Bubble Tea TUI client that connects to the chat server over TCP and exchanges messages live.
  - **success:** Running `docker compose run --rm tui_tcp_chat_client` connects to the running server; messages typed locally appear on other connected clients, and incoming messages render in this client's TUI without manual refresh.

- **CAP-3**
  - **intent:** User can set a username that tags every message they send, visible to other clients.
  - **success:** Each message rendered in any client's TUI shows the sender's username alongside the text.

- **CAP-4**
  - **intent:** Developer can start and rebuild the chat server via Docker Compose from the repo root.
  - **success:** `docker compose up` and `docker compose up --build tui_tcp_chat`, run from the repo root, bring up the server with no manual steps outside Docker.

## Constraints

- Server concurrency model is goroutine-per-connection plus a single broadcaster goroutine coordinating via channels, with no mutex — a stated learning goal, not open to implementer substitution.
- TCP networking uses Go's standard `net` package directly — no higher-level networking framework or library.
- TUI is built with Bubble Tea (Charm ecosystem) — not open for substitution.
- Lives at `apps/tui_tcp_chat/` inside the `ai-training-go` monorepo and uses its single root `go.mod` — see `conventions.md` for the full monorepo convention.
- Server and client are separate Compose services sharing the app's Dockerfile/image: `tui_tcp_chat` (server, long-running, managed with `up`/`--build`) and `tui_tcp_chat_client` (client, never `up`'d — invoked per-session with `docker compose run --rm` to get a real attached TTY). See `conventions.md`.

## Non-goals

- No chat rooms in v1 — a single shared channel only.
- No message history or persistence — messages exist only live, for the session.
- No authentication, TLS, or encryption — plain TCP; username is a display label, not an identity.
- No reconnect/resilience handling — a dropped connection simply ends that client's session.

## Success signal

v1 is done when two or more Bubble Tea TUI clients connect to the Go TCP server, exchange text messages tagged with usernames, and the whole thing starts via `docker compose up` from the repo root — no manual steps outside Docker.

## Assumptions

- Assumed that setting up the Go and Docker toolchain on the developer's machine is a prerequisite to executing this spec, not itself a spec capability — the source brief notes neither is installed yet.
- Assumed two separate Compose service names: `tui_tcp_chat` for the server (preserves the brief's literal example command `docker compose up --build tui_tcp_chat`) and `tui_tcp_chat_client` for the client. Rename before build if a different scheme is wanted.
