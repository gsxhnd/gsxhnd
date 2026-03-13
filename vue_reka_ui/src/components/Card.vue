<script setup lang="ts">
interface Props {
  hoverable?: boolean;
  clickable?: boolean;
}

withDefaults(defineProps<Props>(), {
  hoverable: false,
  clickable: false,
});
</script>

<template>
  <div
    :class="[
      'card',
      {
        'card--hoverable': hoverable,
        'card--clickable': clickable,
      },
    ]"
  >
    <header v-if="$slots.header" class="card__header">
      <slot name="header" />
    </header>

    <div class="card__body">
      <slot />
    </div>

    <footer v-if="$slots.footer" class="card__footer">
      <slot name="footer" />
    </footer>
  </div>
</template>

<style scoped lang="scss">
.card {
  background-color: var(--card);
  border: 1px solid var(--card-border);
  border-radius: var(--radius-lg);
  overflow: hidden;
  transition: all 0.2s ease;

  &--hoverable {
    &:hover {
      background-color: var(--card-hover);
      border-color: var(--border-secondary);
      box-shadow: var(--shadow-md);
    }
  }

  &--clickable {
    cursor: pointer;

    &:hover {
      background-color: var(--card-hover);
      border-color: var(--border-secondary);
      box-shadow: var(--shadow-md);
    }

    &:active {
      transform: scale(0.995);
    }
  }
}

.card__header {
  padding: 1rem 1.5rem;
  border-bottom: 1px solid var(--card-border);
  background-color: var(--background);

  &:empty {
    display: none;
  }
}

.card__body {
  padding: 1rem 1.5rem;
}

.card__footer {
  padding: 1rem 1.5rem;
  border-top: 1px solid var(--card-border);
  background-color: var(--background);
  display: flex;
  gap: 0.75rem;
  justify-content: flex-end;

  &:empty {
    display: none;
  }
}
</style>
