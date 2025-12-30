# Copilot Instructions for Astro Website

## Project Overview

Astro 5.x static site with strict TypeScript, file-based routing, and **Tailwind CSS + SCSS** styling. Zero JavaScript by default; all `.astro` components render to static HTML at build time.

## Architecture & File Structure

- **`src/pages/`** - File-based routing (e.g., `pages/about.astro` → `/about`)
- **`src/components/`** - Reusable `.astro` components; always scoped CSS/SCSS
- **`src/layouts/`** - Base layouts wrapping page content via `<slot />`
- **`src/assets/`** - Static images/media; import with named imports (e.g., `import logo from '../assets/logo.svg'`) then use `{logo.src}`
- **`src/styles/global.css`** - Global Tailwind directives and CSS
- **`public/`** - Directly served static files (favicon.svg, etc.)

## Styling System

**Tailwind CSS 4** configured via Vite plugin in `astro.config.mjs`:

```javascript
import tailwindcss from "@tailwindcss/vite";
export default defineConfig({
  vite: { plugins: [tailwindcss()] },
});
```

**SCSS/Sass** installed (`sass@^1.97.1`). Use in components:

```astro
<style lang="scss">
  $color: #333;
  h1 { color: $color; }
</style>
```

**Global styles** in `src/styles/global.css` imported in `Layout.astro`.

## Key Patterns

### Component Frontmatter (Server-Side Only)

```astro
---
import OtherComponent from './OtherComponent.astro';
const greeting = 'Hello';
---
<h1>{greeting}</h1>
<OtherComponent />
<style lang="scss">
  h1 { color: blue; }
</style>
```

### Asset Usage

Always destructure `.src` from imported assets:

```astro
import banner from '../assets/hero.svg';
<img src={banner.src} alt="Hero" />  <!-- ✓ Correct -->
<!-- NOT: <img src={banner} alt="Hero" /> -->
```

## Development & Build

| Command                   | Purpose                                    |
| ------------------------- | ------------------------------------------ |
| `npm run dev`             | Dev server at `localhost:4321`, hot reload |
| `npm run build`           | Build to `./dist/`                         |
| `npm run preview`         | Preview production build locally           |
| `npm run astro add <pkg>` | Add Astro integrations (React, etc.)       |

**Config**: `astro.config.mjs` minimal; Astro auto-compiles TypeScript and `.astro` files.

## Type Safety

- Extends `astro/tsconfigs/strict`
- Auto-generated types at `.astro/types.d.ts`
- Use TypeScript freely in frontmatter; prefer interfaces for props

## When Adding Features

- **New page?** Create `.astro` file in `src/pages/` (auto-routed)
- **Reusable UI?** Build as `.astro` component in `src/components/`
- **Heavy interactivity?** Consider framework integration (`npm run astro add react`)—static-first is the default
- **Styling?** Use Tailwind utility classes + scoped SCSS in components

## Common Gotchas

- Don't use bare string paths for assets—always import and use `.src`
- CSS in `.astro` components is **automatically scoped**, no naming conventions needed
- Astro generates **zero JS** unless you explicitly add client interactivity (framework components or `<script>` tags)
- SCSS variables don't cross component boundaries (scoped CSS limitation)
