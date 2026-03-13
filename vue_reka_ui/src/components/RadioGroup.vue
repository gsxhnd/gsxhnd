<script setup lang="ts">
import { RadioGroupRoot } from "reka-ui";

interface Props {
  modelValue?: string;
  defaultValue?: string;
  disabled?: boolean;
  required?: boolean;
  name?: string;
  orientation?: "horizontal" | "vertical";
}

interface Emits {
  (e: "update:modelValue", value: string): void;
}

withDefaults(defineProps<Props>(), {
  disabled: false,
  required: false,
  orientation: "vertical",
});

defineEmits<Emits>();
</script>

<template>
  <RadioGroupRoot
    :model-value="modelValue"
    :default-value="defaultValue"
    :disabled="disabled"
    :required="required"
    :name="name"
    :orientation="orientation"
    :class="['radio-group', `radio-group--${orientation}`]"
    @update:model-value="(value: any) => $emit('update:modelValue', value)"
  >
    <slot />
  </RadioGroupRoot>
</template>

<style scoped lang="scss">
.radio-group {
  display: flex;

  &--vertical {
    flex-direction: column;
    gap: 0.75rem;
  }

  &--horizontal {
    flex-direction: row;
    gap: 1.5rem;
    align-items: center;
  }
}
</style>
