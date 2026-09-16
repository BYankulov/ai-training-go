# ai-training-go

A Docker-based Go learning playground. Each exercise lives in its own `apps/<name>/` folder and runs as its own Compose service — see `_bmad-output/specs/spec-tui-tcp-chat/conventions.md` for the full monorepo convention.

## Apps

- `apps/tui_tcp_chat` — TUI TCP chat server + Bubble Tea client. From the repo root:
  - `docker compose up` — start the server.
  - `docker compose up --build tui_tcp_chat` — rebuild and (re)start the server.
  - Connect a client: `nc localhost 9001` (or `telnet localhost 9001`), from two or more terminals. Text typed in one appears in the others.