# Project Context — ai-training-go

_Last updated: 2026-09-16_

## Current State
`ai-training-go` is a Go monorepo built through BMad planning workflows. The first app being built is `tui_tcp_chat` (a TCP chat server + Bubble Tea TUI client), broken into 5 implementation stories in `stories.yaml`. Story 1 (scaffolding) is complete and committed; stories 2–5 are fully scoped but not started.

This session did no BMad story work. Its only activity was installing two custom slash commands — `/project:start` and `/project:stop` — copied from an external `ai-delivery-platform` template into `.claude/commands/project/`, to give future sessions this same context-persistence workflow.

## Active Work
- **tui_tcp_chat** — Story 1 (Scaffolding) done: `apps/tui_tcp_chat/` scaffolded (main.go stub, Dockerfile), root `go.mod`, `compose.yaml` skeleton. Committed in `b1626f7`.
- Next up: Story 2 — TCP broadcast server. Not yet started.
- No multi-person lane/ownership info found in the spec; treat as a single-track effort unless told otherwise.

## Key Decisions
- Concurrency model for the chat server is fixed by SPEC.md: goroutine-per-connection + a single broadcaster goroutine coordinating via channels, **no mutex anywhere**. (Story 2, `stories.yaml`)
- Bubble Tea is the required TUI library for the client — no substitute. (Story 3, `stories.yaml`)
- Username is a display label only, not identity/auth — no uniqueness enforcement or validation beyond non-empty input. (Story 4, `stories.yaml`)
- Repo uses a single root `go.mod` for the whole monorepo (not per-app), and Docker Compose build context is the repo root — per `conventions.md`.
- 2026-09-16: User prefers BMad build work to pause for review after every story rather than running straight through all 5 — factor this into how Story 2+ gets kicked off.

## Open Questions / Blockers
- None currently blocking. Stories 2–5 are fully scoped in `stories.yaml` with per-story dev instructions (`invoke_dev_with`).

## Next Steps
1. Kick off Story 2 (TCP broadcast server) via the dev workflow; pause for a checkpoint/review after it lands rather than continuing straight into Story 3.
2. Story 3 — Bubble Tea TUI client.
3. Story 4 — Usernames.
4. Story 5 — Compose lifecycle finish (final integration + `docker compose up` end-to-end check).

## Artifact Index
- `_bmad-output/planning-artifacts/briefs/brief-ai-training-go-2026-09-14/brief.md` — product brief for ai-training-go
- `_bmad-output/specs/spec-tui-tcp-chat/SPEC.md` — spec for the tui_tcp_chat app
- `_bmad-output/specs/spec-tui-tcp-chat/conventions.md` — repo/app conventions (go.mod, Dockerfile, compose layout)
- `_bmad-output/specs/spec-tui-tcp-chat/stories.yaml` — 5-story implementation breakdown
- `_bmad-output/specs/spec-tui-tcp-chat/stories/1-scaffolding.md` — Story 1 detail (done)
- `apps/tui_tcp_chat/` — app source (main.go stub, Dockerfile) from Story 1
- `.claude/commands/project/start.md`, `stop.md` — session start/stop workflow commands (added this session)
