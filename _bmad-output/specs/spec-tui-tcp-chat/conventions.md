# Monorepo Conventions

Standing structure for the `ai-training-go` monorepo. Applies to `tui_tcp_chat` and every learning app added after it.

## Layout

- Each learning app lives in its own `apps/<name>/` folder, e.g. `apps/tui_tcp_chat/`.
- Each app folder has its own `Dockerfile`.
- Docker Compose's `build.context` stays at the repo root so every app's `Dockerfile` can reach the shared root-level `go.mod`.

## Modules

- Single root `go.mod` for the whole monorepo — not a per-app module. Chosen for simplicity while module/workspace structure is still being learned; revisit if apps outgrow it.

## Running apps

- From the repo root: `docker compose up` starts the configured app(s).
- `docker compose up --build <service>` rebuilds and (re)starts one specific app, e.g. `docker compose up --build tui_tcp_chat`.
- **Interactive/TUI services are never `up`'d.** Plain `docker compose up` multiplexes every service's logs into one non-interactive stream — it can't hand you a real keyboard/screen session. Any service with an interactive terminal is still defined in `compose.yaml`, but is only ever launched with `docker compose run --rm <service>`, which allocates a real attached TTY. Long-running background services (a server, for example) keep using `up`/`--build`.

## Source control

- The repo is git-initialized on `main` and bound to `git@github.com:BYankulov/ai-training-go.git` over SSH.
