# Astro Website - Copilot Instructions

Quick reference for developing the stargazers.club Astro website. Details in [AGENTS.md](../AGENTS.md).

## 🚨 Critical: Draft Filtering

**Always check `import.meta.env.PROD` when querying Content Collections.** Drafts must not reach production.

```typescript
// ✅ CORRECT: Drafts filtered in production
const allNews = await getCollection("news", ({ data }) => {
  return import.meta.env.PROD ? data.draft !== true : true;
});

// ❌ AVOID: Drafts will publish
const allNews = await getCollection("news");
```

This pattern is **required** in:

- `src/pages/news/[...page].astro` (pagination)
- `src/pages/news/[...slug].astro` (individual posts)
- Any other dynamic pages listing content

---

## ⚡ Quick Start

| Task                 | Command                                      | Location             |
| -------------------- | -------------------------------------------- | -------------------- |
| Start dev server     | `npm run dev` → localhost:4321               | -                    |
| Build for production | `npm run build` → `./dist/`                  | -                    |
| Preview build        | `npm run preview`                            | -                    |
| Add blog post        | Create `YYYY-MM-DD-slug.md` with frontmatter | `src/content/news/`  |
| Add static page      | Create `page-name.astro` → auto-routed       | `src/pages/`         |
| Update navigation    | Edit `SITE_CONFIG`                           | `src/config/site.ts` |
| Style component      | Scoped `<style lang="scss">` with BEM        | `*.astro` files      |

---

## 🌍 Project Overview

**Type**: Static site generator (Astro 5.18+)  
**Language**: Chinese (default) + English routes (`/` vs `/en/`)  
**Build**: TypeScript strict mode, SCSS, MDX support  
**Deployment**: Static (no backend; all content in `/src/content/`)

### Tech Stack

- **astro** 5.18.0 — SSG framework
- **@astrojs/mdx** — MDX in content collections
- **sass** — CSS preprocessing
- **TailwindCSS 4** — Utility framework (global)

---

## 🗂️ Architecture

### Content Collections

News articles live in `src/content/news/` as Markdown/MDX with typed schema:

- **Required**: `title` (string), `date` (Date, not string)
- **Optional**: `draft` (boolean), `type` (string), `description` (string)

**Schema validation** in [src/content.config.ts](../src/content.config.ts) enforces types and rejects missing fields.

### Routing Patterns

**Pagination** (`news/[...page].astro`):

```typescript
export async function getStaticPaths({ paginate }) {
  const allNews = await getCollection("news", ({ data }) => {
    return import.meta.env.PROD ? data.draft !== true : true;
  });

  const sorted = allNews.sort((a, b) => b.data.date - a.data.date);
  return paginate(sorted, { pageSize: 2 }); // Fixed 2 per page
}

const { page } = Astro.props; // {data, currentPage, lastPage}
```

**Dynamic routes** (`news/[...slug].astro`):

```typescript
export async function getStaticPaths() {
  const allNews = await getCollection("news", ({ data }) => {
    return import.meta.env.PROD ? data.draft !== true : true;
  });

  return allNews.map((entry) => ({
    params: { slug: entry.slug },
    props: { entry },
  }));
}

const { entry } = Astro.props;
const { Content } = await render(entry);
```

Use `Astro.props.slug` only in dynamic routes. For pagination, use `Astro.props.page`.

### Component Pattern

All Astro components should have scoped SCSS:

```astro
---
interface Props {
  title?: string;
}
const { title = "默认标题" } = Astro.props;
---

<div class="my-component">
  <h1>{title}</h1>
</div>

<style lang="scss">
  .my-component {
    padding: 2rem;

    h1 {
      font-size: 1.5rem;
    }
  }
</style>
```

CSS is **auto-scoped** per component; class name collisions are impossible.

---

## 🌐 i18n Architecture

**Default language**: Chinese (`zh`)

- **Chinese routes** (no prefix): `/`, `/about`, `/news`
- **English routes** (prefixed): `/en/`, `/en/about`, `/en/news`

**Language detection**:

```typescript
import { getLangFromUrl } from "../i18n/utils.ts";
const lang = getLangFromUrl(Astro.url); // Returns "zh" or "en"

import { useTranslations } from "../i18n/utils.ts";
const t = useTranslations(lang);
const header = t("header"); // Looks up i18n/locales/[lang].json
```

Strings live in [src/i18n/locales/](../src/i18n/locales/) (zh.json, en.json).

---

## 📁 Project Structure

```
src/
├── pages/              # File-based routing
│   ├── index.astro     # Homepage
│   ├── news/[...page].astro    # Paginated list
│   ├── news/[...slug].astro    # Article detail
│   ├── en/index.astro  # English homepage
│   └── ...
├── components/         # Reusable Astro components
├── layouts/Layout.astro        # Main wrapper (Header + Footer)
├── content/news/       # Markdown/MDX blog posts
├── config/site.ts      # SITE_CONFIG (navigation, pagination size, etc.)
├── styles/global.css   # Global styles (imports Tailwind)
└── content.config.ts   # Content Collections schema
```

---

## ✅ Best Practices

1. **Always use strict TypeScript**: Define interfaces for all props
2. **Sort by date descending**: News in reverse chronological order
3. **Draft in dev, not prod**: Use `import.meta.env.PROD` check
4. **Scoped styles**: Never use global CSS in components (use `<style>`)
5. **BEM naming**: `.component { }`, `.component__element { }`, `.component--modifier { }`
6. **Test builds**: Run `npm run build` before committing to catch schema or type errors

---

## ❌ Common Mistakes

| Mistake                                 | Impact                                 | Fix                                                |
| --------------------------------------- | -------------------------------------- | -------------------------------------------------- |
| Forgot draft filter                     | Draft articles published to production | Check `import.meta.env.PROD` in `getStaticPaths`   |
| `date` as string, not Date              | Schema validation fails at build-time  | Wrap `date: new Date("2024-05-15")` in frontmatter |
| Missing `title` field                   | Content Collection build error         | Always include `title` in `.md` frontmatter        |
| Global styles in component              | Styles leak to other components        | Use `<style lang="scss">` (scoped by default)      |
| Sorting by filename not date            | Wrong order on lists                   | Sort by `data.date.valueOf()` descending           |
| Using `Astro.params.slug` in pagination | Undefined behavior                     | Use `Astro.props.page` for paginated routes        |

---

## 🔗 References

- [Astro Docs](https://docs.astro.build)
- [Content Collections](https://docs.astro.build/en/guides/content-collections/)
- [Dynamic Routes](https://docs.astro.build/en/guides/routing/#dynamic-routes)
- Project repo: <https://stargazers.club>
