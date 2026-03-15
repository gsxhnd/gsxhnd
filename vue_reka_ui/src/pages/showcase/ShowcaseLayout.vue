<script setup lang="ts">
import { useI18n } from "vue-i18n";
import Card from "@/components/Card.vue";
import Label from "@/components/Label.vue";
import Input from "@/components/Input.vue";
import Separator from "@/components/Separator.vue";

const { t } = useI18n();
</script>

<template>
  <div class="showcase-section">
    <h2 class="section-title">{{ t("showcase.sections.layout") }}</h2>
    <div class="section-grid">
      <!-- Card -->
      <Card class="component-card full-width">
        <template #header>
          <h3>{{ t("showcase.components.card") }}</h3>
        </template>
        <p>
          This is a card component. It provides a container for grouping related
          content.
        </p>
      </Card>

      <!-- Label -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.label") }}</h3>
        </template>
        <div class="demo-label">
          <Label for="demo">Form Label:</Label>
          <Input id="demo" placeholder="Associated input" />
        </div>
      </Card>

      <!-- Separator -->
      <Card class="component-card full-width">
        <template #header>
          <h3>{{ t("showcase.components.separator") }}</h3>
        </template>
        <div class="demo-separator">
          <p>Content before separator</p>
          <Separator />
          <p>Content after separator</p>
        </div>
      </Card>
    </div>
  </div>
</template>

<style scoped lang="scss">
.showcase-section {
  margin-bottom: 3rem;

  .section-title {
    font-size: 1.5rem;
    font-weight: 700;
    color: var(--foreground);
    margin: 0 0 2rem;
    padding-bottom: 0.5rem;
    border-bottom: 2px solid var(--primary);
    display: inline-block;
  }

  .section-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    margin-top: 2rem;

    .component-card {
      height: 100%;

      &.full-width {
        grid-column: 1 / -1;
      }
    }
  }
}

.demo-label,
.demo-separator {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}
</style>
