## What / why
<!-- One paragraph. Link the handoff or ADR: docs/handoffs/<task>.md -->

## How to test
<!-- Commands and manual steps. CI must be green before merge. -->

## Screenshots
<!-- UI changes: 390 / 768 / 1280px, light + dark if supported. Delete if no UI change. -->

## Migration & rollback
<!-- Migration file(s), lock risk, up/down/up evidence on a throwaway Neon branch. Rollback plan. "None" if no migration. -->

## Risks
<!-- Auth, data loss, perf, behaviour change. -->

## Cost impact
<!-- New AWS/Cloudflare/Neon resources? No NAT/ALB/always-on Fargate/VPC Lambda without an ADR. Log retention set? "None" if no change. -->

## Follow-ups
<!-- Unrelated problems noticed but intentionally not fixed. -->

## Checklist
- [ ] Behaviour change has a test
- [ ] Lint, typecheck, tests pass locally
- [ ] No secrets, `.env*`, or production targets touched
- [ ] Portability kept: `DATABASE_URL` only, JWKS via `auth.Verifier`, plain SQL migrations
- [ ] UI changes reviewed with `ui-review` (WCAG 2.2 AA)
- [ ] Diff is scoped to the plan; no drive-by refactors
