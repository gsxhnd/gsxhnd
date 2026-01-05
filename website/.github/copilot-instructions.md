# Copilot Instructions for Astro Website

## Project Overview

This is a corporate portfolio website built with Astro 5.x + React + TailwindCSS + SASS. The site showcases services, portfolio works, blog/news with pagination and dynamic routing, using Astro Content Collections for CMS-like content management.

**Languages**: Chinese (primary), English metadata present. **Tech Stack**: Astro, React, TypeScript, TailwindCSS 4, SASS, MDX.

## Architecture & File Structure

- **`src/pages/`** - File-based routing (`.astro` files become routes)
  - **`news/[...page].astro`** - Pagination for news/blog listing (paginated by 2 items)
  - **`news/[...slug].astro`** - Dynamic individual news/blog post pages with MDX rendering
  - **`works/[name].astro`** - Static portfolio projects (no dynamic routing needed)
- **`src/components/`** - Reusable `.astro` components (no client JavaScript)
- **`src/layouts/Layout.astro`** - Main layout with Header/Footer/slot pattern
- **`src/content/news/`** - Markdown/MDX blog posts with frontmatter metadata (title, date, draft flag)
- **`src/styles/global.css`** - TailwindCSS + global SCSS utilities; component styles use scoped `<style lang="scss">`
- **`src/content.config.ts`** - Astro Content Collections schema definition (enforces title, date, type, description, draft)

## Key Patterns & Conventions

### Content Collections (Astro Content Layer)

Use `getCollection()` with optional filtering for draft posts:

```astro
// In src/pages/news/[...slug].astro
const allNews = await getCollection("news", ({ data }) => {
  return import.meta.env.PROD ? data.draft !== true : true; // Hide drafts in production
});
```

Schema defined in `src/content.config.ts`; all content must match Zod schema (title, date required).

### Dynamic Routing

- **Pagination**: `[...page].astro` uses `paginate()` from `getStaticPaths()` → `Astro.props.page` contains `currentPage`, `data`, `lastPage`
- **Dynamic slugs**: `[...slug].astro` uses `getStaticPaths()` → `Astro.params.slug` + `Astro.props.entry` for content entry
- Both patterns use `export async function getStaticPaths()` for static generation

### Component Structure & Scoped Styling

Astro components use scoped SCSS (`.astro` files have `<style lang="scss">` blocks):

```astro
---
// Server-side TypeScript only
---
<!-- Rendered HTML -->
<style lang="scss">
  /* Auto-scoped to component; use BEM naming: .block__element--modifier */
  .contact {
    &__container { /* nesting */ }
  }
</style>
```

### Styling Architecture

- **`src/styles/global.css`** - TailwindCSS imports + global resets
- **Component styles** - Scoped SCSS in `<style lang="scss">` within `.astro` files
- **Class naming** - BEM convention (e.g., `.service-card__number`) consistent across codebase
- **Responsive patterns** - Use `@media (min-width: 768px)` in SCSS for breakpoints

## Development Workflow

### Commands

- `npm run dev` - Start local dev server at `localhost:4321` (hot reload enabled)
- `npm run build` - Build to `./dist/` for production
- `npm run preview` - Locally preview the production build
- `npm run astro ...` - Run Astro CLI commands (e.g., `npm run astro add react`)

### Build & Deployment

- Static generation: All pages built at build time (no SSR)
- Production env: `import.meta.env.PROD` is `true` during build; used to filter draft posts
- Sitemap: Auto-generated via `@astrojs/sitemap` integration

### Integrations Currently Active

- **`@astrojs/mdx`** - MDX support for `.mdx` content
- **`@astrojs/react`** - React 19 available but not used in pages (zero interactive JS by default)
- **`@astrojs/sitemap`** - Auto-sitemap generation
- **TailwindCSS 4** - Via `@tailwindcss/vite` plugin

## Adding Features

### New Blog Posts

1. Create `.md` or `.mdx` file in `src/content/news/` with filename pattern `YYYY-MM-DD-slug.{md,mdx}`
2. Add frontmatter: `title`, `date` (YAML date format), optional `type`, `description`, `draft`
3. Auto-routed via dynamic `[...slug].astro` pattern

### New Portfolio Projects

Create `.astro` file in `src/pages/works/` and link from `src/pages/works.astro`

### Adding Client Interactivity

Use `client:` directives sparingly:

```astro
import ReactComponent from '../components/Reactive.jsx';
<ReactComponent client:load />  <!-- only if needed -->
```

## Common Pitfalls

- **Draft posts in production**: Always check `import.meta.env.PROD` when filtering content
- **Asset paths**: Import images with named imports; don't use string paths
- **Scoped CSS isolation**: Component styles are auto-scoped; no class name collision risk
- **No JSX in `.astro`**: Use HTML syntax in templates; React components are for client interactivity only
- **Content schema violations**: All news entries must match schema (add missing required fields to existing posts if errors occur)

## Documentation

Refer to [Astro docs](https://docs.astro.build) for detailed guidance. The project structure matches the "Basics" template exactly.
