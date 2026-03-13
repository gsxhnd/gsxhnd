<script setup lang="ts">
import { TabsRoot, TabsList, TabsTrigger, TabsContent } from "reka-ui";
import { computed } from "vue";

export interface Tab {
  id: string;
  label: string;
  disabled?: boolean;
}

interface Props {
  modelValue?: string;
  defaultValue?: string;
  tabs: Tab[];
  orientation?: "horizontal" | "vertical";
  activationMode?: "automatic" | "manual";
  size?: "sm" | "md" | "lg";
}

interface Emits {
  (e: "update:modelValue", value: string): void;
}

const props = withDefaults(defineProps<Props>(), {
  orientation: "horizontal",
  activationMode: "automatic",
  size: "md",
});

const emit = defineEmits<Emits>();

const tabsListClass = computed(() => {
  return ["tabs__list", `tabs__list--${props.orientation}`];
});

const tabsTriggerClass = (disabled?: boolean) => {
  return [
    "tabs__trigger",
    `tabs__trigger--${props.size}`,
    { "tabs__trigger--disabled": disabled },
  ];
};

const tabsContentClass = () => {
  return ["tabs__content"];
};
</script>

<template>
  <TabsRoot
    :model-value="modelValue"
    :default-value="defaultValue"
    :orientation="orientation"
    :activation-mode="activationMode"
    :class="'tabs'"
    @update:model-value="(value) => $emit('update:modelValue', value)"
  >
    <TabsList :class="tabsListClass">
      <TabsTrigger
        v-for="tab in tabs"
        :key="tab.id"
        :value="tab.id"
        :disabled="tab.disabled"
        :class="tabsTriggerClass(tab.disabled)"
      >
        {{ tab.label }}
      </TabsTrigger>
    </TabsList>

    <template v-for="tab in tabs" :key="tab.id">
      <TabsContent :value="tab.id" :class="tabsContentClass()">
        <slot :name="`content-${tab.id}`" />
      </TabsContent>
    </template>
  </TabsRoot>
</template>

<style scoped lang="scss">
.tabs {
  width: 100%;
}

.tabs__list {
  display: flex;
  gap: 0.5rem;
  border-bottom: 1px solid var(--border);
  background-color: var(--background);

  &--vertical {
    flex-direction: column;
    border-bottom: none;
    border-right: 1px solid var(--border);
    min-width: 150px;
  }
}

.tabs__trigger {
  position: relative;
  padding: 0.75rem 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  line-height: 1.5;
  color: var(--foreground-secondary);
  background-color: transparent;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease;
  outline: none;
  white-space: nowrap;

  &:hover:not(.tabs__trigger--disabled) {
    color: var(--foreground);
    background-color: var(--background-muted);
  }

  &:focus-visible {
    box-shadow: 0 0 0 var(--focus-ring-width) var(--primary-light);
  }

  &[data-state="active"] {
    color: var(--foreground);
    border-bottom: 2px solid var(--primary);
    margin-bottom: -1px;
  }

  &--disabled {
    color: var(--foreground-disabled);
    cursor: not-allowed;
    opacity: 0.5;
  }

  &--sm {
    padding: 0.5rem 0.75rem;
    font-size: 0.75rem;
  }

  &--lg {
    padding: 1rem 1.5rem;
    font-size: 1rem;
  }
}

.tabs__content {
  margin-top: 1rem;
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
</style>
