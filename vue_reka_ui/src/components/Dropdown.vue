<script setup lang="ts">
import {
  DropdownMenuRoot,
  DropdownMenuTrigger,
  DropdownMenuPortal,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuCheckboxItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
} from "reka-ui";

export interface DropdownItem {
  id: string;
  label?: string;
  disabled?: boolean;
  checkbox?: boolean;
  checked?: boolean;
  separator?: boolean;
}

interface Props {
  items: DropdownItem[];
  side?: "top" | "right" | "bottom" | "left";
  align?: "start" | "center" | "end";
  sideOffset?: number;
}

interface Emits {
  (e: "select", itemId: string): void;
  (e: "update:checked", itemId: string, checked: boolean): void;
}

withDefaults(defineProps<Props>(), {
  side: "bottom",
  align: "start",
  sideOffset: 4,
});

defineEmits<Emits>();
</script>

<template>
  <DropdownMenuRoot>
    <DropdownMenuTrigger as-child>
      <slot name="trigger" />
    </DropdownMenuTrigger>

    <DropdownMenuPortal>
      <DropdownMenuContent
        :side="side"
        :align="align"
        :side-offset="sideOffset"
        class="dropdown__content"
      >
        <template v-for="item in items" :key="item.id">
          <!-- Separator -->
          <DropdownMenuSeparator
            v-if="item.separator"
            class="dropdown__separator"
          />

          <!-- Label -->
          <DropdownMenuLabel v-else-if="!item.checkbox" class="dropdown__label">
            {{ item.label }}
          </DropdownMenuLabel>

          <!-- Checkbox Item -->
          <DropdownMenuCheckboxItem
            v-else-if="item.checkbox"
            :checked="item.checked"
            :disabled="item.disabled"
            :class="[
              'dropdown__item',
              { 'dropdown__item--disabled': item.disabled },
            ]"
            @update:checked="
              (checked: boolean) => $emit('update:checked', item.id, checked)
            "
          >
            <span class="dropdown__item-indicator">✓</span>
            {{ item.label }}
          </DropdownMenuCheckboxItem>

          <!-- Regular Item -->
          <DropdownMenuItem
            v-else
            :disabled="item.disabled"
            :class="[
              'dropdown__item',
              { 'dropdown__item--disabled': item.disabled },
            ]"
            @select="() => $emit('select', item.id)"
          >
            {{ item.label }}
          </DropdownMenuItem>
        </template>
      </DropdownMenuContent>
    </DropdownMenuPortal>
  </DropdownMenuRoot>
</template>

<style scoped lang="scss">
.dropdown__content {
  background-color: var(--popover);
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-lg);
  min-width: 180px;
  padding: 0.5rem 0;
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

.dropdown__item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 0.625rem 1rem;
  font-size: 0.875rem;
  line-height: 1.5;
  cursor: pointer;
  transition: background-color 0.15s ease;
  outline: none;
  color: var(--popover-foreground);
  position: relative;

  &:hover:not(.dropdown__item--disabled) {
    background-color: var(--background-muted);
  }

  &:focus-visible {
    background-color: var(--primary-light);
    color: var(--primary);
  }

  &--disabled {
    color: var(--foreground-disabled);
    cursor: not-allowed;
    opacity: 0.5;
  }

  [data-state="checked"] & {
    background-color: var(--primary-light);
    color: var(--primary);
  }
}

.dropdown__item-indicator {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 1rem;
  height: 1rem;
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--primary);

  [data-state="unchecked"] & {
    display: none;
  }
}

.dropdown__separator {
  height: 1px;
  background-color: var(--border);
  margin: 0.5rem 0;
}

.dropdown__label {
  display: block;
  padding: 0.5rem 1rem;
  font-size: 0.75rem;
  font-weight: 600;
  line-height: 1.25;
  color: var(--foreground-muted);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}
</style>
