---
description: Resume work — read persisted project context and lessons learned before doing anything else
allowed-tools: Bash, Read, Glob, Grep
---

# Project Start

You are starting (or resuming) a work session on this project. Before doing anything else, load the context left by prior sessions so you don't re-derive it from scratch or ask the user to re-explain work already captured.

## 1. Read the persisted state

Read `docs/project-context.md` and `docs/lessons-learned.md`.

- If neither exists: say so plainly — this looks like the first session using this workflow (or they haven't been created yet) — and proceed without them. Don't fabricate context to fill the gap.
- If `docs/project-context.md` exists but looks stale (e.g. "Last updated" is old, or it references branches/artifacts that no longer match reality), read it anyway but flag the staleness rather than presenting it as current fact.

## 2. Sanity-check against reality before trusting it

A context doc is a claim about the past, not a live view. Quickly verify the load-bearing parts rather than assuming they still hold:
- `git status` and `git log --oneline -5` — does the current branch/state match what "Active Work" describes?
- For any file paths named in the **Artifact Index** or **Active Work** sections, spot-check that a few actually exist (`Glob`/`Grep`), especially if the doc looks more than a few sessions old.
- If something in the doc conflicts with what you observe now, trust the live state and flag the discrepancy to the user — don't silently act on stale information.

## 3. Brief the user

Summarize back in a few sentences, not a wall of text:
- Where the project stands (from **Current State**).
- What's actively in progress and whose lane it's in, if relevant.
- Anything open/blocked that they should know before diving in.
- Any lessons from `docs/lessons-learned.md` that are directly relevant to what they're likely to do next — don't recite the whole file, just what's load-bearing right now.

Then ask what they want to work on, or wait for their next message — don't start executing a plan on their behalf just because a "Next Steps" list exists in the doc.
