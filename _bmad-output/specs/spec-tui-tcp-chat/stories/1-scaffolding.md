---
title: 'Scaffolding: apps/tui_tcp_chat/, root go.mod, Dockerfile, compose skeleton'
type: 'chore'
created: '2026-09-14'
status: 'done'
route: 'oneshot'
review_loop_iteration: 0
context: ['_bmad-output/specs/spec-tui-tcp-chat/SPEC.md', '_bmad-output/specs/spec-tui-tcp-chat/conventions.md']
---

<frozen-after-approval reason="human-owned intent — do not modify unless human renegotiates">

## Intent

**Problem:** The `ai-training-go` monorepo has no code yet — no Go module, no app folder, no Docker/Compose wiring for the first learning app, `tui_tcp_chat`.

**Approach:** Scaffold the structure per `conventions.md`: a single root `go.mod` (module path `github.com/BYankulov/ai-training-go`), `apps/tui_tcp_chat/` with a placeholder `main.go` that compiles and runs (prints a startup message; no chat logic yet), its own `Dockerfile`, and a repo-root `compose.yaml` defining the `tui_tcp_chat` server service with build context at the repo root. Skeleton only — no TCP, broadcast, or TUI logic in this story; those land in stories 2–5.

</frozen-after-approval>

## Implementation Notes

- `go.mod` at repo root: module `github.com/BYankulov/ai-training-go` (matches the git remote), `go 1.23`.
- `apps/tui_tcp_chat/main.go`: placeholder prints a startup message, then blocks on SIGINT/SIGTERM via `os/signal` — not a bare `select{}`, which the Go runtime flags as `fatal error: all goroutines are asleep - deadlock!` when it is the only goroutine (caught by testing `docker compose up`).
- `apps/tui_tcp_chat/Dockerfile`: multi-stage build, `golang:1.23-alpine` builder, `alpine:3.20` runtime, binary at `/usr/local/bin/tui_tcp_chat`.
- `compose.yaml` at repo root: `tui_tcp_chat` service, `build.context: .`, `dockerfile: apps/tui_tcp_chat/Dockerfile`, per `conventions.md`. No ports yet — the server doesn't listen until story 2.
- Added `.dockerignore` (`.git`, `.claude`, `_bmad`, `_bmad-output`) since the build context is the repo root and would otherwise ship the whole BMad toolchain into the build.
- Verified: `docker compose build tui_tcp_chat`, `docker compose up tui_tcp_chat` (stays `Up`, logs the startup line), and `docker compose up --build tui_tcp_chat` all succeed from the repo root with no manual steps outside Docker.

## Review Triage Log

- `medium` — `.gitignore` had no Go build/OS/editor ignores. Real: repo has none today, would pick up stray files (`/bin/`, IDE dirs) the moment someone builds locally or adds an editor. Fixed: added Go build artifact and OS/editor entries.
- `medium` — Dockerfile runtime stage ran as root by default. Real: no `USER` was set. Fixed: added a dedicated `appuser` (uid 10001) and switched to it before `ENTRYPOINT`; verified via `docker compose exec ... id`.
- `low` — README didn't mention `apps/tui_tcp_chat` or how to run it. Real but cosmetic. Fixed: added an "Apps" section with the `docker compose up` / `up --build` commands.
- `low` — builder stage didn't set `CGO_ENABLED=0` explicitly; worked only because builder and runtime are both musl/Alpine. Fixed: added `ENV CGO_ENABLED=0` to the builder stage to remove the implicit coupling.
- `false` — frontmatter `status` looked stale (`in-progress`) next to Implementation Notes describing verified builds. Not a real finding: this file's own Finalize step (this edit) sets `status: done` as the last action, per the oneshot workflow's own sequencing.
- `false` — flagged `.dockerignore` as incomplete for hypothetical future noise (`.vscode/`, a future `bin/`). Rejected: no such files exist in this repo yet; nothing to ignore today, and adding speculative entries would be guessing at future structure.
- `low` — `main.go` uses `fmt.Println` instead of `log`. Rejected: this placeholder is fully replaced by real server logic in story 2; picking a logging convention now for throwaway scaffolding code isn't worth it.
- Deferred (see `_bmad-output/implementation-artifacts/deferred-work.md`): Dockerfile only `COPY`s `go.mod`, not `go.sum` — a real gap once story 3 adds Bubble Tea as a dependency, but not a defect in this story (zero deps today).
