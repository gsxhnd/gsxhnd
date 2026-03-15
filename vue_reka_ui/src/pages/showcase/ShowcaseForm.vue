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
  <div class="showcase-section">
    <h2 class="section-title">{{ t("showcase.sections.form") }}</h2>
    <div class="section-grid">
      <!-- Button -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.button") }}</h3>
        </template>
        <div class="demo-buttons">
          <Button>{{ t("common.confirm") }}</Button>
          <Button variant="secondary">{{ t("common.cancel") }}</Button>
          <Button variant="destructive">{{ t("common.delete") }}</Button>
          <Button disabled>{{ t("common.loading") }}</Button>
        </div>
      </Card>

      <!-- Input -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.input") }}</h3>
        </template>
        <div class="demo-input">
          <Label for="input-demo">{{ t("common.name") }}:</Label>
          <Input
            id="input-demo"
            v-model="formData.name"
            placeholder="Enter name..."
          />
        </div>
      </Card>

      <!-- Checkbox -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.checkbox") }}</h3>
        </template>
        <div class="demo-checkbox">
          <Checkbox v-model="formData.agreed" id="agree" />
          <label for="agree">{{ t("common.confirm") }}</label>
        </div>
      </Card>

      <!-- Radio Group -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.radio") }}</h3>
        </template>
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
      </Card>

      <!-- Select -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.select") }}</h3>
        </template>
        <div class="demo-select">
          <Label for="color-select">Color:</Label>
          <Select
            id="color-select"
            v-model="formData.selectedColor"
            :options="colorOptions"
          />
        </div>
      </Card>

      <!-- Combobox -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.combobox") }}</h3>
        </template>
        <div class="demo-combobox">
          <Combobox
            v-model="selectedComboboxValue"
            :options="comboboxOptions"
            placeholder="Select framework..."
          />
        </div>
      </Card>
    </div>
  </div>
</template>

<style scoped lang="scss">
.showcase-section {
  margin-bottom: 3rem;

  .section-title {
    font-size: 1.5rem;
    font-weight: 700;
    color: var(--foreground);
    margin: 0 0 2rem;
    padding-bottom: 0.5rem;
    border-bottom: 2px solid var(--primary);
    display: inline-block;
  }

  .section-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
    margin-top: 2rem;

    .component-card {
      height: 100%;
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
