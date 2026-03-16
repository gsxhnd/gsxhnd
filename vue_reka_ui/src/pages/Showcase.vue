<script setup lang="ts">
import { useShowcase } from "@/composables/useShowcase";
import ShowcaseHeader from "@/components/showcase/ShowcaseHeader.vue";
import ShowcaseSidebar from "@/components/showcase/ShowcaseSidebar.vue";
import ShowcaseForm from "./showcase/ShowcaseForm.vue";
import ShowcaseFeedback from "./showcase/ShowcaseFeedback.vue";
import ShowcaseLayout from "./showcase/ShowcaseLayout.vue";
import ShowcaseOverlay from "./showcase/ShowcaseOverlay.vue";
import ShowcaseAdvanced from "./showcase/ShowcaseAdvanced.vue";

const {
  selectedCategory,
  selectedComponent,
  expandedCategories,
  componentCategories,
  selectComponent,
  toggleCategory,
} = useShowcase();
</script>

<template>
  <div class="showcase-wrapper">
    <ShowcaseHeader />

    <div class="showcase-layout">
      <ShowcaseSidebar
        :categories="componentCategories"
        :selected-category="selectedCategory"
        :selected-component="selectedComponent"
        :expanded-categories="expandedCategories"
        @select="selectComponent"
        @toggle="toggleCategory"
      />

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

.showcase-layout {
  display: flex;
  flex: 1;
  max-width: 1600px;
  width: 100%;
  margin: 0 auto;
  gap: 2rem;
  padding: 2rem;
}

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

@media (max-width: 1024px) {
  .showcase-layout {
    gap: 1.5rem;
    padding: 1.5rem;
  }
}

@media (max-width: 768px) {
  .showcase-layout {
    flex-direction: column;
    gap: 1rem;
    padding: 1rem;
  }

  .showcase-content {
    .content-container {
      padding: 1.5rem;
    }
  }
}
</style>
