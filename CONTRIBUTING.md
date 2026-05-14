# Contributing to agent-guard

Thank you for your interest! :wave:

This project is run on volunteer time, so please have patience.

## Before you open a PR

1. **Open an issue first.** Every commit in this repo closes a same-repo issue (`closes #N` in the commit body). Discussion happens in the issue, the PR is the change itself. This applies even to trivial fixes, the issue gives the change a stable URL.
2. **Stay close to scope.** agent-guard is intentionally small. It exposes a generic dev surface on top of [cli-guard](https://github.com/coilysiren/cli-guard). Features that pull this package out of its lane will get pushed back. Personal verbs belong in [coily](https://github.com/coilysiren/coily), repo-specific verbs belong in the downstream repo.
3. **Run the dev verbs before pushing.** Local dev routes through [`coily`](https://github.com/coilysiren/coily):

   ```
   coily exec build
   coily exec test
   coily exec vet
   coily exec lint
   ```

   The `.coily/coily.yaml` ↔ Makefile contract is checked by `coily lint` and by CI on every push.

## Code of Conduct

Participation in this community is governed by the [Code of Conduct](CODE_OF_CONDUCT.md), adapted from the [Contributor Covenant 2.1](https://www.contributor-covenant.org/version/2/1/code_of_conduct/).

## Security disclosures

See [SECURITY.md](SECURITY.md). Do not file vulnerabilities as public issues.
