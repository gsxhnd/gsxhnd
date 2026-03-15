<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import Card from "@/components/Card.vue";
import Badge from "@/components/Badge.vue";
import Alert from "@/components/Alert.vue";
import Progress from "@/components/Progress.vue";
import Tabs from "@/components/Tabs.vue";
import type { Tab } from "@/components/Tabs.vue";

const { t } = useI18n();

const progressValue = ref(65);
const activeTab = ref("tab1");

const tabs: Tab[] = [
  { id: "tab1", label: "Overview" },
  { id: "tab2", label: "Details" },
  { id: "tab3", label: "Settings" },
];
</script>

<template>
  <div class="showcase-section">
    <h2 class="section-title">{{ t("showcase.sections.feedback") }}</h2>
    <div class="section-grid">
      <!-- Badge -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.badge") }}</h3>
        </template>
        <div class="demo-badge">
          <Badge>New</Badge>
          <Badge variant="secondary">Secondary</Badge>
          <Badge variant="destructive">Danger</Badge>
        </div>
      </Card>

      <!-- Alert -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.alert") }}</h3>
        </template>
        <Alert type="success">
          <strong>{{ t("common.success") }}!</strong> Operation completed
          successfully.
        </Alert>
      </Card>

      <!-- Progress -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.progress") }}</h3>
        </template>
        <div class="demo-progress">
          <Progress :value="progressValue" />
          <p class="progress-text">{{ progressValue }}%</p>
        </div>
      </Card>

      <!-- Tabs -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.tabs") }}</h3>
        </template>
        <Tabs v-model="activeTab" :tabs="tabs">
          <div v-if="activeTab === 'tab1'" class="tab-content">
            Tab 1 Content
          </div>
          <div v-if="activeTab === 'tab2'" class="tab-content">
            Tab 2 Content
          </div>
          <div v-if="activeTab === 'tab3'" class="tab-content">
            Tab 3 Content
          </div>
        </Tabs>
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
    }
  }
}

.demo-badge {
  display: flex;
  gap: 0.75rem;
  flex-wrap: wrap;
}

.demo-progress {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;

  .progress-text {
    font-size: 0.875rem;
    color: var(--foreground-secondary);
  }
}

.tab-content {
  padding: 1rem 0;
  color: var(--foreground-secondary);
}
</style>
