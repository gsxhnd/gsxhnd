<script setup lang="ts">
import { RadioGroupItem, RadioGroupIndicator } from "reka-ui";

interface Props {
  value: string;
  disabled?: boolean;
  required?: boolean;
  id?: string;
  name?: string;
  ariaLabel?: string;
}

withDefaults(defineProps<Props>(), {
  disabled: false,
  required: false,
});
</script>

<template>
  <div class="radio">
    <RadioGroupItem
      :value="value"
      :disabled="disabled"
      :required="required"
      :id="id"
      :name="name"
      :aria-label="ariaLabel"
      :class="['radio__control', { 'radio__control--disabled': disabled }]"
    >
      <RadioGroupIndicator :class="['radio__indicator']" />
    </RadioGroupItem>
  </div>
</template>

<style scoped lang="scss">
.radio {
  display: inline-flex;
  align-items: center;
}

.radio__control {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 1.25rem;
  height: 1.25rem;
  border: 1px solid var(--border);
  border-radius: var(--radius-full);
  background-color: var(--background);
  cursor: pointer;
  transition: all 0.2s ease;
  outline: none;

  &:hover:not(.radio__control--disabled) {
    border-color: var(--border-secondary);
    background-color: var(--primary-light);
  }

  &:focus-visible {
    box-shadow: 0 0 0 var(--focus-ring-width) var(--primary-light);
    border-color: var(--primary);
  }

  &[data-state="checked"] {
    border-color: var(--primary);
    background-color: var(--background);

    &:hover:not(.radio__control--disabled) {
      border-color: var(--primary-hover);
    }
  }

  &--disabled {
    cursor: not-allowed;
    opacity: 0.5;
    pointer-events: none;
  }
}

.radio__indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 0.5rem;
  height: 0.5rem;
  background-color: var(--primary);
  border-radius: var(--radius-full);

  [data-state="unchecked"] & {
    display: none;
  }
}
</style>
