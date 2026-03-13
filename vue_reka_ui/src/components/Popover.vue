<script setup lang="ts">
import {
  PopoverRoot,
  PopoverTrigger,
  PopoverPortal,
  PopoverContent,
  PopoverClose,
  PopoverArrow,
} from "reka-ui";
import { computed } from "vue";

interface Props {
  open?: boolean;
  side?: "top" | "right" | "bottom" | "left";
  align?: "start" | "center" | "end";
  sideOffset?: number;
}

interface Emits {
  (e: "update:open", value: boolean): void;
}

const props = withDefaults(defineProps<Props>(), {
  side: "bottom",
  align: "center",
  sideOffset: 4,
});

const emit = defineEmits<Emits>();

const isOpen = computed({
  get: () => props.open ?? false,
  set: (value: boolean) => emit("update:open", value),
});

const handleOpenChange = (value: boolean) => {
  isOpen.value = value;
};
</script>

<template>
  <PopoverRoot :open="isOpen" @update:open="handleOpenChange">
    <PopoverTrigger as-child>
      <slot name="trigger" />
    </PopoverTrigger>

    <PopoverPortal>
      <PopoverContent
        :side="side"
        :align="align"
        :side-offset="sideOffset"
        class="popover__content"
      >
        <div class="popover__header">
          <PopoverClose class="popover__close" aria-label="Close popover"
            >✕</PopoverClose
          >
        </div>

        <div class="popover__body">
          <slot />
        </div>

        <PopoverArrow class="popover__arrow" />
      </PopoverContent>
    </PopoverPortal>
  </PopoverRoot>
</template>

<style scoped lang="scss">
.popover__content {
  background-color: var(--popover);
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-lg);
  padding: 1rem;
  z-index: 50;
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

.popover__header {
  display: flex;
  justify-content: flex-end;
  margin-bottom: 0.5rem;
}

.popover__close {
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

.popover__body {
  font-size: 0.875rem;
  line-height: 1.6;
  color: var(--popover-foreground);
}

.popover__arrow {
  fill: var(--popover);
  stroke: var(--border);
  stroke-width: 1px;
}
</style>
