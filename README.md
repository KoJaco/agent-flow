# agent-flow: agentic project template

GitHub template for projects built with Claude Code (plan/architect, mobile, escalation) and Cursor (desktop implementation), with shared rules, skills, CI/CD, and the default stack:
**Cloudflare + Neon (Postgres/Auth) + Go on AWS Lambda + R2** (Supabase swap when an app needs better social auth/MFA/SAML).

Start with `docs/decisions/0001-default-stack.md` (confirmed vs proposed vs open items).

## Use this template
1. On GitHub: **Use this template** (or `gh repo create <name> --template <owner>/agent-flow --private`).
2. Prerequisites on your machine: `jq`, `git`, `gh`, `pnpm`, `go`, `node` (hooks fail closed without `jq`).
3. Run the `stack-decision` skill, then `bootstrap-project` (checklist: `.claude/skills/bootstrap-project/references/checklist.md`).
4. Fill in `docs/design/tokens.md` and `docs/runbooks/environments.md`.
5. Configure Cursor command approval and sandbox: `docs/runbooks/cursor-guardrails.md` (not covered by `.claude/settings.json`).
6. Set branch protection on `main` (require PR, review, `ci-ok`) and create the `staging` / `production` Environments.

## Layout
| Path | Purpose | Read by |
|---|---|---|
| `AGENTS.md` | Conventions and hard rules (single source of truth) | Claude + Cursor |
| `CLAUDE.md` | Imports `AGENTS.md`, Claude-specific roles | Claude |
| `.cursor/rules/project.mdc` | Thin pointer to `AGENTS.md` | Cursor |
| `.claude/skills/` | All skills, committed (cloud agents can't see `~/.agents`) | Claude + Cursor |
| `.claude/agents/`, `.claude/hooks/`, `.claude/settings.json` | Subagents, guard hooks, permissions | Claude only |
| `.mcp.json`, `.cursor/mcp.json` | MCP servers; keep identical, versions pinned | per tool |
| `.github/` | CI (`ci-ok` required check), deploy/preview workflows, CODEOWNERS, PR template, Dependabot | GitHub |
| `docs/` | ADRs, handoffs, runbooks, design tokens | everyone |

Skills whose description starts with `Claude Code only:` orchestrate subagents; Cursor should not run them. The rest are tool-neutral.

## Guardrails
- Claude: `.claude/settings.json` allow/ask/deny + `block-dangerous.sh` (patterns in `blocked-patterns.txt`).
- Cursor: **manual setup**, see `docs/runbooks/cursor-guardrails.md`.
- Both: CI, branch protection, staging-only credentials, CODEOWNERS.
- Verify after creating a repo: ask Claude to read `.env` (denied) and to force-push (blocked).

## Workflows
- `ci.yml` self-skips until `package.json` / `api/go.mod` exist; require the `ci-ok` check.
- `deploy-staging.yml`, `deploy-prod.yml`, `preview-db.yml` are inert until their repo variables are set (see the bootstrap checklist). Prod deploys sit behind the `production` Environment approval.
- Actions are pinned by SHA; Dependabot keeps them current. Compile/test the Go skeleton in `bootstrap-project/assets/go-api/` before relying on it.

## Third-party skills
Committed under `.claude/skills/` and tracked in `skills-lock.json` (Neon, Vercel React/composition, web-design-guidelines, accessibility, humanizer). Read each `SKILL.md` before updating; they can instruct agents to run code.

## Customising
- `.claude/hooks/blocked-patterns.txt`: add/remove blocked command patterns.
- `.claude/hooks/post-edit.sh`: formats and runs a scoped check per edit; `AGENT_HOOK_SKIP_CHECKS=1` formats only.
- `.claude/settings.json`: widen/narrow the allowlist as you trust workflows.
- Agent models/tools: frontmatter in `.claude/agents/*.md`.
- `my-design/SKILL.md`: your UI taste; add a rule whenever you correct the same thing twice.
