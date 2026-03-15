<script setup lang="ts">
import { useI18n } from "vue-i18n";
import Card from "@/components/Card.vue";
import Separator from "@/components/Separator.vue";

interface Props {
  selectedComponent?: string;
}

defineProps<Props>();

const { t } = useI18n();
</script>

<template>
  <div class="component-demo">
    <!-- Card -->
    <div
      v-if="!selectedComponent || selectedComponent === 'card'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.card") }}</h3>
        </template>
        <div class="demo-content">
          <Card class="nested-card">
            <template #header>
              <h4>Nested Card Example</h4>
            </template>
            <p>
              This is a card component. It provides a container for grouping
              related content.
            </p>
          </Card>
          <p class="demo-description">
            Card component for grouping and containing content with headers and
            styling
          </p>
        </div>
      </Card>
    </div>

    <!-- Separator -->
    <div
      v-if="!selectedComponent || selectedComponent === 'separator'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.separator") }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-separator">
            <p>Content before separator</p>
            <Separator />
            <p>Content after separator</p>
          </div>
          <p class="demo-description">
            Separator component for visual division between content sections
          </p>
        </div>
      </Card>
    </div>
  </div>
</template>

<style scoped lang="scss">
.component-demo {
  display: flex;
  flex-direction: column;
  gap: 2rem;

  .demo-container {
    .component-card {
      height: auto;

      :deep(.card-header) {
        padding: 1rem 1.5rem;
        border-bottom: 1px solid var(--background-tertiary);

        h3 {
          margin: 0;
          font-size: 1.125rem;
          font-weight: 600;
          color: var(--foreground);
        }
      }

      .demo-content {
        padding: 2rem;

        .demo-description {
          margin-top: 1.5rem;
          padding-top: 1.5rem;
          border-top: 1px solid var(--background-tertiary);
          font-size: 0.875rem;
          color: var(--foreground-secondary);
        }
      }
    }
  }
}

.nested-card {
  margin: 0 0 1.5rem 0;

  h4 {
    margin: 0;
    font-size: 1rem;
  }
}

.demo-separator {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
</style>
