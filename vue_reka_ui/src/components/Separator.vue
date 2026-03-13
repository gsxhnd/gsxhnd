<script setup lang="ts">
import { computed } from "vue";

interface Props {
  orientation?: "horizontal" | "vertical";
  decorative?: boolean;
  ariaOrientation?: "horizontal" | "vertical";
}

const props = withDefaults(defineProps<Props>(), {
  orientation: "horizontal",
  decorative: true,
});

const separatorClass = computed(() => {
  return ["separator", `separator--${props.orientation}`];
});
</script>

<template>
  <div
    :orientation="orientation"
    :aria-orientation="ariaOrientation || orientation"
    :class="separatorClass"
    :aria-hidden="decorative"
  />
</template>

<style scoped lang="scss">
.separator {
  background-color: var(--border);
  flex-shrink: 0;

  &--horizontal {
    width: 100%;
    height: 1px;
    margin: 1rem 0;
  }

  &--vertical {
    width: 1px;
    height: 100%;
    margin: 0 1rem;
  }
}
</style>
