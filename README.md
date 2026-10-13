# Agentic starter kit

Claude Code (plan/architect, mobile, escalation) + Cursor (desktop implementation), with shared rules, a 7-agent roster, 14 skills, hooks/permissions, CI/CD templates, and the confirmed default stack:
**Cloudflare + Neon (Postgres/Auth) + Go on AWS Lambda + R2** (Supabase swap when an app needs better social auth/MFA/SAML).

Start with `docs/decisions/0001-default-stack.md` (confirmed vs proposed vs open items).

## Install into a repo
```bash
cp -r AGENTS.md CLAUDE.md .claude .cursor docs templates  /path/to/your-repo/
chmod +x /path/to/your-repo/.claude/hooks/*.sh
# requires: jq, git, pnpm, go (for hooks)
```
Then fill the TODOs in `AGENTS.md`, copy `my-design/references/tokens-template.md` to `docs/design/tokens.md`, and fill `docs/runbooks/environments.md`.
Skills you want available in **every** repo can also live in `~/.claude/skills/`.

## Rollout order (best value first)
1. AGENTS.md + CLAUDE.md, permissions + hooks (`.claude/settings.json`, `.claude/hooks/`).
2. Remote Control (`docs/runbooks/mobile-remote-control.md`).
3. Agents + `plan-feature`, `handoff-*`, `review-diff`.
4. `my-design` + `ui-review` (+ optional third-party UI skills below).
5. `stack-decision` → `bootstrap-project` on your first app; CI workflows.
6. Grow `add-*`, `debug-triage`, `security-pass` from real corrections.

## Optional third-party skills / MCP (read every SKILL.md first; pin versions/commits)
- Anthropic `frontend-design` (base for `my-design`), Vercel `web-design-guidelines` + `react-best-practices`, an accessibility skill (Addy Osmani web-quality-skills or AccessLint), shadcn skill. Install via the skills CLI, e.g. `npx skills add <owner/repo@skill>` — verify the exact source and version.
- Neon agent skills: the Neon docs show `neon skills -s neon -y` (verify).
- Playwright MCP: `claude mcp add playwright -- npx @playwright/mcp@<pinned>`
- Chrome DevTools MCP: `claude mcp add chrome-devtools --scope user -- npx chrome-devtools-mcp@<pinned>`

## Customising
- `.claude/hooks/blocked-patterns.txt` — add/remove blocked command patterns.
- `.claude/settings.json` — widen/narrow the allowlist as you trust workflows.
- Agent models/tools: edit the frontmatter in `.claude/agents/*.md`.
- `my-design/SKILL.md` — your UI taste; add a rule whenever you correct the same thing twice.

## Verify after install
- Ask an agent to read `.env` (should be denied) and to run `git push --force` (should be blocked).
- Run `/agents` and confirm the roster loads; check frontmatter fields against your Claude Code version.
- This kit is a draft: Go skeleton and workflows are references — compile/test them and pin all versions and Action SHAs.
