<script setup lang="ts">
import { ref, computed } from "vue";

export interface CommandItem {
  id: string;
  label: string;
  icon?: string;
  shortcut?: string;
  category?: string;
  disabled?: boolean;
}

interface Props {
  items: CommandItem[];
  open?: boolean;
  placeholder?: string;
}

interface Emits {
  (e: "update:open", value: boolean): void;
  (e: "select", itemId: string): void;
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: "Type a command...",
});

const emit = defineEmits<Emits>();

const searchQuery = ref("");
const selectedIndex = ref(0);

const isOpen = computed({
  get: () => props.open ?? false,
  set: (value: boolean) => emit("update:open", value),
});

const grouped = computed(() => {
  const groups: Map<string, CommandItem[]> = new Map();

  for (const item of props.items) {
    const key = item.category || "General";
    if (!groups.has(key)) {
      groups.set(key, []);
    }
    groups.get(key)!.push(item);
  }

  return groups;
});

const filteredItems = computed(() => {
  const results: CommandItem[] = [];
  const query = searchQuery.value.toLowerCase();
  const groupMap = grouped.value;

  for (const [_, items] of groupMap.entries()) {
    for (const item of items) {
      if (
        !item.disabled &&
        (item.label.toLowerCase().includes(query) ||
          item.id.toLowerCase().includes(query))
      ) {
        results.push(item);
      }
    }
  }

  return results;
});

const handleSelect = (itemId: string) => {
  emit("select", itemId);
  isOpen.value = false;
  searchQuery.value = "";
};

const handleKeyDown = (e: KeyboardEvent) => {
  switch (e.key) {
    case "ArrowDown":
      e.preventDefault();
      selectedIndex.value = Math.min(
        selectedIndex.value + 1,
        filteredItems.value.length - 1,
      );
      break;
    case "ArrowUp":
      e.preventDefault();
      selectedIndex.value = Math.max(selectedIndex.value - 1, 0);
      break;
    case "Enter":
      e.preventDefault();
      if (
        selectedIndex.value >= 0 &&
        filteredItems.value[selectedIndex.value]
      ) {
        handleSelect(filteredItems.value[selectedIndex.value].id);
      }
      break;
    case "Escape":
      e.preventDefault();
      isOpen.value = false;
      break;
  }
};
</script>

<template>
  <div v-if="isOpen" class="command">
    <div class="command__overlay" @click="isOpen = false" />

    <div class="command__modal">
      <div class="command__input-container">
        <span class="command__icon">⌘</span>
        <input
          v-model="searchQuery"
          :placeholder="placeholder"
          class="command__input"
          type="text"
          @keydown="handleKeyDown"
          autofocus
        />
      </div>

      <div class="command__list">
        <div v-if="filteredItems.length === 0" class="command__empty">
          <p>No results found.</p>
          <p class="command__empty-hint">
            Try a different command or search term.
          </p>
        </div>

        <template v-for="(item, index) in filteredItems" :key="item.id">
          <button
            :class="[
              'command__item',
              {
                'command__item--selected': index === selectedIndex,
                'command__item--disabled': item.disabled,
              },
            ]"
            :disabled="item.disabled"
            @click="handleSelect(item.id)"
          >
            <div class="command__item-content">
              <span v-if="item.icon" class="command__item-icon">{{
                item.icon
              }}</span>
              <span class="command__item-label">{{ item.label }}</span>
            </div>
            <span v-if="item.shortcut" class="command__item-shortcut">
              {{ item.shortcut }}
            </span>
          </button>
        </template>
      </div>

      <div class="command__footer">
        <span class="command__hint"
          >Type to search • ↑↓ to navigate • ↵ to select • ESC to exit</span
        >
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.command {
  position: fixed;
  inset: 0;
  z-index: 50;
  display: flex;
  align-items: flex-start;
  justify-content: center;
  padding-top: 20vh;
}

.command__overlay {
  position: absolute;
  inset: 0;
  background-color: var(--overlay-background);
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

.command__modal {
  position: relative;
  background-color: var(--popover);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-xl);
  width: 90%;
  max-width: 640px;
  max-height: 60vh;
  display: flex;
  flex-direction: column;
  animation: slide-in 0.3s ease;
  overflow: hidden;

  @keyframes slide-in {
    from {
      opacity: 0;
      transform: translateY(-20px);
    }
    to {
      opacity: 1;
      transform: translateY(0);
    }
  }
}

.command__input-container {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  padding: 1rem 1.5rem;
  border-bottom: 1px solid var(--border);
}

.command__icon {
  color: var(--foreground-secondary);
  font-size: 1.25rem;
  font-weight: 600;
}

.command__input {
  flex: 1;
  border: none;
  background-color: transparent;
  color: var(--popover-foreground);
  font-size: 1rem;
  outline: none;

  &::placeholder {
    color: var(--foreground-muted);
  }
}

.command__list {
  flex: 1;
  overflow-y: auto;
  padding: 0.5rem 0;
}

.command__empty {
  padding: 2rem 1.5rem;
  text-align: center;
  color: var(--foreground-secondary);

  p {
    margin: 0.5rem 0;
    font-size: 0.875rem;

    &:first-child {
      font-weight: 500;
    }
  }
}

.command__empty-hint {
  color: var(--foreground-muted);
  font-size: 0.75rem;
}

.command__item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  padding: 0.75rem 1.5rem;
  border: none;
  background-color: transparent;
  color: var(--popover-foreground);
  cursor: pointer;
  text-align: left;
  font-size: 0.875rem;
  transition: background-color 0.15s ease;
  outline: none;

  &:hover:not(.command__item--disabled) {
    background-color: var(--background-muted);
  }

  &--selected {
    background-color: var(--primary-light);
    color: var(--primary);
  }

  &--disabled {
    color: var(--foreground-disabled);
    cursor: not-allowed;
    opacity: 0.5;
  }

  &:focus-visible {
    background-color: var(--primary-light);
    color: var(--primary);
  }
}

.command__item-content {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.command__item-icon {
  font-size: 1.125rem;
}

.command__item-label {
  font-weight: 500;
}

.command__item-shortcut {
  font-size: 0.75rem;
  color: var(--foreground-muted);
  font-family: monospace;
}

.command__footer {
  padding: 0.75rem 1.5rem;
  border-top: 1px solid var(--border);
  background-color: var(--background);
  font-size: 0.75rem;
  color: var(--foreground-muted);
  text-align: center;
}
</style>
