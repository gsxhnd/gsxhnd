<script setup lang="ts">
import { computed } from "vue";

interface Props {
  variant?:
    | "default"
    | "primary"
    | "secondary"
    | "success"
    | "warning"
    | "destructive"
    | "info";
  size?: "sm" | "md" | "lg";
  pill?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  variant: "default",
  size: "md",
  pill: false,
});

const badgeClass = computed(() => {
  return [
    "badge",
    `badge--${props.variant}`,
    `badge--${props.size}`,
    {
      "badge--pill": props.pill,
    },
  ];
});
</script>

<template>
  <span :class="badgeClass">
    <slot />
  </span>
</template>

<style scoped lang="scss">
.badge {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  font-weight: 600;
  white-space: nowrap;
  border-radius: var(--radius-md);
  transition: all 0.2s ease;
  line-height: 1;

  &--sm {
    padding: 0.25rem 0.5rem;
    font-size: 0.75rem;
  }

  &--md {
    padding: 0.375rem 0.75rem;
    font-size: 0.875rem;
  }

  &--lg {
    padding: 0.5rem 1rem;
    font-size: 1rem;
  }

  &--pill {
    border-radius: var(--radius-full);
  }

  // Variant colors
  &--default {
    background-color: var(--background-muted);
    color: var(--foreground);
    border: 1px solid var(--border);
  }

  &--primary {
    background-color: var(--primary);
    color: var(--primary-foreground);
  }

  &--secondary {
    background-color: var(--secondary);
    color: var(--secondary-foreground);
  }

  &--success {
    background-color: var(--success);
    color: var(--success-foreground);
  }

  &--warning {
    background-color: var(--warning);
    color: var(--warning-foreground);
  }

  &--destructive {
    background-color: var(--destructive);
    color: var(--destructive-foreground);
  }

  &--info {
    background-color: var(--info);
    color: var(--info-foreground);
  }
}
</style>
