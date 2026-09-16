---
description: Wrap up the work session — persist context and lessons learned for a clean resume, then clear context
argument-hint: "[optional: anything to flag for next time]"
allowed-tools: Bash, Read, Write, Edit, Glob, Grep
---

# Project Stop

You are closing out a work session on this project. Your job is to make sure nothing important is lost, and that the *next* session — which will have zero memory of this conversation — can pick up cleanly from disk alone. Do not skip steps because "nothing much happened"; a short, accurate update is still the point.

## 1. Gather what actually happened this session

- `git status` and `git diff` (staged + unstaged) to see what changed on disk.
- `git log --oneline -10` for recent commits, in case this session included commits.
- Check for BMad workflow activity: any `_bmad-output/**/.memlog.md` files touched this session (these are the authoritative decision trail for PRDs/architecture/etc. in progress — read the tail of any that changed).
- Review the conversation itself for: decisions made, dead ends ruled out, open questions raised, things the user corrected you on, and anything explicitly deferred ("later," "not now," "leave that for next time").
- Note the user's extra input if provided: "$ARGUMENTS"

## 2. Read the existing state (don't overwrite blind)

Read `docs/project-context.md` and `docs/lessons-learned.md` if they exist. If either is missing, create it — this is the first stop.

## 3. Update `docs/project-context.md`

This file's only job is letting a fresh session (or a teammate on this 4-person team) resume without re-deriving anything from git archaeology or asking the user to re-explain. Structure:

```markdown
# Project Context — {project name}

_Last updated: {date}_

## Current State
1-2 paragraphs: where the project stands right now, in plain terms.

## Active Work
- What's in progress, with paths to the relevant artifacts (PRDs, branches, workspaces). Mark whose lane it's in if known (FE/BE/AI/Infra) — this repo's 4-person team works individually within sprints, so ownership matters.

## Key Decisions
- Append-only log: decision — why — where it's documented in full (e.g. a PRD section or memlog). Don't delete old entries; if a decision was later reversed, add a new entry noting the supersession rather than rewriting history.

## Open Questions / Blockers
- Anything unresolved that the next session should know about before making assumptions. Point at where it's tracked in full (e.g. a PRD's Open Questions section) rather than duplicating it.

## Next Steps
- Ordered, concrete. What should happen next and why.

## Artifact Index
- path — one-line description, for every meaningfully persistent thing produced (PRDs, architecture docs, key scripts, etc.)
```

Update every section to reflect current reality — this is a living document, not an append-only log (except **Key Decisions**, which is append-only). If something in the existing file is now stale or wrong, fix it; don't leave it for someone to trip over.

## 4. Update `docs/lessons-learned.md`

Append **only genuinely new, durable lessons** — not a session recap. A lesson earns a place here if it would change how a future session (human or agent) approaches this specific project. Good candidates: a mistake made and corrected (with why it happened), a non-obvious constraint discovered, an approach the user explicitly confirmed or rejected and why, a gotcha in this codebase/toolchain that isn't derivable by reading the code. Skip this step entirely if nothing this session rises to that bar — don't pad it.

Format, newest first:

```markdown
## {date}
- **Lesson:** one-line statement of the rule.
  **Context:** what happened that taught it.
```

## 5. Clean up loose ends

- Check for any TODO items (if a todo list exists in this session) that are stale, abandoned, or already done but unmarked — resolve or note them in Open Questions rather than leaving them dangling.
- Flag (don't silently delete) any scratch/temp files this session created that won't be needed next time.
- If there are uncommitted changes the user hasn't addressed, surface that clearly — don't let "wrap up" quietly imply "and I committed it for you."

## 6. Close out

Show the user a short summary (a few bullets, not a wall of text) of what was written to `docs/project-context.md` and `docs/lessons-learned.md`. Then tell them plainly: everything needed to resume is now on disk, and it's safe to run `/clear` (or end the session) — you cannot clear the conversation yourself, so say this as a suggestion to them, not something you're about to do.
