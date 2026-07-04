# Frontend Diff Review — Panel Drawer & Delete Mutations

## Overview

This review covers 6 scoped files plus the new `base-store.svelte.ts`. The changes lift panel drawer state from local `$state` into a shared context-based store, adjust delete mutations, and wire `setBaseState()` into the `[project]` layout.

---

### 2c. `invalidateQueries` replaces optimistic filter — intentional UX trade-off

```ts
// OLD (optimistic): instant removal
queryClient.setQueryData(key, (rows) =>
  rows.filter((r) => r.id !== service_id),
);

// NEW (invalidate): refetch, stale data visible during refetch
queryClient.invalidateQueries({ queryKey: key });
```

`invalidateQueries` marks stale and refetches in background. Until the refetch resolves, the old list (still containing the deleted service) is shown. The old `setQueryData` filter removed it instantly.

If API latency is high, the deleted card flickers back briefly. If this matters, combine both:

```ts
queryClient.setQueryData(key, (rows) =>
  rows?.filter((r) => r.id !== service_id),
);
queryClient.invalidateQueries({ queryKey: key });
```

### 2d. Guard added — improvement

New `if (!instance.current.id) return;` guard before `invalidateQueries`. Old code used `instance.current.id as string` (unsafe cast). This is a strict improvement.

---

## 4. `psql-service.svelte` — Props cleanup

### 4a. `drawerOpen` prop removed cleanly

- Close button: `onclick={() => base.setPanelDrawerState(false)}` ✓
- `StreamLogs`: `open={base.inlinePanelDrawer}` ✓

No dangling references. Clean removal.

### 4b. StreamLogs lifecycle

`StreamLogs` effect cleanup runs `closeStream()` + `teardownTerminal()` when `open` goes false. Reopening the drawer creates a fresh xterm instance and reconnects SSE. This means:

- Log scrollback is lost on close/reopen
- SSE reconnection overhead on each open

**Not a bug** — consistent with "fresh view each time" behavior. If persistent logs are desired later, hide the terminal instead of destroying it.

### 4c. Query runs even when drawer is closed

`PsqlService` renders inside `InlinePanel`. The panel uses `translate-x-full` when closed but still renders children. So `useGetPsqlServiceDetailsQuery` fires when `+page.svelte` mounts, before the drawer opens. This pre-warms data and is the existing pattern. Not a bug (only one service renders at a time).

---

## 5. `+page.svelte` — Binding to base store

### 5a. `bind:open={base.inlinePanelDrawer}` — correct

`InlinePanel` has `open = $bindable(false)`. When the overlay is clicked, it writes `open = false`, which propagates to `base.inlinePanelDrawer` via `$state`'s bindability. The X button also sets `base.setPanelDrawerState(false)`. Both paths work correctly.

### 5b. Service select ordering — correct

```ts
selectedServiceId = service.id; // set before drawer opens
selectedServiceType = service.type;
base.setPanelDrawerState(true); // opens drawer
```

The drawer renders `PsqlService` which immediately queries. ID is set first — no race.

---

## 6. `+layout.svelte` — Base state initialization

### 6a. Timing

```svelte
useGetAllInstanceQuery(() => project);  // returns query observer (sync)
setBaseState();                          // runs immediately after
```

`{@render children()}` follows both. Base state is available before any child renders. No race.

### 6b. `import { setBaseState } from '@/features/base/base-store.svelte.js'`

Already covered in §1b. Should be `from '@/features/base'` once barrel export is added.

---

## 7. Negative space — Complete migration check

✅ Zero remaining `drawerOpen` references after diff is applied.
✅ All drawer state consumers use base store.
✅ No orphaned `drawerOpen` props being passed.

No additional migration needed.

---

## Summary

| #   | Severity       | Issue                                                                    | Location                     |
| --- | -------------- | ------------------------------------------------------------------------ | ---------------------------- |
| 1   | **🔴 BUG**     | `goto('.')` is a no-op — user stays on deleted service page              | `mutation.svelte.ts:144`     |
| 2   | **🟡 BUG**     | Toast fires before guard check; cache may not invalidate                 | `mutation.svelte.ts:133-139` |
| 3   | **🟡 MISSING** | `base-store.svelte.ts` not exported from barrel                          | `base/index.ts`              |
| 4   | **🔵 LOW**     | Inconsistent import paths (`.js` vs no extension, direct file vs barrel) | 4 files                      |
| 5   | **🔵 LOW**     | `setBaseState()` no guard against re-initialization                      | `+layout.svelte:9`           |
| 6   | **⚪ INFO**    | `invalidateQueries` may flash stale data (intentional)                   | `mutation.svelte.ts:157-159` |
| 7   | **⚪ INFO**    | `selectedServiceId`/`selectedServiceType` not cleared on drawer close    | `+page.svelte:24-25`         |

### Fix priority

1. **Fix `goto('.')` → `goto('..')`** — user-facing navigation regression.
2. **Add barrel export** for `base-store.svelte.ts`.
3. **Reorder toast + guard** in `useDeleteAppServiceMutation`.
4. Normalize imports once barrel is available.
5. (Optional) Add init guard to `setBaseState()`.
6. (Optional) Clear selected service IDs on drawer close.
