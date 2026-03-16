<script setup lang="ts">
import type { ComponentCategory } from '@/composables/useShowcase';

interface Props {
  categories: ComponentCategory[];
  selectedCategory: string;
  selectedComponent: string;
  expandedCategories: Set<string>;
}

interface Emits {
  (e: 'select', categoryId: string, componentId: string): void;
  (e: 'toggle', categoryId: string): void;
}

defineProps<Props>();
const emit = defineEmits<Emits>();
</script>

<template>
  <aside class="showcase-sidebar">
    <nav class="component-tree">
      <div
        v-for="category in categories"
        :key="category.id"
        class="category"
      >
        <button
          class="category-header"
          :class="{ expanded: expandedCategories.has(category.id) }"
          @click="emit('toggle', category.id)"
        >
          <span class="toggle-icon">{{
            expandedCategories.has(category.id) ? '▼' : '▶'
          }}</span>
          <span class="category-label">{{ category.label }}</span>
          <span class="component-count">{{ category.components.length }}</span>
        </button>

        <div
          v-if="expandedCategories.has(category.id)"
          class="component-list"
        >
          <button
            v-for="component in category.components"
            :key="component.id"
            class="component-item"
            :class="{
              active:
                selectedComponent === component.id &&
                selectedCategory === category.id,
            }"
            @click="emit('select', category.id, component.id)"
          >
            {{ component.label }}
          </button>
        </div>
      </div>
    </nav>
  </aside>
</template>

<style scoped lang="scss">
.showcase-sidebar {
  width: 280px;
  flex-shrink: 0;
  background-color: var(--background-secondary);
  border: 1px solid var(--background-tertiary);
  border-radius: 0.75rem;
  padding: 1.5rem;
  max-height: calc(100vh - 180px);
  overflow-y: auto;
  position: sticky;
  top: 120px;

  .component-tree {
    .category {
      margin-bottom: 1rem;

      .category-header {
        width: 100%;
        padding: 0.75rem;
        background: none;
        border: none;
        cursor: pointer;
        display: flex;
        align-items: center;
        gap: 0.5rem;
        border-radius: 0.5rem;
        transition: all 0.2s ease;
        color: var(--foreground);
        font-weight: 600;
        font-size: 0.875rem;

        &:hover {
          background-color: var(--background-tertiary);
        }

        .toggle-icon {
          flex-shrink: 0;
          width: 1rem;
          display: inline-flex;
          align-items: center;
          font-size: 0.75rem;
        }

        .category-label {
          flex: 1;
          text-align: left;
        }

        .component-count {
          background-color: var(--primary);
          color: var(--primary-foreground);
          padding: 0.125rem 0.5rem;
          border-radius: 0.25rem;
          font-size: 0.75rem;
          font-weight: 700;
        }
      }

      .component-list {
        display: flex;
        flex-direction: column;
        gap: 0.25rem;
        padding-left: 1.5rem;
        margin-top: 0.5rem;

        .component-item {
          padding: 0.5rem 0.75rem;
          background: none;
          border: none;
          cursor: pointer;
          text-align: left;
          color: var(--foreground-secondary);
          font-size: 0.875rem;
          border-radius: 0.375rem;
          transition: all 0.2s ease;

          &:hover {
            color: var(--foreground);
            background-color: var(--background-tertiary);
          }

          &.active {
            background-color: var(--primary);
            color: var(--primary-foreground);
            font-weight: 600;
          }
        }
      }
    }
  }
}

.showcase-sidebar::-webkit-scrollbar {
  width: 6px;
}

.showcase-sidebar::-webkit-scrollbar-track {
  background: transparent;
}

.showcase-sidebar::-webkit-scrollbar-thumb {
  background-color: var(--background-tertiary);
  border-radius: 3px;

  &:hover {
    background-color: var(--background-secondary);
  }
}

@media (max-width: 1024px) {
  .showcase-sidebar {
    width: 240px;
  }
}

@media (max-width: 768px) {
  .showcase-sidebar {
    width: 100%;
    max-height: 300px;
    position: static;
  }
}
</style>
