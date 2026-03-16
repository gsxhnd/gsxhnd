<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import Button from "@/components/Button.vue";
import Input from "@/components/Input.vue";
import Label from "@/components/Label.vue";
import Checkbox from "@/components/Checkbox.vue";
import RadioGroup from "@/components/RadioGroup.vue";
import Radio from "@/components/Radio.vue";
import Select from "@/components/Select.vue";
import Card from "@/components/Card.vue";
import Combobox from "@/components/Combobox.vue";
import type { SelectOption } from "@/components/Select.vue";

interface Props {
  selectedComponent?: string;
}

defineProps<Props>();

const { t } = useI18n();

const formData = ref({
  name: "",
  email: "",
  agreed: false,
  role: "",
  selectedColor: "blue",
});

const selectedComboboxValue = ref("");

const colorOptions: SelectOption[] = [
  { label: "Red", value: "red" },
  { label: "Blue", value: "blue" },
  { label: "Green", value: "green" },
];

const comboboxOptions = [
  { label: "Vue", value: "vue", disabled: false },
  { label: "React", value: "react", disabled: false },
  { label: "Svelte", value: "svelte", disabled: false },
];
</script>

<template>
  <div class="component-demo">
    <!-- Button -->
    <div
      v-if="!selectedComponent || selectedComponent === 'button'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.button") }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-buttons">
            <Button>{{ t("common.confirm") }}</Button>
            <Button variant="secondary">{{ t("common.cancel") }}</Button>
            <Button variant="destructive">{{ t("common.delete") }}</Button>
            <Button disabled>{{ t("common.loading") }}</Button>
          </div>
          <p class="demo-description">Button component with multiple variants and states</p>
        </div>
      </Card>
    </div>

    <!-- Input -->
    <div
      v-if="!selectedComponent || selectedComponent === 'input'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.input") }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-input">
            <Label for="input-demo">{{ t("common.name") }}:</Label>
            <Input
              id="input-demo"
              v-model="formData.name"
              placeholder="Enter name..."
            />
          </div>
          <p class="demo-description">Input component for text entry</p>
        </div>
      </Card>
    </div>

    <!-- Label -->
    <div
      v-if="!selectedComponent || selectedComponent === 'label'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.label") }}</h3>
        </template>
        <div class="demo-content">
          <Label for="label-demo">Example Label</Label>
          <p class="demo-description">Label component for form field labels</p>
        </div>
      </Card>
    </div>

    <!-- Checkbox -->
    <div
      v-if="!selectedComponent || selectedComponent === 'checkbox'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.checkbox") }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-checkbox">
            <Checkbox v-model="formData.agreed" id="agree" />
            <label for="agree">{{ t("common.confirm") }}</label>
          </div>
          <p class="demo-description">Checkbox component for boolean selections</p>
        </div>
      </Card>
    </div>

    <!-- Radio Group -->
    <div
      v-if="!selectedComponent || selectedComponent === 'radio'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.radio") }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-radio">
            <RadioGroup v-model="formData.role">
              <div class="radio-item">
                <Radio value="admin" id="admin" />
                <label for="admin">Admin</label>
              </div>
              <div class="radio-item">
                <Radio value="user" id="user" />
                <label for="user">User</label>
              </div>
            </RadioGroup>
          </div>
          <p class="demo-description">Radio component for single selection from multiple options</p>
        </div>
      </Card>
    </div>

    <!-- Select -->
    <div
      v-if="!selectedComponent || selectedComponent === 'select'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.select") }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-select">
            <Label for="color-select">Color:</Label>
            <Select
              id="color-select"
              v-model="formData.selectedColor"
              :options="colorOptions"
            />
          </div>
          <p class="demo-description">Select component for dropdown selections</p>
        </div>
      </Card>
    </div>

    <!-- Combobox -->
    <div
      v-if="!selectedComponent || selectedComponent === 'combobox'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.combobox") }}</h3>
        </template>
        <div class="demo-content">
          <div class="demo-combobox">
            <Combobox
              v-model="selectedComboboxValue"
              :options="comboboxOptions"
              placeholder="Select framework..."
            />
          </div>
          <p class="demo-description">Combobox component for searchable selections</p>
        </div>
      </Card>
    </div>
  </div>
</template>

<style scoped lang="scss">
.component-demo {
  display: flex;
  flex-direction: column;
  gap: 2rem;

  .demo-container {
    .component-card {
      height: auto;

      :deep(.card-header) {
        padding: 1rem 1.5rem;
        border-bottom: 1px solid var(--background-tertiary);

        h3 {
          margin: 0;
          font-size: 1.125rem;
          font-weight: 600;
          color: var(--foreground);
        }
      }

      .demo-content {
        padding: 2rem;

        .demo-description {
          margin-top: 1.5rem;
          padding-top: 1.5rem;
          border-top: 1px solid var(--background-tertiary);
          font-size: 0.875rem;
          color: var(--foreground-secondary);
        }
      }
    }
  }
}

.demo-buttons {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.demo-input,
.demo-select,
.demo-combobox {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.demo-checkbox,
.demo-radio {
  display: flex;
  align-items: center;
  gap: 0.5rem;

  label {
    cursor: pointer;
  }
}

.radio-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin: 0.5rem 0;

  label {
    cursor: pointer;
  }
}
</style>
