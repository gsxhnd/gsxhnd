# Copilot Instructions for Astro Website

## Project Overview

This is a corporate portfolio website built with **Astro 5.16.6** + **React 19** + **TailwindCSS 4** + **SASS**. The site showcases services, portfolio works, and a blog/news section with pagination and dynamic routing using **Astro Content Collections** for CMS-like content management.

**Primary Language**: Chinese. **Tech Stack**: Astro, React, TypeScript, TailwindCSS 4, SASS, MDX. **Site URL**: https://stargazers.club

## Architecture & File Structure

- **`src/pages/`** - File-based routing (`.astro` files become routes)
  - **`news/[...page].astro`** - Pagination for news/blog listing (pageSize: 2 items per page; uses `getStaticPaths()` + `paginate()`)
  - **`news/[...slug].astro`** - Dynamic individual news/blog post pages (uses `getStaticPaths()` + `render(entry)` for MDX)
  - **`works/`** - Static portfolio project pages (one per project, no dynamic routing)
  - **Home pages**: `index.astro`, `services.astro`, `about.astro`, `company.astro`, `contact.astro`, `privacy.astro`
- **`src/components/`** - Reusable `.astro` components: `Header.astro`, `Footer.astro`, `Contact.astro`, `Box.tsx` (React)
- **`src/layouts/Layout.astro`** - Main layout wrapper with Header/Footer/slot pattern
- **`src/content/news/`** - Markdown/MDX blog posts with naming pattern: `YYYY-MM-DD-slug.{md,mdx}` + frontmatter
- **`src/styles/global.css`** - Global TailwindCSS imports and resets; component-level styling uses scoped `<style lang="scss">`
- **`src/content.config.ts`** - Zod schema enforces: `title` (string), `date` (z.date()), `draft` (boolean, default false), optional `type`, `description`

## Key Patterns & Conventions

### Content Collections & Draft Filtering

**Always use this pattern when querying content:**

```astro
const allNews = await getCollection("news", ({ data }) => {
  return import.meta.env.PROD ? data.draft !== true : true;
});
```

- Dev mode: Shows all posts (including `draft: true`)
- Production build: Hides draft posts
- Schema enforced in `src/content.config.ts` (title, date required; draft optional, defaults false)
- Content files use format: `src/content/news/YYYY-MM-DD-slug.{md,mdx}`

### Dynamic Routing Patterns

**Pagination** (`news/[...page].astro`):

```astro
export async function getStaticPaths({ paginate }) {
  const allNews = await getCollection("news", ...);
  return paginate(allNews, { pageSize: 2 }); // Fixed: 2 items per page
}
const { page } = Astro.props; // page.currentPage, page.data[], page.lastPage
```

**Dynamic Content** (`news/[...slug].astro`):

```astro
export async function getStaticPaths() {
  const allNews = await getCollection("news", ...);
  return allNews.map((entry) => ({
    params: { slug: entry.slug },
    props: { entry }
  }));
}
const { entry } = Astro.props;
const { Content } = await render(entry); // Renders MDX/Markdown
```

### Component Structure & Scoped Styling

All Astro components use **scoped SCSS** with **BEM naming**:

```astro
---
import Layout from "./Layout.astro";
---
<div class="page-header">
  <div class="page-header__container"></div>
</div>

<style lang="scss">
  .page-header {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
    padding: 2rem 1.5rem;

    @media (min-width: 768px) {
      padding: 3rem 2rem;
    }

    &__container {
      max-width: 1200px;
      margin: 0 auto;
    }
  }
</style>
```

- **Scoping**: Styles automatically scoped to component; no collision risk
- **Naming**: BEM convention (`.block__element--modifier`)
- **Responsive**: Use `@media (min-width: 768px)` for tablet breakpoints
- **Global imports**: `src/styles/global.css` (TailwindCSS) imported in `Header.astro`

### Assets & Image Handling

Always use `Image` component from `astro:assets`:

```astro
import { Image } from "astro:assets";
import logoImage from "../assets/astro.svg";
<Image src={logoImage} alt="description" />
```

Never use string paths for images.

## Development Workflow

### Commands

```bash
npm run dev      # Start localhost:4321 (hot reload)
npm run build    # Build to ./dist/ for production
npm run preview  # Locally preview production build
npm run astro    # Direct Astro CLI access
```

### Build & Deployment

- **Static Generation**: All pages pre-built at build time (no SSR); routes in `getStaticPaths()` become static HTML
- **Production Filter**: `import.meta.env.PROD` (boolean) is true during build; used for filtering draft posts
- **Sitemap**: Auto-generated to `dist/sitemap.xml` via `@astrojs/sitemap` integration
- **Site Config**: Base URL set in `astro.config.mjs` as `site: "https://stargazers.club"`

### Active Integrations

- **`@astrojs/mdx`** (v4.3.13) - Enables `.mdx` files with JSX in content
- **`@astrojs/react`** (v4.4.2) - React 19 support; sparingly used (client-side interactivity via `client:` directives)
- **`@astrojs/sitemap`** (v3.6.0) - Auto-generates sitemap
- **`@astrojs/ts-plugin`** (v1.10.6) - TypeScript support in Astro templates
- **TailwindCSS 4** - Configured in `global.css`; auto-scoped component SCSS works alongside

## Common Tasks

### Add a New Blog Post

1. Create `src/content/news/YYYY-MM-DD-slug.md` (or `.mdx` for JSX)
2. Include frontmatter:
   ```yaml
   ---
   title: "Post Title"
   date: 2026-01-05
   draft: false
   type: "News"
   description: "Optional summary"
   ---
   ```
3. Auto-routed; visible at `/news/YYYY-MM-DD-slug/` and listed on paginated `/news/` page

### Add a New Static Page

1. Create `.astro` file in `src/pages/` (auto-routed by filename)
2. Use Layout wrapper: `import Layout from "../layouts/Layout.astro";`
3. Add link to `Header.astro` navigation if needed

### Add Portfolio Project

1. Create `.astro` file in `src/pages/works/project-name.astro`
2. Link from `src/pages/works.astro` listing

### Add Client-Side Interactivity

Use React sparingly with `client:` directives:

```astro
import ContactForm from '../components/ContactForm.jsx';
<ContactForm client:load />  <!-- loads JS only when component renders -->
```

- `client:load` - Eagerly hydrate (immediate interaction needed)
- `client:idle` - Hydrate when browser idle (default, preferred)
- `client:visible` - Hydrate on scroll into view

## Critical Details

### TypeScript Configuration

- **`tsconfig.json`** extends `astro/tsconfigs/strict`
- JSX mode: `react-jsx` (auto imports React)
- `@astrojs/ts-plugin` enabled for Astro template type checking

### Testing & Debugging

- Console output visible in dev terminal (not browser console for server code)
- `.astro` files are server-only; use `console.log()` in frontmatter for debugging
- Build errors show full context; check terminal during `npm run build`

### Common Mistakes

1. **Draft filtering**: Missing `import.meta.env.PROD` check → drafts published in production
2. **Image imports**: Using string paths instead of `import { Image }` → broken images
3. **JSX in `.astro` templates**: Use HTML syntax; JSX only in `.jsx` components
4. **Frontmatter validation**: Missing `title` or `date` fields → schema errors; check `src/content.config.ts`
5. **Route params**: Use `Astro.params.slug` not `Astro.props.slug` in dynamic routes
6. **Scoped styles**: Styles auto-scoped; safe to use short class names (no global collisions)

### References

- [Astro Docs](https://docs.astro.build) - Official reference
- [Content Collections](https://docs.astro.build/en/guides/content-collections/) - Schema, filtering, rendering
- [Dynamic Routes](https://docs.astro.build/en/guides/routing/#dynamic-routes) - `getStaticPaths()` patterns
- [Astro Components](https://docs.astro.build/en/basics/astro-components/) - Component syntax, scoped styles
