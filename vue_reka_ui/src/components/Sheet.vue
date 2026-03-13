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
  side?: "top" | "right" | "bottom" | "left";
}

interface Emits {
  (e: "update:open", value: boolean): void;
}

const props = withDefaults(defineProps<Props>(), {
  side: "right",
});

const emit = defineEmits<Emits>();

const isOpen = computed({
  get: () => props.open ?? false,
  set: (value: boolean) => emit("update:open", value),
});

const sheetContentClass = computed(() => {
  return ["sheet__content", `sheet__content--${props.side}`];
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
      <DialogOverlay class="sheet__overlay" />
      <DialogContent :class="sheetContentClass">
        <div class="sheet__header">
          <DialogTitle v-if="title" class="sheet__title">{{
            title
          }}</DialogTitle>
          <DialogClose class="sheet__close" aria-label="Close sheet"
            >✕</DialogClose
          >
        </div>

        <DialogDescription v-if="description" class="sheet__description">
          {{ description }}
        </DialogDescription>

        <div class="sheet__body">
          <slot />
        </div>

        <div class="sheet__footer">
          <slot name="footer" />
        </div>
      </DialogContent>
    </DialogPortal>
  </DialogRoot>
</template>

<style scoped lang="scss">
.sheet__overlay {
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

.sheet__content {
  position: fixed;
  background-color: var(--modal-background);
  border: 1px solid var(--border);
  box-shadow: var(--shadow-xl);
  display: flex;
  flex-direction: column;
  gap: 1rem;
  z-index: 50;

  &--top {
    inset: 0 0 auto 0;
    max-height: 50vh;
    padding: 2rem;
    border-bottom: 1px solid var(--border);
    border-top: none;
    border-left: none;
    border-right: none;
    animation: slide-down-out 0.3s ease;

    @keyframes slide-down-out {
      from {
        transform: translateY(-100%);
      }
      to {
        transform: translateY(0);
      }
    }
  }

  &--right {
    inset: 0 0 0 auto;
    max-width: 512px;
    width: 100%;
    padding: 2rem;
    border-left: 1px solid var(--border);
    border-top: none;
    border-bottom: none;
    border-right: none;
    animation: slide-left-out 0.3s ease;

    @keyframes slide-left-out {
      from {
        transform: translateX(100%);
      }
      to {
        transform: translateX(0);
      }
    }

    @media (max-width: 640px) {
      max-width: none;
      width: 100%;
    }
  }

  &--bottom {
    inset: auto 0 0 0;
    max-height: 50vh;
    padding: 2rem;
    border-top: 1px solid var(--border);
    border-bottom: none;
    border-left: none;
    border-right: none;
    animation: slide-up-out 0.3s ease;

    @keyframes slide-up-out {
      from {
        transform: translateY(100%);
      }
      to {
        transform: translateY(0);
      }
    }
  }

  &--left {
    inset: 0 auto 0 0;
    max-width: 512px;
    width: 100%;
    padding: 2rem;
    border-right: 1px solid var(--border);
    border-top: none;
    border-bottom: none;
    border-left: none;
    animation: slide-right-out 0.3s ease;

    @keyframes slide-right-out {
      from {
        transform: translateX(-100%);
      }
      to {
        transform: translateX(0);
      }
    }

    @media (max-width: 640px) {
      max-width: none;
      width: 100%;
    }
  }
}

.sheet__header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 1rem;
}

.sheet__title {
  margin: 0;
  font-size: 1.25rem;
  font-weight: 600;
  line-height: 1.5;
  color: var(--foreground);
}

.sheet__description {
  font-size: 0.875rem;
  line-height: 1.5;
  color: var(--foreground-secondary);
}

.sheet__close {
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

.sheet__body {
  flex: 1;
  overflow-y: auto;
  color: var(--foreground);
  font-size: 0.875rem;
  line-height: 1.6;
}

.sheet__footer {
  display: flex;
  gap: 0.75rem;
  justify-content: flex-end;
  padding-top: 1rem;
  border-top: 1px solid var(--border);

  &:empty {
    display: none;
  }
}
</style>
