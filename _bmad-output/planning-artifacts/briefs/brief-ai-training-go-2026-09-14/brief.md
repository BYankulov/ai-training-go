---
title: "Product Brief: ai-training-go — TUI TCP Chat (v1)"
status: final
created: 2026-09-14
updated: 2026-09-14
---

# Product Brief: ai-training-go — TUI TCP Chat (v1)

## What This Is

ai-training-go is a personal Go-learning monorepo: a Docker-based playground where each learning exercise lives in its own `apps/<name>/` folder and runs as its own Compose service. The first concrete app is `apps/tui_tcp_chat/` — a minimal TUI TCP chat, with both server and client written in Go and built with Bubble Tea.

## Why

This exists to learn Go by building something real, not just working through isolated tutorials. In priority order, the goal is to understand:

1. **Goroutines and channels** — the chat server's core concurrency model (a goroutine per connection, a broadcaster goroutine coordinating everything through channels).
2. **The `net` package** — raw TCP socket programming: listening, accepting, dialing, reading and writing over `net.Conn`.
3. **Project structure and modules** — approached step by step as the monorepo grows, not all at once.

No deadline — this proceeds at whatever pace fits.

## The App (v1 Scope)

- Server and client, both written in Go, communicating over raw TCP.
- TUI built with **Bubble Tea** (Charm ecosystem).
- Multiple clients can connect to one server; a single broadcaster goroutine coordinates them via channels — no rooms, no mutex — connection state stays confined to its own goroutine.
- One feature beyond raw text: **usernames** — each message is tagged and displayed with the sender's chosen name.

**Explicitly deferred to later versions** (not rejected — just not v1):
- Message history / persistence
- Authentication, TLS, encryption
- Reconnect / resilience handling on dropped connections

## Infrastructure

- From the repo root: `docker compose up` starts the app(s); `docker compose up --build tui_tcp_chat` rebuilds and (re)starts this service specifically.
- Each learning app gets its own `apps/<name>/` folder with its own Dockerfile. The Compose build context stays at the repo root so every app's Dockerfile can reach the **single root `go.mod`** — one module for the whole monorepo rather than per app, chosen for simplicity while module/workspace structure is still being learned (revisit later if apps outgrow it).
- Go and Docker are **not yet installed/working locally** — getting the toolchain running is itself part of the learning curve here, not an assumed starting point.
- The repo is tracked in git locally and bound to `git@github.com:BYankulov/ai-training-go.git` over SSH (initial planning setup already pushed to `main`).

## Success Criteria

v1 is done when two or more Bubble Tea TUI clients connect to the Go TCP server, exchange text messages tagged with usernames, and when the whole thing starts with `docker compose up` from the repo root — no manual steps outside Docker.

## What's Next

Once `tui_tcp_chat` works end to end, the same `apps/<name>/` + Compose pattern repeats for whatever comes next in the Go learning course — each new exercise gets its own folder and service without disturbing what's already working. No specific next app is decided yet.
