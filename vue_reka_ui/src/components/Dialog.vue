<script setup lang="ts">
import {
  DialogRoot,
  DialogTrigger,
  DialogPortal,
  DialogOverlay,
  DialogContent,
  DialogTitle,
  DialogDescription,
  DialogClose,
} from "reka-ui";
import { computed } from "vue";

interface Props {
  open?: boolean;
  title?: string;
  description?: string;
  size?: "sm" | "md" | "lg" | "xl";
}

interface Emits {
  (e: "update:open", value: boolean): void;
}

const props = withDefaults(defineProps<Props>(), {
  size: "md",
});

const emit = defineEmits<Emits>();

const isOpen = computed({
  get: () => props.open ?? false,
  set: (value: boolean) => emit("update:open", value),
});

const dialogContentClass = computed(() => {
  return ["dialog__content", `dialog__content--${props.size}`];
});

const handleOpenChange = (value: boolean) => {
  isOpen.value = value;
};
</script>

<template>
  <DialogRoot :open="isOpen" @update:open="handleOpenChange">
    <DialogTrigger as-child>
      <slot name="trigger" />
    </DialogTrigger>

    <DialogPortal>
      <DialogOverlay class="dialog__overlay" />
      <DialogContent :class="dialogContentClass">
        <div class="dialog__header">
          <DialogTitle v-if="title" class="dialog__title">{{
            title
          }}</DialogTitle>
          <DialogClose class="dialog__close" aria-label="Close dialog"
            >✕</DialogClose
          >
        </div>

        <DialogDescription v-if="description" class="dialog__description">
          {{ description }}
        </DialogDescription>

        <div class="dialog__body">
          <slot />
        </div>

        <div class="dialog__footer">
          <slot name="footer" />
        </div>
      </DialogContent>
    </DialogPortal>
  </DialogRoot>
</template>

<style scoped lang="scss">
.dialog__overlay {
  position: fixed;
  inset: 0;
  background-color: var(--overlay-background);
  animation: fade-in 0.2s ease;

  @keyframes fade-in {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
}

.dialog__content {
  position: fixed;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
  background-color: var(--modal-background);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-xl);
  display: flex;
  flex-direction: column;
  gap: 1rem;
  max-height: 90vh;
  overflow-y: auto;
  z-index: 50;
  animation: slide-in 0.2s ease;

  &--sm {
    width: 90%;
    max-width: 400px;
    padding: 1.5rem;
  }

  &--md {
    width: 90%;
    max-width: 512px;
    padding: 2rem;
  }

  &--lg {
    width: 90%;
    max-width: 640px;
    padding: 2rem;
  }

  &--xl {
    width: 90%;
    max-width: 896px;
    padding: 2rem;
  }

  @keyframes slide-in {
    from {
      opacity: 0;
      transform: translate(-50%, -48%);
    }
    to {
      opacity: 1;
      transform: translate(-50%, -50%);
    }
  }
}

.dialog__header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
}

.dialog__title {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  line-height: 1.5;
  color: var(--foreground);
}

.dialog__description {
  font-size: 0.875rem;
  line-height: 1.5;
  color: var(--foreground-secondary);
}

.dialog__close {
  flex-shrink: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 2rem;
  height: 2rem;
  padding: 0;
  border: none;
  background-color: transparent;
  color: var(--foreground-secondary);
  cursor: pointer;
  border-radius: var(--radius-md);
  transition: all 0.2s ease;
  outline: none;
  font-size: 1.25rem;

  &:hover {
    background-color: var(--background-muted);
    color: var(--foreground);
  }

  &:focus-visible {
    box-shadow: 0 0 0 var(--focus-ring-width) var(--primary-light);
  }
}

.dialog__body {
  color: var(--foreground);
  font-size: 0.875rem;
  line-height: 1.6;
}

.dialog__footer {
  display: flex;
  gap: 0.75rem;
  justify-content: flex-end;
  margin-top: 1rem;

  &:empty {
    display: none;
  }
}
</style>
