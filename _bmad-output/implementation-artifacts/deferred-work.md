- source_spec: `_bmad-output/specs/spec-tui-tcp-chat/stories/1-scaffolding.md`
  summary: apps/tui_tcp_chat/Dockerfile only COPYs go.mod, not go.sum — story 3 will add Bubble Tea as a dependency and produce a go.sum the Dockerfile currently won't copy.
  evidence: No go.sum exists yet (zero external deps in this story), so nothing breaks today; flagged so story 3's implementer remembers to update the COPY line when adding dependencies.
- source_spec: `_bmad-output/specs/spec-tui-tcp-chat/stories/2-tcp-broadcast-server.md`
  summary: No automated unit/integration tests for the broadcaster's concurrency logic (self-echo suppression, drop-on-full-buffer backpressure, register/unregister lifecycle).
  evidence: Verified only by manual two-`nc`-session testing this story. Real risk once stories 3-5 build on top (usernames, TUI client) without regression coverage for the core broadcast mechanism.
