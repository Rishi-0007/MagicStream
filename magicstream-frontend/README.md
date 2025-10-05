# MagicStream Frontend — P0 (Next.js 15 + Tailwind v4)

Cinematic UI foundation with design tokens, typography, and layout. Ready for shadcn components, TanStack Query, and Axios integration.

## Stack
- **Next.js 15** (App Router, React 19) — stable.
- **Tailwind CSS v4** — config-less with `@theme` tokens.
- **shadcn-style primitives** — Button, Input, Card, Skeleton, Badge.
- **TanStack Query** — data layer.
- **Axios** — API client.
- **next-themes** — dark/system themes via `class`.
- **TypeScript**, **ESLint** (Next core web vitals).

## Quick start
```bash
# 1) install deps
pnpm install   # or: npm i / yarn

# 2) env
cp .env.example .env.local

# 3) dev
pnpm dev       # http://localhost:3000
```

## Design system (P0)
- Color tokens defined in `src/app/globals.css` with Tailwind v4 `@theme`:
  - `bg`, `surface`, `elevated`, `border`, `text`, `muted`
  - `primary`, `secondary`, `accent`, `success`, `danger`
- Dark theme token overrides in `html.dark { ... }`.
- Typography: **Inter** (UI) + **Plus Jakarta Sans** (headings).
- Radii, shadows, selection, focus ring are consistent via tokens.

Use tokens directly in Tailwind:
```tsx
<div className="bg-surface text-text border border-border rounded-[var(--radius)]" />
```

## Providers
- `ThemeProvider` (`next-themes`) — toggles `html.dark`.
- `QueryProvider` — configures TanStack Query.

## What’s next (P1)
- Auth pages (login/register) with Zod + React Hook Form.
- Axios interceptors + token refresh integration with backend.
- Search and catalog pages (P2).

## Notes
- Tailwind v4 theming/tokens for colors use `@theme` (top-level). Dark mode overrides variables in `.dark` — utilities like `bg-primary` read updated values at runtime.
