<script setup lang="ts">
import { computed } from "vue";

interface Props {
  size?: "sm" | "md" | "lg" | "xl";
  variant?:
    | "default"
    | "primary"
    | "secondary"
    | "outline"
    | "ghost"
    | "destructive";
  disabled?: boolean;
  loading?: boolean;
  fullWidth?: boolean;
  iconOnly?: boolean;
  type?: "button" | "submit" | "reset";
  asChild?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
  size: "md",
  variant: "primary",
  type: "button",
  disabled: false,
  loading: false,
  fullWidth: false,
  iconOnly: false,
  asChild: false,
});

const buttonClass = computed(() => {
  const baseClasses = ["btn"];

  // Size variants
  if (props.size !== "md") {
    baseClasses.push(`btn--${props.size}`);
  }

  // Style variants
  if (props.variant !== "primary") {
    baseClasses.push(`btn-${props.variant}`);
  } else {
    baseClasses.push("btn-primary");
  }

  // Modifiers
  if (props.fullWidth) baseClasses.push("btn--full");
  if (props.iconOnly) baseClasses.push("btn--icon");
  if (props.loading) baseClasses.push("btn--loading");
  if (props.disabled) baseClasses.push("btn--disabled");

  return baseClasses.join(" ");
});
</script>

<template>
  <button
    :type="type"
    :disabled="disabled || loading"
    :class="buttonClass"
    v-bind="$attrs"
  >
    <slot />
  </button>
</template>

<style scoped lang="scss">
// Button styles are imported from index.scss
// which references the design system CSS variables
// from src/style/_button.scss
</style>
