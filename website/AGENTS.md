# AI 编程助手指南

本文档为 AI 编程助手（Cursor、Copilot、Cline 等）提供项目开发指导。

## 项目概述

这是一个使用 **Astro 5** 构建的企业展示网站，展示服务、作品集和新闻动态。

**技术栈**：
- Astro 5.18.0（静态站点生成器）
- TypeScript（严格模式）
- React 19（用于交互组件）
- SASS（样式预处理器）
- MDX（内容管理）

**站点 URL**：https://stargazers.club
**主要语言**：中文

## 项目结构

```
src/
├── pages/              # 基于文件的路由
│   ├── index.astro     # 首页
│   ├── news/           # 新闻动态（分页 + 动态路由）
│   │   ├── [...page].astro   # 分页列表（每页 2 条）
│   │   └── [...slug].astro   # 文章详情页
│   └── works/          # 作品集页面
├── components/         # 可复用组件
├── layouts/            # 布局模板
│   └── Layout.astro    # 主布局（Header/Footer）
├── content/            # 内容集合
│   └── news/           # MDX/Markdown 文章
├── styles/
│   └── global.css      # 全局样式
└── content.config.ts   # Content Collections 配置
```

## 开发命令

```bash
npm run dev      # 启动开发服务器（localhost:4321）
npm run build    # 生产构建（输出到 ./dist/）
npm run preview  # 预览生产构建
```

**注意**：项目未配置测试框架或 linting 工具。

## TypeScript 配置

- 继承 `astro/tsconfigs/strict`（严格模式）
- JSX 模式：`react-jsx`（自动导入 React）
- 启用 `@astrojs/ts-plugin` 用于 Astro 模板类型检查
- 所有组件和页面应使用 TypeScript

## 代码风格指南

### 导入规范

```typescript
// 1. Astro 核心导入
import { defineCollection, z } from "astro:content";
import { getCollection, render } from "astro:content";

// 2. 布局和组件导入（相对路径）
import Layout from "../layouts/Layout.astro";
import Header from "../components/Header.astro";

// 3. 样式导入
import "../styles/global.css";
```

### 命名约定

- **文件名**：
  - Astro 组件：PascalCase（`Header.astro`, `NewsCard.astro`）
  - 页面：kebab-case（`index.astro`, `services.astro`）
  - 内容文件：`YYYY-MM-DD-slug.{md,mdx}`（如 `2024-01-15-news-title.md`）

- **CSS 类名**：使用 BEM 命名规范
  ```scss
  .page-header { }
  .page-header__container { }
  .page-header__title { }
  .page-header__title--active { }
  ```

- **变量**：camelCase
- **常量**：UPPER_SNAKE_CASE
- **接口/类型**：PascalCase

### 组件结构

Astro 组件标准结构：

```astro
---
// 1. 导入
import Layout from "../layouts/Layout.astro";

// 2. 接口定义
interface Props {
  title?: string;
}

// 3. Props 解构
const { title = "默认标题" } = Astro.props;

// 4. 数据获取逻辑
const data = await fetchData();
---

<!-- 5. HTML 模板 -->
<Layout title={title}>
  <div class="component">
    <h1>{title}</h1>
  </div>
</Layout>

<!-- 6. 样式（scoped SCSS） -->
<style lang="scss">
  .component {
    padding: 2rem;
  }
</style>
```

### 样式规范

- **优先使用 scoped styles**：所有组件样式默认作用域隔离
- **全局样式**：仅在 `global.css` 或使用 `:global` 包裹
- **CSS 变量**：使用语义化命名
  ```css
  var(--color-text-primary)
  var(--color-bg-primary)
  ```
- **响应式设计**：移动优先，使用媒体查询

### Content Collections 模式

**必须遵循的草稿过滤模式**：

```typescript
// ✅ 正确：生产环境过滤草稿
const allNews = await getCollection("news", ({ data }) => {
  return import.meta.env.PROD ? data.draft !== true : true;
});

// ❌ 错误：草稿会发布到生产环境
const allNews = await getCollection("news");
```

**Schema 验证**（`src/content.config.ts`）：
- `title`: string（必需）
- `date`: Date（必需）
- `draft`: boolean（可选，默认 false）
- `type`: string（可选）
- `description`: string（可选）

### 动态路由模式

**分页路由**（`news/[...page].astro`）：

```typescript
export async function getStaticPaths({ paginate }) {
  const allNews = await getCollection("news", ({ data }) => {
    return import.meta.env.PROD ? data.draft !== true : true;
  });

  // 按日期降序排序
  const sortedNews = allNews.sort((a, b) =>
    b.data.date.valueOf() - a.data.date.valueOf()
  );

  return paginate(sortedNews, { pageSize: 2 }); // 固定每页 2 条
}

const { page } = Astro.props;
// page.data - 当前页数据
// page.currentPage - 当前页码
// page.lastPage - 总页数
```

**动态内容路由**（`news/[...slug].astro`）：

```typescript
export async function getStaticPaths() {
  const allNews = await getCollection("news", ({ data }) => {
    return import.meta.env.PROD ? data.draft !== true : true;
  });

  return allNews.map((entry) => ({
    params: { slug: entry.slug },
    props: { entry }
  }));
}

const { entry } = Astro.props;
const { Content } = await render(entry); // 渲染 MDX/Markdown
```

### React 集成

使用客户端指令控制 hydration：

```astro
---
import ContactForm from '../components/ContactForm.jsx';
---

<!-- 推荐：浏览器空闲时加载 -->
<ContactForm client:idle />

<!-- 立即加载（需要即时交互） -->
<ContactForm client:load />

<!-- 滚动到视图时加载 -->
<ContactForm client:visible />
```

## 常见错误和最佳实践

### ❌ 避免的错误

1. **草稿过滤缺失**：未使用 `import.meta.env.PROD` 检查 → 草稿发布到生产环境
2. **路由参数混淆**：使用 `Astro.props.slug` 而非 `Astro.params.slug`
3. **JSX 语法混用**：在 `.astro` 文件中使用 JSX 语法（应使用 HTML）
4. **Frontmatter 验证**：缺少 `title` 或 `date` 字段 → schema 错误
5. **全局样式泄漏**：未使用 scoped styles 导致样式冲突

### ✅ 最佳实践

1. **类型安全**：为所有 Props 定义 TypeScript 接口
2. **内容排序**：始终按日期降序排序新闻/博客
3. **错误处理**：在 `getStaticPaths` 中处理空数据情况
4. **调试输出**：服务端代码使用 `console.log()`（输出在终端，非浏览器控制台）
5. **构建验证**：修改后运行 `npm run build` 检查类型错误

## 参考资源

- [Astro 官方文档](https://docs.astro.build)
- [Content Collections](https://docs.astro.build/en/guides/content-collections/)
- [动态路由](https://docs.astro.build/en/guides/routing/#dynamic-routes)
- [Astro 组件](https://docs.astro.build/en/basics/astro-components/)
