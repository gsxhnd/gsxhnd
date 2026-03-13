<script setup lang="ts">
import { computed } from "vue";

interface Props {
  value: number;
  max?: number;
  variant?: "default" | "success" | "warning" | "destructive" | "info";
  size?: "sm" | "md" | "lg";
  showLabel?: boolean;
  animated?: boolean;
  striped?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  max: 100,
  variant: "default",
  size: "md",
  showLabel: false,
  animated: true,
  striped: false,
});

const percentage = computed(() => {
  const percent = Math.min(Math.max((props.value / props.max) * 100, 0), 100);
  return percent;
});

const progressClass = computed(() => {
  return [
    "progress",
    `progress--${props.size}`,
    {
      "progress--striped": props.striped,
      "progress--animated": props.animated && props.striped,
    },
  ];
});

const barClass = computed(() => {
  return ["progress__bar", `progress__bar--${props.variant}`];
});
</script>

<template>
  <div :class="progressClass">
    <div class="progress__background">
      <div :class="barClass" :style="{ width: `${percentage}%` }" />
    </div>
    <div v-if="showLabel" class="progress__label">
      {{ Math.round(percentage) }}%
    </div>
  </div>
</template>

<style scoped lang="scss">
.progress {
  display: flex;
  align-items: center;
  gap: 0.75rem;

  &--sm {
    .progress__background {
      height: 0.375rem;
    }
  }

  &--md {
    .progress__background {
      height: 0.5rem;
    }
  }

  &--lg {
    .progress__background {
      height: 0.75rem;
    }
  }

  &--striped {
    .progress__bar {
      background-image: linear-gradient(
        45deg,
        currentColor 25%,
        transparent 25%,
        transparent 50%,
        currentColor 50%,
        currentColor 75%,
        transparent 75%,
        transparent
      );
      background-size: 1rem 1rem;
    }
  }

  &--animated {
    .progress__bar {
      animation: progress-stripes 1s linear infinite;
    }
  }
}

.progress__background {
  flex: 1;
  width: 100%;
  background-color: var(--background-muted);
  border-radius: var(--radius-full);
  overflow: hidden;
  border: 1px solid var(--border);
}

.progress__bar {
  height: 100%;
  border-radius: var(--radius-full);
  transition: width 0.3s ease;
  background-color: var(--primary);

  &--default {
    background-color: var(--primary);
  }

  &--success {
    background-color: var(--success);
  }

  &--warning {
    background-color: var(--warning);
  }

  &--destructive {
    background-color: var(--destructive);
  }

  &--info {
    background-color: var(--info);
  }
}

.progress__label {
  min-width: 3rem;
  text-align: right;
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--foreground-secondary);
}

@keyframes progress-stripes {
  from {
    background-position: 0 0;
  }
  to {
    background-position: 1rem 1rem;
  }
}
</style>
