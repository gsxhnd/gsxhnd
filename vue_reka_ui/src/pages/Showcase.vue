<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import ShowcaseForm from "./showcase/ShowcaseForm.vue";
import ShowcaseFeedback from "./showcase/ShowcaseFeedback.vue";
import ShowcaseLayout from "./showcase/ShowcaseLayout.vue";
import ShowcaseOverlay from "./showcase/ShowcaseOverlay.vue";
import ShowcaseAdvanced from "./showcase/ShowcaseAdvanced.vue";

const { t, locale } = useI18n();
const currentTheme = ref((window as any).__theme?.current?.value || "dark");

const selectedCategory = ref("form");
const selectedComponent = ref("button");

interface ComponentCategory {
  id: string;
  label: string;
  components: { id: string; label: string }[];
}

const componentCategories: ComponentCategory[] = [
  {
    id: "form",
    label: t("showcase.sections.form"),
    components: [
      { id: "button", label: t("showcase.components.button") },
      { id: "input", label: t("showcase.components.input") },
      { id: "label", label: t("showcase.components.label") },
      { id: "checkbox", label: t("showcase.components.checkbox") },
      { id: "radio", label: t("showcase.components.radio") },
      { id: "select", label: t("showcase.components.select") },
      { id: "combobox", label: t("showcase.components.combobox") },
    ],
  },
  {
    id: "feedback",
    label: t("showcase.sections.feedback"),
    components: [
      { id: "badge", label: t("showcase.components.badge") },
      { id: "alert", label: t("showcase.components.alert") },
      { id: "progress", label: t("showcase.components.progress") },
      { id: "tabs", label: t("showcase.components.tabs") },
    ],
  },
  {
    id: "layout",
    label: t("showcase.sections.layout"),
    components: [
      { id: "card", label: t("showcase.components.card") },
      { id: "separator", label: t("showcase.components.separator") },
    ],
  },
  {
    id: "overlay",
    label: t("showcase.sections.overlay"),
    components: [
      { id: "dialog", label: t("showcase.components.dialog") },
      { id: "popover", label: t("showcase.components.popover") },
      { id: "dropdown", label: t("showcase.components.dropdown") },
      { id: "sheet", label: t("showcase.components.sheet") },
    ],
  },
  {
    id: "advanced",
    label: t("showcase.sections.advanced"),
    components: [
      { id: "command", label: t("showcase.components.command") },
      { id: "calendar", label: t("showcase.components.calendar") },
      { id: "dataTable", label: t("showcase.components.dataTable") },
    ],
  },
];

const expandedCategories = ref(new Set(["form"]));

const selectComponent = (categoryId: string, componentId: string) => {
  selectedCategory.value = categoryId;
  selectedComponent.value = componentId;
};

const toggleCategory = (categoryId: string) => {
  if (expandedCategories.value.has(categoryId)) {
    expandedCategories.value.delete(categoryId);
  } else {
    expandedCategories.value.add(categoryId);
  }
};

const toggleTheme = () => {
  const newTheme = currentTheme.value === "dark" ? "light" : "dark";
  currentTheme.value = newTheme;
  (window as any).__theme?.set?.(newTheme);
};

const switchLanguage = (lang: string) => {
  locale.value = lang;
};
</script>

<template>
  <div class="showcase-wrapper">
    <!-- Header -->
    <header class="showcase-header">
      <div class="header-content">
        <h1>{{ t("showcase.title") }}</h1>
        <div class="header-controls">
          <!-- Language Switcher -->
          <div class="control-group">
            <button
              :class="['lang-btn', { active: locale === 'zh' }]"
              @click="switchLanguage('zh')"
            >
              中文
            </button>
            <button
              :class="['lang-btn', { active: locale === 'en' }]"
              @click="switchLanguage('en')"
            >
              EN
            </button>
          </div>

          <!-- Theme Switcher -->
          <button class="theme-btn" @click="toggleTheme">
            {{ currentTheme === "dark" ? "☀️" : "🌙" }}
            {{
              currentTheme === "dark"
                ? t("showcase.lightMode")
                : t("showcase.darkMode")
            }}
          </button>
        </div>
      </div>
    </header>

    <!-- Main Layout -->
    <div class="showcase-layout">
      <!-- Sidebar -->
      <aside class="showcase-sidebar">
        <nav class="component-tree">
          <div
            v-for="category in componentCategories"
            :key="category.id"
            class="category"
          >
            <button
              class="category-header"
              :class="{ expanded: expandedCategories.has(category.id) }"
              @click="toggleCategory(category.id)"
            >
              <span class="toggle-icon">{{
                expandedCategories.has(category.id) ? "▼" : "▶"
              }}</span>
              <span class="category-label">{{ category.label }}</span>
              <span class="component-count">{{
                category.components.length
              }}</span>
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
                @click="selectComponent(category.id, component.id)"
              >
                {{ component.label }}
              </button>
            </div>
          </div>
        </nav>
      </aside>

      <!-- Content -->
      <main class="showcase-content">
        <div v-if="selectedCategory === 'form'" class="content-container">
          <ShowcaseForm :selected-component="selectedComponent" />
        </div>
        <div
          v-else-if="selectedCategory === 'feedback'"
          class="content-container"
        >
          <ShowcaseFeedback :selected-component="selectedComponent" />
        </div>
        <div
          v-else-if="selectedCategory === 'layout'"
          class="content-container"
        >
          <ShowcaseLayout :selected-component="selectedComponent" />
        </div>
        <div
          v-else-if="selectedCategory === 'overlay'"
          class="content-container"
        >
          <ShowcaseOverlay :selected-component="selectedComponent" />
        </div>
        <div
          v-else-if="selectedCategory === 'advanced'"
          class="content-container"
        >
          <ShowcaseAdvanced :selected-component="selectedComponent" />
        </div>
      </main>
    </div>
  </div>
</template>

<style scoped lang="scss">
.showcase-wrapper {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  background-color: var(--background);
  color: var(--foreground);
}

// Header
.showcase-header {
  background-color: var(--background-secondary);
  border-bottom: 1px solid var(--background-tertiary);
  padding: 1.5rem 2rem;
  position: sticky;
  top: 0;
  z-index: 100;

  .header-content {
    max-width: 1600px;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 2rem;

    h1 {
      margin: 0;
      font-size: 1.875rem;
      font-weight: 700;
      color: var(--foreground);
    }

    .header-controls {
      display: flex;
      gap: 1rem;
      align-items: center;

      .control-group {
        display: flex;
        gap: 0.5rem;
        border: 1px solid var(--background-tertiary);
        border-radius: 0.5rem;
        padding: 0.25rem;
        background-color: var(--background);

        .lang-btn {
          padding: 0.5rem 1rem;
          background: none;
          border: none;
          cursor: pointer;
          color: var(--foreground-secondary);
          font-size: 0.875rem;
          font-weight: 600;
          border-radius: 0.375rem;
          transition: all 0.2s ease;

          &.active {
            background-color: var(--primary);
            color: var(--primary-foreground);
          }

          &:hover:not(.active) {
            color: var(--foreground);
          }
        }
      }

      .theme-btn {
        padding: 0.5rem 1rem;
        background-color: var(--primary);
        color: var(--primary-foreground);
        border: none;
        border-radius: 0.5rem;
        cursor: pointer;
        font-size: 0.875rem;
        font-weight: 600;
        transition: all 0.2s ease;
        white-space: nowrap;

        &:hover {
          background-color: var(--primary-hover);
        }

        &:active {
          background-color: var(--primary-active);
        }
      }
    }
  }
}

// Main Layout
.showcase-layout {
  display: flex;
  flex: 1;
  max-width: 1600px;
  width: 100%;
  margin: 0 auto;
  gap: 2rem;
  padding: 2rem;
}

// Sidebar
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

        &.expanded {
          .toggle-icon {
            transform: none;
          }
        }

        .toggle-icon {
          flex-shrink: 0;
          width: 1rem;
          display: inline-flex;
          align-items: center;
          font-size: 0.75rem;
          transition: transform 0.2s ease;
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

// Content Area
.showcase-content {
  flex: 1;
  min-width: 0;

  .content-container {
    background-color: var(--background-secondary);
    border: 1px solid var(--background-tertiary);
    border-radius: 0.75rem;
    padding: 2rem;
    min-height: 400px;
  }
}

// Scrollbar styling
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

// Responsive Design
@media (max-width: 1024px) {
  .showcase-layout {
    gap: 1.5rem;
    padding: 1.5rem;
  }

  .showcase-sidebar {
    width: 240px;
  }
}

@media (max-width: 768px) {
  .showcase-layout {
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
  }

  .showcase-sidebar {
    width: 100%;
    max-height: 300px;
    position: static;
  }

  .showcase-header {
    padding: 1rem 1rem;

    .header-content {
      flex-direction: column;
      align-items: flex-start;
      gap: 1rem;

      h1 {
        font-size: 1.5rem;
      }

      .header-controls {
        width: 100%;
        justify-content: flex-start;
      }
    }
  }

  .showcase-content {
    .content-container {
      padding: 1.5rem;
    }
  }
}

@media (max-width: 480px) {
  .showcase-header {
    padding: 1rem;

    .header-content {
      h1 {
        font-size: 1.25rem;
      }

      .header-controls {
        flex-direction: column;
        gap: 0.75rem;

        .control-group {
          width: 100%;
          justify-content: center;
        }

        .theme-btn {
          width: 100%;
        }
      }
    }
  }
}
</style>
