<script setup lang="ts">
import {
  SelectRoot,
  SelectTrigger,
  SelectValue,
  SelectPortal,
  SelectContent,
  SelectViewport,
  SelectItem,
  SelectItemText,
  SelectItemIndicator,
  SelectGroup,
  SelectLabel,
} from "reka-ui";
import { computed } from "vue";

export interface SelectOption {
  label: string;
  value: string;
  disabled?: boolean;
}

export interface SelectGroup {
  label: string;
  options: SelectOption[];
}

type SelectContent = SelectOption | SelectGroup;

interface Props {
  modelValue?: string;
  defaultValue?: string;
  options: SelectContent[];
  placeholder?: string;
  disabled?: boolean;
  required?: boolean;
  size?: "sm" | "md" | "lg";
  name?: string;
  id?: string;
}

interface Emits {
  (e: "update:modelValue", value: string): void;
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: "Select an option...",
  disabled: false,
  required: false,
  size: "md",
});

const emit = defineEmits<Emits>();

const isGroup = (item: SelectContent): item is SelectGroup => {
  return "options" in item;
};

const triggerClass = computed(() => {
  return [
    "select-trigger",
    `select-trigger--${props.size}`,
    {
      "select-trigger--disabled": props.disabled,
    },
  ];
});
</script>

<template>
  <SelectRoot
    :model-value="modelValue"
    :default-value="defaultValue"
    :disabled="disabled"
    :required="required"
    :name="name"
    @update:model-value="(value) => $emit('update:modelValue', value)"
  >
    <SelectTrigger :id="id" :class="triggerClass">
      <SelectValue :placeholder="placeholder" />
    </SelectTrigger>

    <SelectPortal>
      <SelectContent class="select-content" position="popper">
        <SelectViewport class="select-viewport">
          <template v-for="(item, idx) in options" :key="idx">
            <!-- Grouped options -->
            <SelectGroup v-if="isGroup(item)" :class="['select-group']">
              <SelectLabel :class="['select-label']">{{
                item.label
              }}</SelectLabel>
              <SelectItem
                v-for="option in item.options"
                :key="option.value"
                :value="option.value"
                :disabled="option.disabled"
                :class="[
                  'select-item',
                  { 'select-item--disabled': option.disabled },
                ]"
              >
                <SelectItemIndicator :class="['select-item__indicator']">
                  <svg
                    viewBox="0 0 24 24"
                    fill="none"
                    stroke="currentColor"
                    stroke-width="2"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                  >
                    <polyline points="20 6 9 17 4 12" />
                  </svg>
                </SelectItemIndicator>
                <SelectItemText>{{ option.label }}</SelectItemText>
              </SelectItem>
            </SelectGroup>

            <!-- Ungrouped options -->
            <SelectItem
              v-else
              :value="item.value"
              :disabled="item.disabled"
              :class="[
                'select-item',
                { 'select-item--disabled': item.disabled },
              ]"
            >
              <SelectItemIndicator :class="['select-item__indicator']">
                <svg
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <polyline points="20 6 9 17 4 12" />
                </svg>
              </SelectItemIndicator>
              <SelectItemText>{{ item.label }}</SelectItemText>
            </SelectItem>
          </template>
        </SelectViewport>
      </SelectContent>
    </SelectPortal>
  </SelectRoot>
</template>

<style scoped lang="scss">
.select-trigger {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  padding: 0.5rem 0.75rem;
  font-size: 0.875rem;
  line-height: 1.5;
  color: var(--foreground);
  background-color: var(--input-background);
  border: 1px solid var(--input-border);
  border-radius: var(--radius-md);
  cursor: pointer;
  transition: all 0.2s ease;
  outline: none;

  &:hover:not(.select-trigger--disabled) {
    border-color: var(--input-border-focus);
  }

  &:focus-visible {
    border-color: var(--input-border-focus);
    box-shadow: 0 0 0 var(--focus-ring-width) var(--ring, rgba(0, 0, 0, 0.1));
  }

  &[data-state="open"] {
    border-color: var(--primary);
  }

  &--sm {
    padding: 0.375rem 0.5rem;
    font-size: 0.75rem;
  }

  &--lg {
    padding: 0.75rem 1rem;
    font-size: 1rem;
  }

  &--disabled {
    background-color: var(--background-muted);
    color: var(--foreground-disabled);
    cursor: not-allowed;
    opacity: 0.6;
  }
}

.select-content {
  overflow: hidden;
  background-color: var(--popover);
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-md);
}

.select-viewport {
  padding: 0.5rem 0;
  max-height: 300px;
}

.select-group {
  overflow: hidden;
}

.select-label {
  display: block;
  padding: 0.5rem 0.75rem;
  font-size: 0.75rem;
  font-weight: 600;
  line-height: 1.25;
  color: var(--foreground-muted);
  text-transform: uppercase;
  letter-spacing: 0.05em;
}

.select-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 0.75rem;
  cursor: pointer;
  transition: background-color 0.15s ease;
  outline: none;
  background-color: transparent;

  &[data-highlighted] {
    background-color: var(--background-secondary);
  }

  &:hover:not(.select-item--disabled) {
    background-color: var(--background-secondary);
  }

  &:focus-visible {
    background-color: var(--background-secondary);
  }

  &[data-state="checked"] {
    background-color: var(--background-hover);
    color: var(--foreground);
  }

  &--disabled {
    color: var(--foreground-disabled);
    cursor: not-allowed;
    opacity: 0.5;
  }
}

.select-item__indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 1rem;
  height: 1rem;
  color: var(--primary);

  svg {
    width: 100%;
    height: 100%;
  }

  [data-state="unchecked"] & {
    display: none;
  }
}
</style>
