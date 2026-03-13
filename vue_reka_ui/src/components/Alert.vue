<script setup lang="ts">
import { computed } from "vue";

interface Props {
  variant?: "default" | "success" | "warning" | "destructive" | "info";
  title?: string;
  closeable?: boolean;
}

interface Emits {
  (e: "close"): void;
}

const props = withDefaults(defineProps<Props>(), {
  variant: "default",
  closeable: false,
});

const emit = defineEmits<Emits>();

const alertClass = computed(() => {
  return ["alert", `alert--${props.variant}`];
});

const iconMap = {
  success: "✓",
  warning: "⚠",
  destructive: "✕",
  info: "ℹ",
  default: "•",
};

const icon = computed(() => iconMap[props.variant as keyof typeof iconMap]);
</script>

<template>
  <div :class="alertClass" role="alert">
    <div class="alert__icon">{{ icon }}</div>

    <div class="alert__content">
      <h4 v-if="title" class="alert__title">{{ title }}</h4>
      <div class="alert__message">
        <slot />
      </div>
    </div>

    <button
      v-if="closeable"
      class="alert__close"
      aria-label="Close alert"
      @click="$emit('close')"
    >
      ✕
    </button>
  </div>
</template>

<style scoped lang="scss">
.alert {
  display: flex;
  gap: 1rem;
  align-items: flex-start;
  padding: 1rem;
  border-radius: var(--radius-md);
  border: 1px solid;
  transition: all 0.2s ease;

  &--default {
    background-color: var(--background-muted);
    border-color: var(--border);
    color: var(--foreground);

    .alert__icon {
      color: var(--foreground-secondary);
    }
  }

  &--success {
    background-color: var(--success-light);
    border-color: var(--success-muted);
    color: var(--foreground);

    .alert__icon {
      color: var(--success);
    }
  }

  &--warning {
    background-color: var(--warning-light);
    border-color: var(--warning-muted);
    color: var(--foreground);

    .alert__icon {
      color: var(--warning);
    }
  }

  &--destructive {
    background-color: var(--destructive-light);
    border-color: var(--destructive-muted);
    color: var(--foreground);

    .alert__icon {
      color: var(--destructive);
    }
  }

  &--info {
    background-color: var(--info-light);
    border-color: var(--info-muted);
    color: var(--foreground);

    .alert__icon {
      color: var(--info);
    }
  }
}

.alert__icon {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 1.25rem;
  height: 1.25rem;
  border-radius: var(--radius-md);
  font-weight: 600;
  font-size: 0.875rem;
}

.alert__content {
  flex: 1;
}

.alert__title {
  margin: 0 0 0.25rem 0;
  font-size: 0.875rem;
  font-weight: 600;
  line-height: 1.5;
}

.alert__message {
  font-size: 0.875rem;
  line-height: 1.5;
}

.alert__close {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 1.5rem;
  height: 1.5rem;
  padding: 0;
  border: none;
  background-color: transparent;
  color: var(--foreground-secondary);
  cursor: pointer;
  border-radius: var(--radius-sm);
  transition: all 0.2s ease;
  outline: none;
  font-size: 1rem;

  &:hover {
    background-color: var(--background-muted);
    color: var(--foreground);
  }

  &:focus-visible {
    box-shadow: 0 0 0 var(--focus-ring-width) var(--primary-light);
  }
}
</style>
