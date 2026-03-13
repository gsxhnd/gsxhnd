<script setup lang="ts">
import { CheckboxRoot, CheckboxIndicator } from "reka-ui";
import { computed } from "vue";

type CheckedState = boolean | "indeterminate";

interface Props {
  modelValue?: CheckedState;
  defaultChecked?: CheckedState;
  disabled?: boolean;
  required?: boolean;
  name?: string;
  value?: string;
  id?: string;
  ariaLabel?: string;
}

interface Emits {
  (e: "update:modelValue", value: CheckedState): void;
}

const props = withDefaults(defineProps<Props>(), {
  defaultChecked: false,
  disabled: false,
  required: false,
  value: "on",
});

const emit = defineEmits<Emits>();

const isControlled = computed(() => props.modelValue !== undefined);

const checked = computed({
  get: () => (isControlled.value ? props.modelValue : props.defaultChecked),
  set: (value: CheckedState) => {
    emit("update:modelValue", value);
  },
});

const handleUpdateChecked = (value: CheckedState) => {
  checked.value = value;
};
</script>

<template>
  <CheckboxRoot
    :checked="checked"
    :disabled="disabled"
    :required="required"
    :name="name"
    :value="value"
    :id="id"
    :aria-label="ariaLabel"
    :class="['checkbox', { 'checkbox--disabled': disabled }]"
    @update:checked="handleUpdateChecked"
  >
    <CheckboxIndicator :class="['checkbox__indicator']">
      <svg
        v-if="checked === true"
        class="checkbox__icon"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <polyline points="20 6 9 17 4 12" />
      </svg>
      <svg
        v-else-if="checked === 'indeterminate'"
        class="checkbox__icon"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <line x1="5" y1="12" x2="19" y2="12" />
      </svg>
    </CheckboxIndicator>
  </CheckboxRoot>
</template>

<style scoped lang="scss">
.checkbox {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 1.25rem;
  height: 1.25rem;
  border: 1px solid var(--border);
  border-radius: var(--radius-sm);
  background-color: var(--background);
  cursor: pointer;
  transition: all 0.2s ease;
  outline: none;

  &:hover:not(.checkbox--disabled) {
    border-color: var(--border-secondary);
    background-color: var(--primary-light);
  }

  &:focus-visible {
    box-shadow: 0 0 0 var(--focus-ring-width) var(--primary-light);
    border-color: var(--primary);
  }

  &[data-state="checked"] {
    background-color: var(--primary);
    border-color: var(--primary);
    color: var(--primary-foreground);

    &:hover:not(.checkbox--disabled) {
      background-color: var(--primary-hover);
      border-color: var(--primary-hover);
    }
  }

  &[data-state="indeterminate"] {
    background-color: var(--primary);
    border-color: var(--primary);
    color: var(--primary-foreground);
  }

  &--disabled {
    cursor: not-allowed;
    opacity: 0.5;
    pointer-events: none;
  }
}

.checkbox__indicator {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
  color: inherit;

  [data-state="unchecked"] & {
    display: none;
  }
}

.checkbox__icon {
  width: 0.875rem;
  height: 0.875rem;
  stroke-width: 3;
}
</style>
