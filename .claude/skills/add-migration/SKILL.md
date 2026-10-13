---
name: add-migration
description: Create a safe Postgres schema migration (tables, columns, indexes, constraints, backfills) with rollback, tested on a disposable Neon branch. Use for any database schema or data change, including "add a column", "new table", or "fix this data".
---
# add-migration

Delegate execution to the `migration-guard` subagent.

Rules:
- Plain SQL in `db/migrations/NNNN_<slug>.sql` (goose: `-- +goose Up` / `-- +goose Down`). **Portable Postgres only** — no provider-specific features (so we can move Neon ⇄ Supabase).
- **Expand/contract:** add nullable → backfill in batches → enforce → drop old in a later release. Never rename/drop in the same release code still reads.
- Indexes on big tables: `CREATE INDEX CONCURRENTLY` (needs `-- +goose NO TRANSACTION`).
- Add `NOT NULL`/FKs with `NOT VALID` then `VALIDATE` for large tables.
- Test on a throwaway branch of the **staging** Neon project: up → tests → down → up.
- Report lock risk, duration estimate, rollback plan, and data-loss risk. Production application happens in CI after approval.
- Don't bake auth-provider coupling (e.g. RLS policies using a provider's `auth.uid()`) into schema until an ADR approves it.
