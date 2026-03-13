<script setup lang="ts">
export interface Props {
  for?: string;
  required?: boolean;
  disabled?: boolean;
}

withDefaults(defineProps<Props>(), {
  required: false,
  disabled: false,
});
</script>

<template>
  <label
    :for="for"
    :class="[
      'label',
      { 'label--required': required, 'label--disabled': disabled },
    ]"
  >
    <slot />
    <span v-if="required" class="label__required" aria-label="required">*</span>
  </label>
</template>

<style scoped lang="scss">
.label {
  display: inline-block;
  font-weight: 500;
  font-size: 0.875rem;
  line-height: 1.5;
  color: var(--foreground);
  cursor: pointer;
  transition: color 0.2s ease;

  &:hover:not(.label--disabled) {
    color: var(--foreground-secondary);
  }

  &--required {
    .label__required {
      margin-left: 0.25rem;
      color: var(--destructive);
    }
  }

  &--disabled {
    cursor: not-allowed;
    color: var(--foreground-disabled);
  }
}

.label__required {
  color: var(--destructive);
  margin-left: 0.25rem;
}
</style>
