<script setup lang="ts">
import { ref, computed, watch } from "vue";
import Input from "./Input.vue";

export interface ComboboxOption {
  label: string;
  value: string;
  disabled?: boolean;
}

interface Props {
  modelValue?: string;
  options: ComboboxOption[];
  placeholder?: string;
  disabled?: boolean;
  searchable?: boolean;
  clearable?: boolean;
  size?: "sm" | "md" | "lg";
}

interface Emits {
  (e: "update:modelValue", value: string): void;
}

const props = withDefaults(defineProps<Props>(), {
  placeholder: "Search...",
  disabled: false,
  searchable: true,
  clearable: true,
  size: "md",
});

const emit = defineEmits<Emits>();

const isOpen = ref(false);
const searchQuery = ref("");
const selectedIndex = ref(-1);

const filteredOptions = computed(() => {
  if (!props.searchable || !searchQuery.value) {
    return props.options;
  }
  const query = searchQuery.value.toLowerCase();
  return props.options.filter(
    (opt) => !opt.disabled && opt.label.toLowerCase().includes(query),
  );
});

const selectedOption = computed(() => {
  return props.options.find((opt) => opt.value === props.modelValue);
});

watch(isOpen, () => {
  if (!isOpen.value) {
    searchQuery.value = "";
    selectedIndex.value = -1;
  } else {
    selectedIndex.value = -1;
  }
});

const handleSelect = (value: string) => {
  emit("update:modelValue", value);
  isOpen.value = false;
  searchQuery.value = "";
};

const handleClear = () => {
  emit("update:modelValue", "");
  searchQuery.value = "";
};

const handleKeyDown = (e: KeyboardEvent) => {
  if (!isOpen.value) {
    if (e.key === "Enter" || e.key === " ") {
      isOpen.value = true;
      e.preventDefault();
    }
    return;
  }

  switch (e.key) {
    case "ArrowDown":
      e.preventDefault();
      selectedIndex.value = Math.min(
        selectedIndex.value + 1,
        filteredOptions.value.length - 1,
      );
      break;
    case "ArrowUp":
      e.preventDefault();
      selectedIndex.value = Math.max(selectedIndex.value - 1, -1);
      break;
    case "Enter":
      e.preventDefault();
      if (selectedIndex.value >= 0) {
        handleSelect(filteredOptions.value[selectedIndex.value].value);
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
  <div class="combobox">
    <div class="combobox__input-wrapper">
      <Input
        :value="searchQuery || selectedOption?.label || ''"
        :placeholder="placeholder"
        :disabled="disabled"
        :size="size"
        type="text"
        @input="
          (e: Event) => {
            searchQuery = (e.target as HTMLInputElement).value;
            isOpen = true;
          }
        "
        @focus="isOpen = true"
        @keydown="handleKeyDown"
      />

      <div class="combobox__actions">
        <button
          v-if="clearable && props.modelValue"
          class="combobox__clear"
          @click="handleClear"
          :disabled="disabled"
          aria-label="Clear selection"
        >
          ✕
        </button>
        <button
          class="combobox__toggle"
          @click="isOpen = !isOpen"
          :disabled="disabled"
          aria-label="Toggle dropdown"
        >
          ▼
        </button>
      </div>
    </div>

    <div v-if="isOpen" class="combobox__dropdown">
      <div class="combobox__options">
        <div v-if="filteredOptions.length === 0" class="combobox__empty">
          No results found
        </div>

        <button
          v-for="(option, index) in filteredOptions"
          :key="option.value"
          :class="[
            'combobox__option',
            {
              'combobox__option--selected': option.value === props.modelValue,
              'combobox__option--highlighted': index === selectedIndex,
              'combobox__option--disabled': option.disabled,
            },
          ]"
          :disabled="option.disabled"
          @click="handleSelect(option.value)"
        >
          {{ option.label }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped lang="scss">
.combobox {
  position: relative;
  width: 100%;
}

.combobox__input-wrapper {
  position: relative;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.combobox__actions {
  display: flex;
  gap: 0.25rem;
  margin-right: 0.5rem;
}

.combobox__clear,
.combobox__toggle {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 1.5rem;
  height: 1.5rem;
  padding: 0;
  border: none;
  background-color: transparent;
  color: var(--foreground-secondary);
  cursor: pointer;
  border-radius: var(--radius-sm);
  transition: all 0.2s ease;
  font-size: 0.875rem;
  outline: none;

  &:hover:not(:disabled) {
    background-color: var(--background-muted);
    color: var(--foreground);
  }

  &:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }

  &:focus-visible {
    box-shadow: 0 0 0 var(--focus-ring-width) var(--primary-light);
  }
}

.combobox__dropdown {
  position: absolute;
  top: calc(100% + 0.5rem);
  left: 0;
  right: 0;
  background-color: var(--popover);
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  box-shadow: var(--shadow-lg);
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

.combobox__options {
  max-height: 300px;
  overflow-y: auto;
  padding: 0.5rem 0;
}

.combobox__empty {
  padding: 1rem;
  text-align: center;
  font-size: 0.875rem;
  color: var(--foreground-muted);
}

.combobox__option {
  display: flex;
  align-items: center;
  width: 100%;
  padding: 0.625rem 1rem;
  text-align: left;
  font-size: 0.875rem;
  border: none;
  background-color: transparent;
  color: var(--popover-foreground);
  cursor: pointer;
  transition: background-color 0.15s ease;
  outline: none;

  &:hover:not(.combobox__option--disabled) {
    background-color: var(--background-muted);
  }

  &--highlighted {
    background-color: var(--primary-light);
    color: var(--primary);
  }

  &--selected {
    background-color: var(--primary-light);
    font-weight: 600;
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
</style>
