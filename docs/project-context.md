# Project Context — ai-training-go

_Last updated: 2026-09-16_

## Current State
`ai-training-go` is a Go monorepo built through BMad planning workflows. The first app being built is `tui_tcp_chat` (a TCP chat server + Bubble Tea TUI client), broken into 5 implementation stories in `stories.yaml`. Stories 1 (scaffolding) and 2 (TCP broadcast server) are complete and committed; stories 3–5 are fully scoped but not started.

This session ran Story 2 end-to-end via `bmad-build`: drafted the spec, implemented the server, verified it manually with two concurrent `nc` clients through Docker, ran a review pass, patched the real findings, and committed. Then pushed `main` to `origin`.

## Active Work
- **tui_tcp_chat** — Story 1 (Scaffolding) done, committed `b1626f7`. Story 2 (TCP broadcast server) done, committed `2ab86b1`, pushed to `origin/main`.
- Next up: Story 3 — Bubble Tea TUI client. Not yet started.
- No multi-person lane/ownership info found in the spec; treat as a single-track effort unless told otherwise.

## Key Decisions
- Concurrency model for the chat server is fixed by SPEC.md: goroutine-per-connection + a single broadcaster goroutine coordinating via channels, **no mutex anywhere**. Implemented in Story 2 exactly this way — see `apps/tui_tcp_chat/main.go`.
- Bubble Tea is the required TUI library for the client — no substitute. (Story 3, `stories.yaml`)
- Username is a display label only, not identity/auth — no uniqueness enforcement or validation beyond non-empty input. (Story 4, `stories.yaml`)
- Repo uses a single root `go.mod` for the whole monorepo (not per-app), and Docker Compose build context is the repo root — per `conventions.md`.
- 2026-09-16: User prefers BMad build work to pause for review after every story rather than running straight through all 5 — this was followed for Story 2 (spec approval checkpoint + review pass before committing) and should continue for Story 3+.
- 2026-09-16: Server listens on container port `9000` internally; `compose.yaml` maps host `9001`→container `9000` (not `9000`→`9000`) because host port 9000 is already bound by an unrelated local `portainer` container on this dev machine. Story 3's client must dial `9000` inside the Compose network, or `localhost:9001` if run from the host outside Compose.
- 2026-09-16: Newline-delimited text protocol, no self-echo (sender's own TTY already shows what was typed), and drop-on-full-buffer backpressure (a slow/dead client's full send buffer causes its own messages to be dropped rather than blocking broadcast to everyone else) — all decided during Story 2 since `SPEC.md` didn't specify them; documented in `stories/2-tcp-broadcast-server.md`.

## Open Questions / Blockers
- None currently blocking. Stories 3–5 are fully scoped in `stories.yaml` with per-story dev instructions (`invoke_dev_with`).

## Next Steps
1. Story 3 — Bubble Tea TUI client: dials the server over TCP (port `9000` in-network), sends typed text, renders incoming messages live. Runs via `docker compose run --rm tui_tcp_chat_client`. Pause for a checkpoint/review after it lands rather than continuing straight into Story 4.
2. Story 4 — Usernames.
3. Story 5 — Compose lifecycle finish (final integration + `docker compose up` end-to-end check).

## Artifact Index
- `_bmad-output/planning-artifacts/briefs/brief-ai-training-go-2026-09-14/brief.md` — product brief for ai-training-go
- `_bmad-output/specs/spec-tui-tcp-chat/SPEC.md` — spec for the tui_tcp_chat app
- `_bmad-output/specs/spec-tui-tcp-chat/conventions.md` — repo/app conventions (go.mod, Dockerfile, compose layout)
- `_bmad-output/specs/spec-tui-tcp-chat/stories.yaml` — 5-story implementation breakdown
- `_bmad-output/specs/spec-tui-tcp-chat/stories/1-scaffolding.md` — Story 1 detail (done)
- `_bmad-output/specs/spec-tui-tcp-chat/stories/2-tcp-broadcast-server.md` — Story 2 detail (done), includes full review triage log
- `_bmad-output/implementation-artifacts/deferred-work.md` — deferred items, incl. missing test coverage for the broadcaster's concurrency logic (flagged in Story 2 review) and a Dockerfile `go.sum` gap for Story 3
- `apps/tui_tcp_chat/` — app source: `main.go` (real TCP broadcast server as of Story 2), `Dockerfile`
- `compose.yaml` — `tui_tcp_chat` service, host `9001`→container `9000`
- `.claude/commands/project/start.md`, `stop.md` — session start/stop workflow commands
