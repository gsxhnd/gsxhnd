# Development Guidelines

This document provides guidelines for agents working on this codebase.

## Project Overview

Vue 3 + TypeScript + Vite project using Reka UI components. Single-page application demonstrating Reka UI with internationalization support via vue-i18n.

## Build, Lint, and Test Commands

**Development:**
```bash
npm run dev        # Start Vite dev server
```

**Build:**
```bash
npm run build      # TypeScript check + production build
npm run preview    #preview production build locally
```

**Test:**
No test framework configured yet. Add Vitest + Testing Library when needed:
- `npm run test` - Run all tests
- `npm run test -- --run` - Run single test execution (no watch)
- `npm run test -- src/components/MyComponent.spec.ts` - Run specific file
- `npm run test -- -- testNamePattern` - Filter by test name

## Code Style Guidelines

### Imports
- Use ES modules (`import/export`)
- Group imports: standard library → built-in packages → relative paths
- Sort alphabetically within groups
- Prefer named imports; avoid default imports unless idiomatic (e.g., Vue API)

### Formatting
- Single quotes for strings
- Semicolons required
- Tab indentation inside `<script>` blocks; spaces for CSS
- Trailing commas where helpful
- Max line length: 100 characters

### TypeScript & Types
- Strict mode enabled (`"strict": true` in tsconfig)
- Always annotate component props using `defineProps<{ ... }>()` pattern
- Use `defineEmits<{ ... }>()` for events
- Annotate function parameters and return types explicitly
- Prefer type aliases over interfaces for component props
- Avoid `any`; use `unknown` when type is truly uncertain
- Leverage Vue's TypeRef for reactive declarations

### Naming Conventions
- PascalCase for component filenames (`MyComponent.vue`)
- camelCase for functions, variables, and properties
- UPPERCASE for constants
- Composable/functions: `useXxx()` prefix pattern
- Props: kebab-case in templates, camelCase in code

### Error Handling
- Use try/catch for async operations and external calls
- Define custom error types for domain-specific errors
- Graceful degradation expected; show appropriate user feedback
- Log errors with context using `console.error()`

### Component Pattern (Script Setup)
```vue
<script setup lang="ts">
import { ref, computed } from 'vue'

interface Props {
  msg: string
  optional?: number
}

const props = defineProps<Props>()

interface Emits {
  (e: 'update', value: boolean): void
}

const emit = defineEmits<Emits>()

// Reactivity
const count = ref(0)

// Derived state
const double = computed(() => count.value * 2)
</script>
```

### Templates
- Use semantic HTML elements
- Apply `scoped` styles to components
- Use template expressions sparingly; prefer methods/computed
- Accessibility first: proper ARIA labels, alt text, button roles

### Styling
- Use `<style scoped>` by default
- Prefer CSS custom properties for theming
- Organize styles: component vars → layout → typography → interactions

## File Structure
```
src/
├── App.vue           # Root component
├── main.ts           # Application entry point
├── components/       # Reusable components
└── [other directories as needed]
```

## Common Patterns

### Pinia Store Usage
- Create stores in `stores/` directory
- Use typed store definitions
- Follow: `const useUserStore = defineStore('user', { ... })`

### Vue Router Usage
- Lazy-load routes for code splitting
- Guard navigation with `beforeEach` hooks
- Use composables for route logic: `useRoute()`, `useRouter()`

### Internationalization
- Use `vue-i18n` for translations
- Extract strings to locale files under `/locales/`
- Use namespaced keys: `component.nested.key`

## Additional Notes

- No linter currently configured; maintain code quality manually
- Consider adding ESLint + Prettier for consistency
- Keep bundle size reasonable; analyze before adding dependencies
- All scripts assume Node.js 18+ environment
