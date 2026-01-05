# 🎨 Light/Dark 主题系统完整指南

**最后更新**: 2026-01-05

## 📖 目录

1. [快速开始](#快速开始)
2. [系统架构](#系统架构)
3. [主题配置](#主题配置)
4. [使用方法](#使用方法)
5. [实现细节](#实现细节)
6. [常见问题](#常见问题)

---

## 🚀 快速开始

### 5 分钟快速上手

#### 1. 查看主题系统在运行

网站自动包含主题切换功能。查找 Header 中的太阳/月亮图标。

#### 2. 理解主题属性

```css
/* 浅色主题 - 默认 */
:root {
  --color-bg-primary: #ffffff;
  --color-text-primary: #333333;
}

/* 深色主题 - 使用 data-theme 属性 */
:root[data-theme="dark"] {
  --color-bg-primary: #1a1a1a;
  --color-text-primary: #e8e8e8;
}
```

#### 3. 在你的组件中使用

```scss
.my-component {
  background: var(--color-bg-primary);    /* 自动切换 */
  color: var(--color-text-primary);       /* 自动切换 */
  border: 1px solid var(--color-border);  /* 自动切换 */
}
```

#### 4. 完成

组件会在用户切换主题时自动更新。无需额外代码。

---

## 🏗️ 系统架构

### 整体流程

```
用户操作 (点击主题按钮)
    ↓
ThemeToggle.tsx (React 组件)
    ↓
setAttribute() 改变 data-theme 属性
    ↓
localStorage 保存用户选择
    ↓
:root[data-theme="dark"] CSS 规则激活
    ↓
所有使用 CSS 变量的组件立即更新
```

### 文件结构

```
website/
├── public/
│   └── theme.js                    # 初始化脚本，防止闪烁
├── src/
│   ├── styles/
│   │   └── global.css              # 主题变量定义
│   ├── components/
│   │   ├── Header.astro            # 集成 ThemeToggle
│   │   ├── ThemeToggle.tsx          # 主题切换按钮
│   │   ├── Footer.astro            # 使用主题变量
│   │   └── Box.tsx                 # 使用主题变量
│   └── layouts/
│       └── Layout.astro            # 加载 theme.js
```

---

## 🎨 主题配置

### 13 个主题变量

#### 背景色

```css
--color-bg-primary:   主背景颜色（页面背景）
--color-bg-secondary: 次级背景（卡片、侧栏）
--color-bg-tertiary:  第三级背景（悬停状态、强调）
```

#### 文本色

```css
--color-text-primary:    主文本（标题、正文）
--color-text-secondary:  次级文本（描述、说明）
--color-text-tertiary:   三级文本（提示、禁用）
```

#### 边框色

```css
--color-border:       标准边框
--color-border-light: 轻量边框
```

#### 交互色

```css
--color-accent:       链接、按钮的主色调
--color-accent-hover: 悬停状态的颜色
```

#### 状态色

```css
--color-success:  成功/正确
--color-warning:  警告
--color-error:    错误
```

### 完整配置表

| 类别 | 变量名 | 浅色值 | 深色值 | 用途 |
|------|--------|--------|--------|------|
| **背景** | `--color-bg-primary` | #ffffff | #1a1a1a | 主背景 |
| | `--color-bg-secondary` | #f8f9fa | #2d2d2d | 卡片/侧栏 |
| | `--color-bg-tertiary` | #f0f2f5 | #3d3d3d | 三级背景 |
| **文本** | `--color-text-primary` | #333333 | #e8e8e8 | 标题/主要 |
| | `--color-text-secondary` | #666666 | #b0b0b0 | 描述/辅助 |
| | `--color-text-tertiary` | #999999 | #808080 | 提示/禁用 |
| **边框** | `--color-border` | #e0e0e0 | #404040 | 标准边框 |
| | `--color-border-light` | #eeeeee | #2d2d2d | 浅边框 |
| **交互** | `--color-accent` | #0066cc | #4d9fff | 链接/按钮 |
| | `--color-accent-hover` | #0052a3 | #6ab3ff | 悬停态 |
| **状态** | `--color-success` | #28a745 | #4ade80 | 成功 |
| | `--color-warning` | #ffc107 | #facc15 | 警告 |
| | `--color-error` | #dc3545 | #f87171 | 错误 |

### 主题对比

```
┌─────────────────────┬──────────────┬──────────────┐
│ 元素                 │ 浅色         │ 深色         │
├─────────────────────┼──────────────┼──────────────┤
│ 页面背景             │ #ffffff      │ #1a1a1a      │
│ 卡片背景             │ #f8f9fa      │ #2d2d2d      │
│ 主文本               │ #333333      │ #e8e8e8      │
│ 链接/按钮            │ #0066cc      │ #4d9fff      │
│ 边框                 │ #e0e0e0      │ #404040      │
│ 成功                 │ #28a745      │ #4ade80      │
│ 警告                 │ #ffc107      │ #facc15      │
│ 错误                 │ #dc3545      │ #f87171      │
└─────────────────────┴──────────────┴──────────────┘
```

---

## 💻 使用方法

### 方法 1：在 CSS/SCSS 中使用

```scss
/* 单个属性 */
.card {
  background: var(--color-bg-secondary);
  color: var(--color-text-primary);
}

/* 完整示例 */
.component {
  background-color: var(--color-bg-primary);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border);
  
  a {
    color: var(--color-accent);
    
    &:hover {
      color: var(--color-accent-hover);
    }
  }
  
  .success { color: var(--color-success); }
  .warning { color: var(--color-warning); }
  .error { color: var(--color-error); }
  
  /* 平滑过渡 */
  transition: background-color 0.3s ease, color 0.3s ease;
}
```

### 方法 2：在 JavaScript 中控制

```javascript
// 获取当前主题
const currentTheme = document.documentElement.getAttribute('data-theme');
console.log(currentTheme);  // 'light' 或 'dark'

// 设置主题为深色
document.documentElement.setAttribute('data-theme', 'dark');

// 设置主题为浅色
document.documentElement.setAttribute('data-theme', 'light');

// 使用 dataset 语法（更简洁）
document.documentElement.dataset.theme = 'dark';

// 检查是否为深色主题
if (document.documentElement.dataset.theme === 'dark') {
  console.log('深色主题已激活');
}

// 获取 CSS 变量值
const bgColor = getComputedStyle(document.documentElement)
  .getPropertyValue('--color-bg-primary')
  .trim();  // "#1a1a1a"
```

### 方法 3：在 React 中使用

```jsx
import { useState, useEffect } from 'react';

function MyComponent() {
  const [isDark, setIsDark] = useState(false);
  
  // 读取当前主题
  useEffect(() => {
    const currentTheme = document.documentElement.dataset.theme;
    setIsDark(currentTheme === 'dark');
  }, []);
  
  // 切换主题
  const toggleTheme = () => {
    const newTheme = isDark ? 'light' : 'dark';
    document.documentElement.dataset.theme = newTheme;
    localStorage.setItem('theme', newTheme);
    setIsDark(!isDark);
  };
  
  return (
    <button onClick={toggleTheme}>
      {isDark ? '☀️ 浅色' : '🌙 深色'}
    </button>
  );
}
```

### 方法 4：Astro 组件中使用

```astro
---
// src/components/MyComponent.astro
---

<div class="my-component">
  <h1>Hello Theme System</h1>
  <p>这是一个主题感知的组件</p>
</div>

<style lang="scss">
  .my-component {
    background: var(--color-bg-primary);
    color: var(--color-text-primary);
    border: 1px solid var(--color-border);
    padding: 1.5rem;
    border-radius: 8px;
    
    /* 主题切换时平滑过渡 */
    transition: background-color 0.3s ease, color 0.3s ease;
    
    h1 {
      color: var(--color-accent);
      margin-bottom: 0.5rem;
    }
  }
</style>
```

---

## 🔧 实现细节

### 1. 主题初始化脚本 (`public/theme.js`)

```javascript
(function () {
  // 1. 检查 localStorage 中保存的主题选择
  const savedTheme = localStorage.getItem('theme');
  
  // 2. 如果没有保存，检查系统偏好
  const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
  
  // 3. 确定最终使用的主题
  const theme = savedTheme || (prefersDark ? 'dark' : 'light');
  
  // 4. 立即应用（防止闪烁）
  document.documentElement.setAttribute('data-theme', theme);
})();
```

**为什么需要这个脚本？**

- 防止页面加载时的主题闪烁 (FOUC - Flash of Unstyled Content)
- 在其他 JavaScript 之前运行（使用 `is:inline` 指令）
- 检查系统偏好作为后备方案

### 2. CSS 主题定义 (`src/styles/global.css`)

```css
/* 浅色主题 - 默认 */
:root {
  --color-bg-primary: #ffffff;
  --color-text-primary: #333333;
  --color-accent: #0066cc;
  /* ... 其他 10 个变量 ... */
}

/* 深色主题 - 属性选择器 */
:root[data-theme="dark"] {
  --color-bg-primary: #1a1a1a;
  --color-text-primary: #e8e8e8;
  --color-accent: #4d9fff;
  /* ... 其他 10 个变量 ... */
}
```

**为什么使用 `data-theme` 属性？**

- 更语义化，清晰表达目的
- 易于扩展（可以添加其他主题值）
- 与 JavaScript 集成简单
- 性能优于 class 切换

### 3. React 主题切换器 (`src/components/ThemeToggle.tsx`)

```jsx
import { useState, useEffect } from 'react';

export default function ThemeToggle() {
  const [isMounted, setIsMounted] = useState(false);
  const [isDark, setIsDark] = useState(false);
  
  useEffect(() => {
    // 防止水合错误
    setIsMounted(true);
    const theme = document.documentElement.dataset.theme;
    setIsDark(theme === 'dark');
  }, []);
  
  const toggleTheme = () => {
    const root = document.documentElement;
    const newTheme = root.dataset.theme === 'dark' ? 'light' : 'dark';
    root.dataset.theme = newTheme;
    localStorage.setItem('theme', newTheme);
    setIsDark(newTheme === 'dark');
  };
  
  if (!isMounted) return null;
  
  return (
    <button 
      onClick={toggleTheme}
      className="theme-toggle"
      aria-label="切换主题"
    >
      {isDark ? '☀️' : '🌙'}
    </button>
  );
}
```

**关键特点**:

- `isMounted` 检查防止水合错误
- 从 localStorage 恢复用户选择
- 同时更新 UI 和 localStorage
- 可访问性属性 (`aria-label`)

### 4. Layout 集成 (`src/layouts/Layout.astro`)

```astro
---
import Header from '../components/Header.astro';
import Footer from '../components/Footer.astro';
import '../styles/global.css';
---

<!doctype html>
<html lang="zh-CN">
  <head>
    {/* 防止 FOUC 的关键脚本 */}
    <script is:inline src="/theme.js"></script>
  </head>
  <body>
    <Header />
    <main>
      <slot />
    </main>
    <Footer />
  </body>
</html>

<style lang="scss">
  body {
    background-color: var(--color-bg-primary);
    color: var(--color-text-primary);
    transition: background-color 0.3s ease, color 0.3s ease;
  }
</style>
```

---

## ❓ 常见问题

### Q1: 如何添加新的主题变量？

**A**: 在 `src/styles/global.css` 中添加：

```css
:root {
  /* ... 现有变量 ... */
  --color-custom: #your-color;
}

:root[data-theme="dark"] {
  /* ... 现有变量 ... */
  --color-custom: #your-dark-color;
}
```

然后在任何地方使用：

```scss
.component {
  color: var(--color-custom);
}
```

### Q2: 为什么我的主题没有切换？

**A**: 检查以下几点：

1. 确保使用了 `var(--color-*)` CSS 变量
2. 检查浏览器控制台是否有错误
3. 清除浏览器缓存 (Ctrl+Shift+Delete)
4. 确保 `/public/theme.js` 文件存在
5. 检查 `Layout.astro` 中是否包含 `<script is:inline src="/theme.js"></script>`

### Q3: 如何在特定条件下强制使用某个主题？

**A**: 在 JavaScript 中：

```javascript
// 强制深色主题
document.documentElement.dataset.theme = 'dark';
localStorage.setItem('theme', 'dark');

// 强制浅色主题
document.documentElement.dataset.theme = 'light';
localStorage.setItem('theme', 'light');
```

### Q4: 如何检测系统主题偏好变更？

**A**: 使用 MediaQueryList API：

```javascript
const darkMode = window.matchMedia('(prefers-color-scheme: dark)');

darkMode.addEventListener('change', (e) => {
  if (e.matches) {
    console.log('系统切换到深色模式');
  } else {
    console.log('系统切换到浅色模式');
  }
});
```

### Q5: 如何添加第三个主题（如自动）？

**A**: 扩展 `theme.js` 和 CSS：

```javascript
// 在 public/theme.js 中
const theme = savedTheme || (prefersDark ? 'dark' : 'auto');
```

```css
/* 在 src/styles/global.css 中 */
@media (prefers-color-scheme: dark) {
  :root[data-theme="auto"] {
    --color-bg-primary: #1a1a1a;
    /* ... */
  }
}
```

### Q6: 主题变量性能如何？

**A**: 非常好！CSS 变量的性能：

- ✅ 零运行时开销
- ✅ 硬件加速支持
- ✅ 不会触发重排 (reflow)
- ✅ 比 JavaScript 快 100 倍

### Q7: 如何在 MDX 内容中使用主题变量？

**A**: 在 MDX 前置元数据中定义，或在组件中使用：

```mdx
---
title: "My Post"
---

<div style={{ background: 'var(--color-bg-secondary)' }}>
  This respects the theme!
</div>
```

---

## 📚 最佳实践

### ✅ 推荐

```scss
/* 1. 对所有主题敏感的属性使用变量 */
.component {
  background: var(--color-bg-primary);
  color: var(--color-text-primary);
  border-color: var(--color-border);
}

/* 2. 添加过渡以实现平滑切换 */
.component {
  transition: background-color 0.3s ease;
}

/* 3. 为交互元素使用 accent 颜色 */
a, button {
  color: var(--color-accent);
}

/* 4. 对禁用状态使用 tertiary 颜色 */
:disabled {
  color: var(--color-text-tertiary);
}
```

### ❌ 避免

```scss
/* 1. ❌ 不要硬编码颜色 */
.component {
  background: #ffffff;  /* 不行 */
  color: #333333;       /* 不行 */
}

/* 2. ❌ 不要在深色主题中覆盖所有变量 */
:root[data-theme="dark"] {
  --color-bg-primary: #1a1a1a;
  --color-bg-secondary: #2d2d2d;
  /* ... 应该是 13 个变量都在这里 ... */
}

/* 3. ❌ 不要依赖 .dark class */
.dark .component { }  /* 已过时 */
```

---

## 🔗 快速参考

### 属性设置

```javascript
document.documentElement.dataset.theme = 'dark' | 'light'
```

### 属性获取

```javascript
document.documentElement.dataset.theme  // 返回当前主题
```

### 存储

```javascript
localStorage.getItem('theme')      // 获取保存的主题
localStorage.setItem('theme', 'dark')  // 保存主题
```

### CSS 用法

```css
background: var(--color-bg-primary);
color: var(--color-text-primary);
border: var(--color-border);
```

---

## 📞 支持

如有问题：

1. 检查 `global.css` 中是否定义了变量
2. 验证 HTML 是否有 `data-theme` 属性
3. 清除浏览器缓存重试
4. 检查浏览器控制台错误

---

**更新日期**: 2026-01-05  
**版本**: 1.0 - data-theme 属性版本  
**状态**: ✅ 生产就绪  
