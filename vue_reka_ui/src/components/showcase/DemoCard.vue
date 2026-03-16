<script setup lang="ts">
interface Props {
  title: string;
  description?: string;
}

defineProps<Props>();
</script>

<template>
  <div class="demo-card">
    <div class="card-header">
      <h3>{{ title }}</h3>
    </div>
    <div class="card-content">
      <slot />
      <p v-if="description" class="card-description">{{ description }}</p>
    </div>
  </div>
</template>

<style scoped lang="scss">
.demo-card {
  background-color: var(--background-secondary);
  border: 1px solid var(--background-tertiary);
  border-radius: 0.75rem;
  overflow: hidden;

  .card-header {
    padding: 1rem 1.5rem;
    border-bottom: 1px solid var(--background-tertiary);

    h3 {
      margin: 0;
      font-size: 1.125rem;
      font-weight: 600;
      color: var(--foreground);
    }
  }

  .card-content {
    padding: 2rem;

    .card-description {
      margin-top: 1.5rem;
      padding-top: 1.5rem;
      border-top: 1px solid var(--background-tertiary);
      font-size: 0.875rem;
      color: var(--foreground-secondary);
    }
  }
}
</style>
