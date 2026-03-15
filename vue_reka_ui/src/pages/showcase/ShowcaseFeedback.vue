<script setup lang="ts">
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import Card from '@/components/Card.vue';
import Badge from '@/components/Badge.vue';
import Alert from '@/components/Alert.vue';
import Progress from '@/components/Progress.vue';
import Tabs from '@/components/Tabs.vue';
import type { Tab } from '@/components/Tabs.vue';

interface Props {
  selectedComponent?: string;
}

defineProps<Props>();

const { t } = useI18n();

const progressValue = ref(65);
const activeTab = ref('tab1');

const tabs: Tab[] = [
  { id: 'tab1', label: 'Overview' },
  { id: 'tab2', label: 'Details' },
  { id: 'tab3', label: 'Settings' },
];
</script>

<template>
  <div class="component-demo">
    <!-- Badge -->
    <div v-if="!selectedComponent || selectedComponent === 'badge'" class="demo-container">
      <Card class="component-card">
        <template #header>
          <h3>{{ t('showcase.components.badge') }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-badge">
            <Badge>New</Badge>
            <Badge variant="secondary">Secondary</Badge>
            <Badge variant="destructive">Danger</Badge>
          </div>
          <p class="demo-description">Badge component for displaying status tags and labels</p>
        </div>
      </Card>
    </div>

    <!-- Alert -->
    <div v-if="!selectedComponent || selectedComponent === 'alert'" class="demo-container">
      <Card class="component-card">
        <template #header>
          <h3>{{ t('showcase.components.alert') }}</h3>
        </template>
        <div class="demo-content">
          <Alert type="success">
            <strong>{{ t('common.success') }}!</strong> Operation completed successfully.
          </Alert>
          <p class="demo-description">Alert component for displaying messages and notifications</p>
        </div>
      </Card>
    </div>

    <!-- Progress -->
    <div v-if="!selectedComponent || selectedComponent === 'progress'" class="demo-container">
      <Card class="component-card">
        <template #header>
          <h3>{{ t('showcase.components.progress') }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-progress">
            <Progress :value="progressValue" />
            <p class="progress-text">{{ progressValue }}%</p>
          </div>
          <p class="demo-description">Progress component for displaying progress bars and loading states</p>
        </div>
      </Card>
    </div>

    <!-- Tabs -->
    <div v-if="!selectedComponent || selectedComponent === 'tabs'" class="demo-container">
      <Card class="component-card">
        <template #header>
          <h3>{{ t('showcase.components.tabs') }}</h3>
        </template>
        <div class="demo-content">
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
          <p class="demo-description">Tabs component for organizing content into tabbed sections</p>
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
