# Agent instructions

Workspace-level conventions (git workflow, test/lint autonomy, readonly ops, writing voice) are loaded globally via `~/.claude/CLAUDE.md` → `agentic-os-kai/AGENTS.md`. This file covers only what's specific to this repo.

## What agent-guard is

A generic-purpose [cli-guard](https://github.com/coilysiren/cli-guard) consumer. Built for repos with external contributors, where coily's Kai-specific verbs would be inappropriate. Sibling concept to [coily](https://github.com/coilysiren/coily), separated by audience, not by mechanism.

## Scope discipline

The boundary between agent-guard and coily is the load-bearing design choice. Two rules:

- **No personal verbs.** Anything that touches Kai's homelab, vault, AWS account, or other personal infrastructure belongs in coily, not here. If a verb would not make sense to a stranger cloning a downstream repo, it does not ship here.
- **No repo-specific verbs.** agent-guard exposes the generic dev surface (`build`, `test`, `vet`, `lint`, `tidy`). Repo-specific Makefile targets stay in the downstream repo's `.agent-guard/agent-guard.yaml`.

## Dev verbs

agent-guard dogfoods itself. Once installed (`brew tap coilysiren/agent-guard https://github.com/coilysiren/agent-guard && brew install coilysiren/agent-guard/agent-guard`), route through it, not bare go:

- `agent-guard exec build`
- `agent-guard exec test`
- `agent-guard exec vet`
- `agent-guard exec lint`
- `agent-guard exec tidy`

The `.agent-guard/agent-guard.yaml` ↔ `Makefile` contract is checked by `agent-guard lint`.

## v0 API discipline

v0.x. Minor API breaks ship in `main` with a note in the commit body. Consumers pin a specific commit until v1.0.0. Once a second downstream adopter lands beyond the urfave/cli repos, lock the API and bump.

## Filing issues

One issue per discrete additive change. Every commit closes a same-repo issue with `closes #N`.

## Release + post-push

Push to `main` -> `.github/workflows/release.yml`: `mathieudutour/github-tag-action` computes semver (`default_bump: patch`, conventional commits drive minor/major), tags + cuts a GH Release, then `bump-formula` rewrites the formula's url+tag+revision line via the Contents API and pushes it back to main with a skip-CI marker. No tap dispatch. `Formula/agent-guard.rb` is the source of truth here; brew picks up the new tag from this repo on the next `brew upgrade`.

Never write the literal skip-CI token in a commit message body or you'll silently disable the release workflow on that push. GitHub greps the entire message, not just the subject line. Quote it as "skip-CI marker" or "skip CI" without brackets if you need to describe it.

Post-push: verify CI at +120s (`coily ops gh run list --repo coilysiren/agent-guard --limit 1`). Once `completed/success`: `brew upgrade coilysiren/agent-guard/agent-guard`.

## See also

- [README.md](README.md) - human-facing intro.
- [docs/FEATURES.md](docs/FEATURES.md) - inventory of what ships today.
- [.agent-guard/agent-guard.yaml](.agent-guard/agent-guard.yaml) - allowlisted commands.

Cross-reference convention from [coilysiren/agentic-os#59](https://github.com/coilysiren/agentic-os/issues/59).