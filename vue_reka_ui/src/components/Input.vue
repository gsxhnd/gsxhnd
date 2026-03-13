<script setup lang="ts">
import { computed, ref } from "vue";

interface Props {
  modelValue?: string | number;
  type?:
    | "text"
    | "email"
    | "password"
    | "number"
    | "search"
    | "tel"
    | "url"
    | "date"
    | "time";
  placeholder?: string;
  disabled?: boolean;
  required?: boolean;
  readonly?: boolean;
  size?: "sm" | "md" | "lg";
  error?: boolean;
  errorMessage?: string;
  autocomplete?: string;
  maxlength?: number;
  minlength?: number;
  min?: number | string;
  max?: number | string;
  step?: number | string;
  name?: string;
  id?: string;
}

interface Emits {
  (e: "update:modelValue", value: string): void;
  (e: "blur"): void;
  (e: "focus"): void;
  (e: "change"): void;
}

const props = withDefaults(defineProps<Props>(), {
  type: "text",
  size: "md",
  disabled: false,
  required: false,
  readonly: false,
  error: false,
});

const emit = defineEmits<Emits>();

const isFocused = ref(false);

const inputClass = computed(() => {
  return [
    "input",
    `input--${props.size}`,
    {
      "input--error": props.error,
      "input--focused": isFocused.value,
      "input--disabled": props.disabled,
      "input--readonly": props.readonly,
    },
  ];
});

const handleInput = (e: Event) => {
  const target = e.target as HTMLInputElement;
  emit("update:modelValue", target.value);
};

const handleFocus = () => {
  isFocused.value = true;
  emit("focus");
};

const handleBlur = () => {
  isFocused.value = false;
  emit("blur");
};

const handleChange = () => {
  emit("change");
};
</script>

<template>
  <div class="input-wrapper">
    <input
      :id="id"
      :type="type"
      :name="name"
      :value="modelValue"
      :placeholder="placeholder"
      :disabled="disabled"
      :required="required"
      :readonly="readonly"
      :autocomplete="autocomplete"
      :maxlength="maxlength"
      :minlength="minlength"
      :min="min"
      :max="max"
      :step="step"
      :class="inputClass"
      v-bind="$attrs"
      @input="handleInput"
      @focus="handleFocus"
      @blur="handleBlur"
      @change="handleChange"
    />
    <span v-if="error && errorMessage" class="input__error">
      {{ errorMessage }}
    </span>
  </div>
</template>

<style scoped lang="scss">
.input-wrapper {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.input {
  display: block;
  width: 100%;
  padding: 0.5rem 0.75rem;
  font-family: inherit;
  font-size: 0.875rem;
  line-height: 1.5;
  color: var(--input-foreground);
  background-color: var(--input-background);
  border: 1px solid var(--input-border);
  border-radius: var(--radius-md);
  transition: all 0.2s ease;
  outline: none;

  &::placeholder {
    color: var(--input-placeholder);
  }

  &:focus {
    border-color: var(--input-border-focus);
    box-shadow: 0 0 0 var(--focus-ring-width) var(--ring, rgba(0, 0, 0, 0.1));
  }

  &--sm {
    padding: 0.375rem 0.5rem;
    font-size: 0.75rem;
  }

  &--lg {
    padding: 0.75rem 1rem;
    font-size: 1rem;
  }

  &--error {
    border-color: var(--destructive);
    color: var(--destructive);

    &:focus {
      box-shadow: 0 0 0 var(--focus-ring-width) var(--destructive-light);
    }
  }

  &--focused {
    border-color: var(--input-border-focus);
  }

  &--disabled {
    background-color: var(--background-muted);
    color: var(--foreground-disabled);
    cursor: not-allowed;
    opacity: 0.6;
  }

  &--readonly {
    background-color: var(--background-muted);
    cursor: default;
    user-select: none;
  }
}

.input__error {
  font-size: 0.75rem;
  color: var(--destructive);
  margin-top: 0.25rem;
}
</style>
