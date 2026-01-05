# 📚 网站文档中心

**最后更新**: 2026-01-05

## 🎯 文档导航

本文档库包含了关于网站的所有重要信息，分为两个主要部分：响应式设计和主题系统。

### 🚀 首次访问？从这里开始

1. **[RESPONSIVE_DESIGN.md](./RESPONSIVE_DESIGN.md)** - 响应式布局完整指南
   - 适合想了解网站如何在不同设备上表现
   - 包含测试方法和最佳实践

2. **[THEME_SYSTEM.md](./THEME_SYSTEM.md)** - Light/Dark 主题系统
   - 适合想使用或扩展主题功能
   - 包含快速上手和详细实现

---

## 📄 完整文档列表

### 第一部分：响应式设计

#### [响应式设计完整指南](./RESPONSIVE_DESIGN.md)

**内容**:

- ✅ 项目完成情况 (所有页面都已优化)
- 📱 5 个主要响应式断点
- 🎨 组件响应式实现示例
- 📄 各页面响应式表现详解
- 🧪 测试与验证方法
- 💡 最佳实践和避免做法

**适合**:

- 想了解网站响应式布局的开发者
- 需要添加新页面的团队成员
- 想学习响应式设计的人

**关键数据**:

- 页面覆盖: 所有 12 个页面 + 4 个作品详情页
- 断点: 320px, 640px, 768px, 1024px, 1280px, 1536px
- 完成度: 100% ✅

---

### 第二部分：主题系统

#### [Light/Dark 主题系统完整指南](./THEME_SYSTEM.md)

**内容**:

- 🚀 5 分钟快速开始
- 🏗️ 系统架构和流程图
- 🎨 13 个主题变量的完整配置表
- 💻 4 种使用方法 (CSS/JS/React/Astro)
- 🔧 实现细节和代码示例
- ❓ 常见问题和解答
- ✅ 最佳实践

**适合**:

- 想快速集成主题功能的开发者
- 需要扩展主题变量的团队成员
- 想深入理解主题系统的人

**关键特性**:

- 架构: `data-theme` 属性 + CSS 变量
- 变量: 13 个主题变量 (背景、文本、边框、交互、状态)
- 持久化: localStorage 自动保存用户选择
- 性能: 零运行时开销，硬件加速支持

---

## 🗂️ 关键文件位置

### 响应式设计相关

```
src/
├── styles/
│   └── global.css              # 响应式断点定义和全局样式
├── components/
│   ├── Header.astro            # 响应式导航 (移动菜单)
│   └── Footer.astro            # 响应式网格布局
├── layouts/
│   └── Layout.astro            # 基础响应式布局
└── pages/
    ├── index.astro             # 首页 - 所有部分响应式
    ├── about.astro             # 关于页面
    ├── services.astro          # 服务页面
    ├── contact.astro           # 联系页面
    ├── works.astro             # 作品列表
    ├── company.astro           # 公司信息
    ├── privacy.astro           # 隐私政策
    ├── works/                  # 4 个作品详情页
    └── news/                   # 新闻详情 (动态)
```

### 主题系统相关

```
public/
└── theme.js                    # 主题初始化脚本

src/
├── styles/
│   └── global.css              # 13 个主题变量定义
├── components/
│   ├── Header.astro            # 包含 ThemeToggle 按钮
│   ├── ThemeToggle.tsx          # 主题切换 React 组件
│   ├── Footer.astro            # 使用主题变量
│   └── Box.tsx                 # 示例组件
└── layouts/
    └── Layout.astro            # 加载 theme.js 初始化脚本
```

---

## 📊 项目统计

### 响应式覆盖

| 项目 | 状态 | 覆盖 |
|------|------|------|
| 静态页面 | ✅ 完成 | 7 个页面 |
| 动态页面 (作品) | ✅ 完成 | 4 个页面 |
| 动态页面 (新闻) | ✅ 完成 | 可变数量 |
| 组件 | ✅ 完成 | Header, Footer, Layout |
| **总计** | ✅ **100%** | **12+ 页面** |

### 主题变量覆盖

| 类别 | 数量 | 状态 |
|------|------|------|
| 背景色 | 3 个 | ✅ 完成 |
| 文本色 | 3 个 | ✅ 完成 |
| 边框色 | 2 个 | ✅ 完成 |
| 交互色 | 2 个 | ✅ 完成 |
| 状态色 | 3 个 | ✅ 完成 |
| **总计** | **13 个** | ✅ **完成** |

---

## 🎯 常见任务

### 任务：添加新页面的响应式设计

**参考文档**: [RESPONSIVE_DESIGN.md](./RESPONSIVE_DESIGN.md#最佳实践)

**步骤**:

1. 查看"最佳实践"部分的代码示例
2. 使用移动优先的方法
3. 在 640px、768px、1024px 处测试
4. 使用 CSS 变量确保主题兼容性

**参考文件**: `src/pages/about.astro` (完整示例)

### 任务：集成主题到新组件

**参考文档**: [THEME_SYSTEM.md](./THEME_SYSTEM.md#方法1在-cssscss-中使用)

**步骤**:

1. 查看"使用方法"部分的代码示例
2. 替换硬编码颜色为 `var(--color-*)`
3. 添加 `transition` 以实现平滑切换
4. 测试浅色和深色主题

**参考文件**: `src/components/Header.astro` (完整示例)

### 任务：添加新的主题变量

**参考文档**: [THEME_SYSTEM.md](./THEME_SYSTEM.md#常见问题)

**步骤**:

1. 打开 `src/styles/global.css`
2. 在 `:root {}` 和 `:root[data-theme="dark"] {}` 中各添加一行
3. 在组件中使用 `var(--color-new-name)`

**示例**:

```css
:root {
  --color-custom: #your-light-color;
}

:root[data-theme="dark"] {
  --color-custom: #your-dark-color;
}
```

---

## 💡 快速参考

### 响应式断点

```css
/* 在 SCSS 中使用 */
@media (min-width: 640px) { /* 平板 */ }
@media (min-width: 768px) { /* 平板+ */ }
@media (min-width: 1024px) { /* 笔记本 */ }
@media (min-width: 1280px) { /* 桌面 */ }
```

### 主题变量使用

```scss
/* 在任何 SCSS 中 */
.component {
  background: var(--color-bg-primary);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border);
}
```

### 主题切换 (JavaScript)

```javascript
// 设置主题
document.documentElement.dataset.theme = 'dark';

// 获取当前主题
const theme = document.documentElement.dataset.theme;

// 保存用户选择
localStorage.setItem('theme', 'dark');
```

---

## 📞 需要帮助？

### 响应式相关问题

1. **页面在某个尺寸下看起来不对？**
   - 打开 [RESPONSIVE_DESIGN.md](./RESPONSIVE_DESIGN.md#测试与验证)
   - 查看"测试与验证"部分

2. **如何在新页面添加响应式设计？**
   - 参考 [RESPONSIVE_DESIGN.md](./RESPONSIVE_DESIGN.md#最佳实践)
   - 查看"最佳实践"部分的代码示例

3. **需要学习响应式设计？**
   - 阅读 [RESPONSIVE_DESIGN.md](./RESPONSIVE_DESIGN.md#组件响应式实现)
   - 查看"组件响应式实现"部分的详细代码

### 主题系统相关问题

1. **如何快速开始？**
   - 打开 [THEME_SYSTEM.md](./THEME_SYSTEM.md#快速开始)
   - 5 分钟快速开始指南

2. **主题没有切换？**
   - 参考 [THEME_SYSTEM.md](./THEME_SYSTEM.md#q2-为什么我的主题没有切换)
   - 查看常见问题部分

3. **需要添加新主题变量？**
   - 参考 [THEME_SYSTEM.md](./THEME_SYSTEM.md#q1-如何添加新的主题变量)
   - 查看常见问题部分

---

## 📈 项目统计

- **总页面数**: 12+ 个 (包括动态路由)
- **响应式完成度**: ✅ 100%
- **主题变量数**: 13 个
- **支持的主题**: 2 个 (浅色、深色)
- **响应式断点**: 6 个
- **CSS 变量使用**: ✅ 全网站

---

## 📝 文档维护

### 最后更新记录

| 日期 | 内容 | 更新人 |
|------|------|--------|
| 2026-01-05 | 初始整合所有根目录文档到 doc 文件夹 | System |
| 2026-01-05 | 创建响应式设计完整指南 | System |
| 2026-01-05 | 创建主题系统完整指南 | System |

### 文档标准

所有文档遵循以下标准：

- ✅ 包含目录
- ✅ 代码示例
- ✅ 清晰的章节划分
- ✅ 快速参考部分
- ✅ 常见问题解答

---

## 🔗 相关资源

### 官方文档

- [Astro 官方文档](https://docs.astro.build)
- [Astro 组件文档](https://docs.astro.build/zh-cn/basics/astro-components/)
- [CSS 变量 (MDN)](https://developer.mozilla.org/en-US/docs/Web/CSS/--*)
- [响应式设计 (MDN)](https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout/Responsive_Design)

### 项目相关

- **主项目位置**: `/home/gsxhnd/Code/personal/demo/website`
- **技术栈**: Astro 5.16.6 + React 19 + TailwindCSS 4
- **Primary Language**: Chinese (中文)

---

**文档版本**: 1.0  
**最后更新**: 2026-01-05  
**状态**: ✅ 完成  

---

*如需更新或添加内容，请更新此索引和相应的具体文档。*
