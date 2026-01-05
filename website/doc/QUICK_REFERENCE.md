# 🎯 开发快速参考卡

**Last Updated**: 2026-01-05

## 📋 响应式设计速查

### 断点速查表

```css
/* 在 SCSS 中 */
@media (min-width: 640px) { }   /* 小屏幕 */
@media (min-width: 768px) { }   /* 平板 */
@media (min-width: 1024px) { }  /* 笔记本 */
@media (min-width: 1280px) { }  /* 桌面 */
@media (min-width: 1536px) { }  /* 大屏幕 */
```

### 常见网格模式

```scss
/* 1 → 2 列响应式网格 */
.grid {
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;
  
  @media (min-width: 768px) {
    grid-template-columns: repeat(2, 1fr);
    gap: 1.5rem;
  }
}

/* 1 → 2 → 3 列响应式网格 */
.grid {
  grid-template-columns: 1fr;
  
  @media (min-width: 640px) {
    grid-template-columns: repeat(2, 1fr);
  }
  
  @media (min-width: 1024px) {
    grid-template-columns: repeat(3, 1fr);
  }
}

/* 自适应网格 (推荐) */
.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
}
```

### 流体字体大小

```scss
/* 自动缩放字体 */
h1 { font-size: clamp(2rem, 4vw, 3.5rem); }
h2 { font-size: clamp(1.5rem, 3vw, 2.5rem); }
body { font-size: clamp(0.875rem, 1.5vw, 1.125rem); }
```

### 移动优先模式

```scss
.component {
  /* 📱 移动 - 基础样式 */
  font-size: 1rem;
  padding: 1rem;
  grid-template-columns: 1fr;
  
  /* 📱 → 💻 升级 */
  @media (min-width: 768px) {
    font-size: 1.125rem;
    padding: 1.5rem;
    grid-template-columns: repeat(2, 1fr);
  }
  
  /* 💻 → 🖥️ 升级 */
  @media (min-width: 1024px) {
    font-size: 1.25rem;
    padding: 2rem;
    grid-template-columns: repeat(3, 1fr);
  }
}
```

---

## 🎨 主题系统速查

### 主题变量快速查表

| 用途 | 变量名 | 浅色 | 深色 |
|------|--------|------|------|
| 主背景 | `--color-bg-primary` | #fff | #1a1a1a |
| 卡片 | `--color-bg-secondary` | #f8f9fa | #2d2d2d |
| 强调 | `--color-bg-tertiary` | #f0f2f5 | #3d3d3d |
| 主文本 | `--color-text-primary` | #333 | #e8e8 |
| 辅助文本 | `--color-text-secondary` | #666 | #b0b0 |
| 提示文本 | `--color-text-tertiary` | #999 | #808080 |
| 边框 | `--color-border` | #e0e0e0 | #404040 |
| 浅边框 | `--color-border-light` | #eee | #2d2d2d |
| 链接 | `--color-accent` | #0066cc | #4d9fff |
| 悬停 | `--color-accent-hover` | #0052a3 | #6ab3ff |
| 成功 | `--color-success` | #28a745 | #4ade80 |
| 警告 | `--color-warning` | #ffc107 | #facc15 |
| 错误 | `--color-error` | #dc3545 | #f87171 |

### 在 SCSS 中使用主题

```scss
.component {
  /* 背景 */
  background: var(--color-bg-primary);
  
  /* 文本 */
  color: var(--color-text-primary);
  
  /* 边框 */
  border: 1px solid var(--color-border);
  
  /* 链接 */
  a {
    color: var(--color-accent);
    
    &:hover {
      color: var(--color-accent-hover);
    }
  }
  
  /* 状态 */
  .success { color: var(--color-success); }
  .warning { color: var(--color-warning); }
  .error { color: var(--color-error); }
  
  /* 平滑过渡 */
  transition: background-color 0.3s ease, color 0.3s ease;
}
```

### 在 JavaScript 中使用主题

```javascript
/* 获取当前主题 */
const theme = document.documentElement.dataset.theme;

/* 设置主题 */
document.documentElement.dataset.theme = 'dark';    // 深色
document.documentElement.dataset.theme = 'light';   // 浅色

/* 保存选择 */
localStorage.setItem('theme', 'dark');
localStorage.getItem('theme');

/* 检查系统偏好 */
const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches;

/* 获取 CSS 变量值 */
const bgColor = getComputedStyle(document.documentElement)
  .getPropertyValue('--color-bg-primary').trim();
```

### 在 React 中使用主题

```jsx
import { useState, useEffect } from 'react';

function ThemeToggle() {
  const [theme, setTheme] = useState('light');
  
  useEffect(() => {
    const current = document.documentElement.dataset.theme;
    setTheme(current);
  }, []);
  
  const toggle = () => {
    const newTheme = theme === 'dark' ? 'light' : 'dark';
    document.documentElement.dataset.theme = newTheme;
    localStorage.setItem('theme', newTheme);
    setTheme(newTheme);
  };
  
  return <button onClick={toggle}>{theme === 'dark' ? '☀️' : '🌙'}</button>;
}
```

### 在 Astro 组件中使用主题

```astro
---
// src/components/MyComponent.astro
---

<div class="component">
  <h1>Hello</h1>
  <p>This is theme-aware</p>
</div>

<style lang="scss">
  .component {
    background: var(--color-bg-primary);
    color: var(--color-text-primary);
    transition: background-color 0.3s ease;
  }
</style>
```

---

## 📁 文件位置速查

### 响应式相关

```
src/styles/global.css          # 断点定义和全局样式
src/components/Header.astro    # 响应式导航示例
src/components/Footer.astro    # 响应式网格示例
src/pages/index.astro          # 完整响应式示例
```

### 主题相关

```
public/theme.js                     # 初始化脚本
src/styles/global.css               # 主题变量定义
src/components/ThemeToggle.tsx      # 切换按钮
src/components/Header.astro         # 包含切换按钮
src/layouts/Layout.astro            # 加载初始化脚本
```

---

## 🧪 测试速查

### 响应式测试

```bash
# 方法 1: 浏览器开发者工具
F12 或 Ctrl+Shift+M 激活响应式模式

# 方法 2: 手动调整宽度
调整浏览器窗口宽度测试各断点

# 测试清单
- [ ] 320px (超小屏)
- [ ] 640px (手机)
- [ ] 768px (平板)
- [ ] 1024px (笔记本)
- [ ] 1280px (桌面)
- [ ] 1920px (大屏)
```

### 主题测试

```bash
# 在浏览器控制台中测试
document.documentElement.dataset.theme = 'dark'
document.documentElement.dataset.theme = 'light'

# 检查是否保存
localStorage.getItem('theme')

# 检查系统偏好
window.matchMedia('(prefers-color-scheme: dark)').matches
```

---

## ⚡ 常用命令

```bash
# 开发
npm run dev              # 启动开发服务器

# 构建
npm run build            # 构建生产版本

# 预览
npm run preview          # 本地预览生产版本

# 其他
npm run astro            # 直接调用 Astro CLI
```

---

## ✅ 检查清单

### 添加新页面时

- [ ] 使用移动优先方法编写样式
- [ ] 在 640px、768px、1024px 处测试
- [ ] 使用 CSS 变量代替硬编码颜色
- [ ] 确保导航在移动设备上可用
- [ ] 检查没有水平滚动

### 添加新组件时

- [ ] 所有颜色都使用 `var(--color-*)`
- [ ] 为颜色变化添加过渡 (`transition`)
- [ ] 测试浅色和深色主题
- [ ] 确保响应式布局
- [ ] 检查可访问性

### 部署前

- [ ] 运行 `npm run build` 无错误
- [ ] 在多个浏览器测试
- [ ] 测试多个设备尺寸
- [ ] 测试主题切换功能
- [ ] 检查控制台无错误

---

## 🚨 常见错误

### ❌ 硬编码颜色

```scss
/* 不好 */
.component {
  background: #ffffff;  /* 深色主题看不见 */
  color: #333333;
}

/* 好 */
.component {
  background: var(--color-bg-primary);
  color: var(--color-text-primary);
}
```

### ❌ 忘记响应式

```scss
/* 不好 */
.grid {
  grid-template-columns: repeat(4, 1fr);  /* 移动设备上太拥挤 */
}

/* 好 */
.grid {
  grid-template-columns: 1fr;
  
  @media (min-width: 768px) {
    grid-template-columns: repeat(2, 1fr);
  }
  
  @media (min-width: 1024px) {
    grid-template-columns: repeat(3, 1fr);
  }
}
```

### ❌ 没有过渡

```scss
/* 不好 */
.component {
  background: var(--color-bg-primary);  /* 切换时闪烁 */
}

/* 好 */
.component {
  background: var(--color-bg-primary);
  transition: background-color 0.3s ease;  /* 平滑切换 */
}
```

---

## 💡 Pro Tips

### Tip 1: 使用 clamp() 实现流体排版

```scss
/* 自动从 1rem 到 1.5rem，根据视口宽度 */
body {
  font-size: clamp(1rem, 2vw, 1.5rem);
}
```

### Tip 2: 使用 auto-fit 实现自适应网格

```scss
/* 自动根据可用空间调整列数 */
.grid {
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
}
```

### Tip 3: 快速检查主题

在浏览器控制台粘贴：

```javascript
// 切换主题
document.documentElement.dataset.theme = 
  document.documentElement.dataset.theme === 'dark' ? 'light' : 'dark'
```

### Tip 4: 调试 CSS 变量

```javascript
// 获取所有 CSS 变量的当前值
const styles = getComputedStyle(document.documentElement);
const cssVars = {};
for (let i = 0; i < styles.length; i++) {
  const propName = styles[i];
  if (propName.startsWith('--color')) {
    cssVars[propName] = styles.getPropertyValue(propName).trim();
  }
}
console.table(cssVars);
```

### Tip 5: 强制系统颜色模式

```javascript
// 仅使用系统偏好
const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
const theme = isDark ? 'dark' : 'light';
document.documentElement.dataset.theme = theme;
```

---

## 📚 更多资源

详细信息请查看完整文档：

- 响应式设计: [RESPONSIVE_DESIGN.md](./RESPONSIVE_DESIGN.md)
- 主题系统: [THEME_SYSTEM.md](./THEME_SYSTEM.md)
- 文档中心: [README.md](./README.md)

---

**Version**: 1.0  
**Last Updated**: 2026-01-05  
**Status**: ✅ Ready to use
