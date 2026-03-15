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
import Separator from "@/components/Separator.vue";
import Tabs from "@/components/Tabs.vue";
import Alert from "@/components/Alert.vue";
import Badge from "@/components/Badge.vue";
import Progress from "@/components/Progress.vue";
import Dialog from "@/components/Dialog.vue";
import Popover from "@/components/Popover.vue";
import Dropdown from "@/components/Dropdown.vue";
import Sheet from "@/components/Sheet.vue";
import Combobox from "@/components/Combobox.vue";
import Command from "@/components/Command.vue";
import Calendar from "@/components/Calendar.vue";
import DataTable from "@/components/DataTable.vue";
import type { SelectOption } from "@/components/Select.vue";
import type { Tab } from "@/components/Tabs.vue";
import type { DropdownItem } from "@/components/Dropdown.vue";
import type { CommandItem } from "@/components/Command.vue";
import type { Column } from "@/components/DataTable.vue";

const { t, locale } = useI18n();
const currentTheme = ref((window as any).__theme?.current?.value || "dark");

// State for form components
const formData = ref({
  name: "",
  email: "",
  agreed: false,
  role: "",
  selectedColor: "blue",
});

// State for interactive components
const progressValue = ref(65);
const activeTab = ref("tab1");
const dialogOpen = ref(false);
const popoverOpen = ref(false);
const sheetOpen = ref(false);
const selectedComboboxValue = ref("");
const commandOpen = ref(false);
const selectedDate = ref<Date | null>(null);

const colorOptions: SelectOption[] = [
  { label: "Red", value: "red" },
  { label: "Blue", value: "blue" },
  { label: "Green", value: "green" },
];

const tabs: Tab[] = [
  { id: "tab1", label: "Overview" },
  { id: "tab2", label: "Details" },
  { id: "tab3", label: "Settings" },
];

const dropdownItems: DropdownItem[] = [
  { id: "edit", label: "Edit" },
  { id: "duplicate", label: "Duplicate" },
  { id: "separator1", separator: true },
  { id: "delete", label: "Delete" },
];

const comboboxOptions = [
  { label: "Vue", value: "vue", disabled: false },
  { label: "React", value: "react", disabled: false },
  { label: "Svelte", value: "svelte", disabled: false },
];

const commandItems: CommandItem[] = [
  {
    id: "new-file",
    label: "New File",
    icon: "📄",
    shortcut: "Cmd+N",
    category: "File",
  },
  {
    id: "new-folder",
    label: "New Folder",
    icon: "📁",
    shortcut: "Cmd+Shift+N",
    category: "File",
  },
];

interface TableRow {
  id: string;
  name: string;
  status: string;
  date: string;
}

const tableData: TableRow[] = [
  { id: "1", name: "Component 1", status: "Active", date: "2024-01-01" },
  { id: "2", name: "Component 2", status: "Inactive", date: "2024-01-02" },
];

const tableColumns: Column[] = [
  { key: "name", label: "Name", width: "200px" },
  { key: "status", label: "Status", width: "150px" },
  { key: "date", label: "Date", width: "150px" },
];

// Theme and language
const toggleTheme = () => {
  const newTheme = currentTheme.value === "dark" ? "light" : "dark";
  currentTheme.value = newTheme;
  (window as any).__theme?.set?.(newTheme);
};

const switchLanguage = (lang: string) => {
  locale.value = lang;
};
</script>

<template>
  <div class="showcase-container">
    <!-- Header -->
    <header class="showcase-header">
      <div class="header-content">
        <div class="header-left">
          <h1>{{ t("showcase.title") }}</h1>
          <p>{{ t("showcase.subtitle") }}</p>
        </div>
        <div class="header-right">
          <div class="controls">
            <!-- Language Switcher -->
            <div class="control-group">
              <button
                :class="['lang-btn', { active: locale === 'zh' }]"
                @click="switchLanguage('zh')"
              >
                中文
              </button>
              <button
                :class="['lang-btn', { active: locale === 'en' }]"
                @click="switchLanguage('en')"
              >
                EN
              </button>
            </div>

            <!-- Theme Switcher -->
            <button class="theme-btn" @click="toggleTheme">
              {{ currentTheme === "dark" ? "☀️" : "🌙" }}
              {{
                currentTheme === "dark"
                  ? t("showcase.lightMode")
                  : t("showcase.darkMode")
              }}
            </button>
          </div>
        </div>
      </div>
    </header>

    <!-- Main Content -->
    <main class="showcase-content">
      <!-- Form Section -->
      <section class="showcase-section">
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
      </section>

      <Separator />

      <!-- Feedback Section -->
      <section class="showcase-section">
        <h2 class="section-title">{{ t("showcase.sections.feedback") }}</h2>
        <div class="section-grid">
          <!-- Badge -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.badge") }}</h3>
            </template>
            <div class="demo-badge">
              <Badge>New</Badge>
              <Badge variant="secondary">Secondary</Badge>
              <Badge variant="destructive">Danger</Badge>
            </div>
          </Card>

          <!-- Alert -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.alert") }}</h3>
            </template>
            <Alert type="success">
              <strong>{{ t("common.success") }}!</strong> Operation completed
              successfully.
            </Alert>
          </Card>

          <!-- Progress -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.progress") }}</h3>
            </template>
            <div class="demo-progress">
              <Progress :value="progressValue" />
              <p class="progress-text">{{ progressValue }}%</p>
            </div>
          </Card>

          <!-- Tabs -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.tabs") }}</h3>
            </template>
            <Tabs v-model="activeTab" :tabs="tabs">
              <div v-if="activeTab === 'tab1'" class="tab-content">
                Tab 1 Content
              </div>
              <div v-if="activeTab === 'tab2'" class="tab-content">
                Tab 2 Content
              </div>
              <div v-if="activeTab === 'tab3'" class="tab-content">
                Tab 3 Content
              </div>
            </Tabs>
          </Card>
        </div>
      </section>

      <Separator />

      <!-- Layout Section -->
      <section class="showcase-section">
        <h2 class="section-title">{{ t("showcase.sections.layout") }}</h2>
        <div class="section-grid">
          <!-- Card -->
          <Card class="component-card full-width">
            <template #header>
              <h3>{{ t("showcase.components.card") }}</h3>
            </template>
            <p>
              This is a card component. It provides a container for grouping
              related content.
            </p>
          </Card>

          <!-- Label -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.label") }}</h3>
            </template>
            <div class="demo-label">
              <Label for="demo">Form Label:</Label>
              <Input id="demo" placeholder="Associated input" />
            </div>
          </Card>

          <!-- Separator -->
          <Card class="component-card full-width">
            <template #header>
              <h3>{{ t("showcase.components.separator") }}</h3>
            </template>
            <div class="demo-separator">
              <p>Content before separator</p>
              <Separator />
              <p>Content after separator</p>
            </div>
          </Card>
        </div>
      </section>

      <Separator />

      <!-- Overlay Section -->
      <section class="showcase-section">
        <h2 class="section-title">{{ t("showcase.sections.overlay") }}</h2>
        <div class="section-grid">
          <!-- Dialog -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.dialog") }}</h3>
            </template>
            <Button @click="dialogOpen = true">{{
              t("common.confirm")
            }}</Button>
            <Dialog v-model="dialogOpen" title="Confirm Action">
              <p>Are you sure you want to proceed?</p>
              <template #footer>
                <Button variant="secondary" @click="dialogOpen = false">
                  {{ t("common.cancel") }}
                </Button>
                <Button @click="dialogOpen = false">{{
                  t("common.confirm")
                }}</Button>
              </template>
            </Dialog>
          </Card>

          <!-- Popover -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.popover") }}</h3>
            </template>
            <Popover v-model="popoverOpen">
              <Button @click="popoverOpen = !popoverOpen">Open Popover</Button>
              <template #content>
                <div class="popover-content">
                  <p>This is popover content</p>
                </div>
              </template>
            </Popover>
          </Card>

          <!-- Dropdown -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.dropdown") }}</h3>
            </template>
            <Dropdown :items="dropdownItems">
              <Button>Menu</Button>
            </Dropdown>
          </Card>

          <!-- Sheet -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.sheet") }}</h3>
            </template>
            <Button @click="sheetOpen = true">{{ t("common.close") }}</Button>
            <Sheet v-model="sheetOpen" title="Sheet Panel">
              <p>This is the sheet content</p>
              <template #footer>
                <Button @click="sheetOpen = false">{{
                  t("common.close")
                }}</Button>
              </template>
            </Sheet>
          </Card>
        </div>
      </section>

      <Separator />

      <!-- Advanced Section -->
      <section class="showcase-section">
        <h2 class="section-title">{{ t("showcase.sections.advanced") }}</h2>
        <div class="section-grid">
          <!-- Command -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.command") }}</h3>
            </template>
            <Button @click="commandOpen = !commandOpen">Open Command</Button>
            <Command
              v-if="commandOpen"
              :items="commandItems"
              placeholder="Search commands..."
            />
          </Card>

          <!-- Calendar -->
          <Card class="component-card">
            <template #header>
              <h3>{{ t("showcase.components.calendar") }}</h3>
            </template>
            <Calendar v-model="selectedDate" />
          </Card>

          <!-- Data Table -->
          <Card class="component-card full-width">
            <template #header>
              <h3>{{ t("showcase.components.dataTable") }}</h3>
            </template>
            <DataTable :columns="tableColumns" :data="tableData" />
          </Card>
        </div>
      </section>
    </main>
  </div>
</template>

<style scoped lang="scss">
.showcase-container {
  min-height: 100vh;
  background-color: var(--background);
  color: var(--foreground);
}

// Header Styles
.showcase-header {
  background-color: var(--background-secondary);
  border-bottom: 1px solid var(--background-tertiary);
  padding: 2rem;
  position: sticky;
  top: 0;
  z-index: 100;

  .header-content {
    max-width: 1400px;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 2rem;

    .header-left {
      flex: 1;

      h1 {
        margin: 0;
        font-size: 2rem;
        font-weight: 700;
        color: var(--foreground);
      }

      p {
        margin: 0.5rem 0 0;
        font-size: 1rem;
        color: var(--foreground-secondary);
      }
    }

    .header-right {
      display: flex;
      gap: 1rem;
      align-items: center;

      .controls {
        display: flex;
        gap: 1rem;
        align-items: center;

        .control-group {
          display: flex;
          gap: 0.5rem;
          border: 1px solid var(--background-tertiary);
          border-radius: 0.5rem;
          padding: 0.25rem;
          background-color: var(--background);

          .lang-btn {
            padding: 0.5rem 1rem;
            background: none;
            border: none;
            cursor: pointer;
            color: var(--foreground-secondary);
            font-size: 0.875rem;
            font-weight: 600;
            border-radius: 0.375rem;
            transition: all 0.2s ease;

            &.active {
              background-color: var(--primary);
              color: var(--primary-foreground);
            }

            &:hover:not(.active) {
              color: var(--foreground);
            }
          }
        }

        .theme-btn {
          padding: 0.5rem 1rem;
          background-color: var(--primary);
          color: var(--primary-foreground);
          border: none;
          border-radius: 0.5rem;
          cursor: pointer;
          font-size: 0.875rem;
          font-weight: 600;
          transition: all 0.2s ease;
          white-space: nowrap;

          &:hover {
            background-color: var(--primary-hover);
          }

          &:active {
            background-color: var(--primary-active);
          }
        }
      }
    }
  }
}

// Main Content Styles
.showcase-content {
  max-width: 1400px;
  margin: 0 auto;
  padding: 3rem 2rem;

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

        &.full-width {
          grid-column: 1 / -1;
        }
      }
    }
  }
}

// Demo Content Styles
.demo-buttons {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.demo-input {
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

.demo-select,
.demo-combobox {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.demo-badge {
  display: flex;
  gap: 0.75rem;
  flex-wrap: wrap;
}

.demo-progress {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;

  .progress-text {
    font-size: 0.875rem;
    color: var(--foreground-secondary);
  }
}

.tab-content {
  padding: 1rem 0;
  color: var(--foreground-secondary);
}

.demo-label,
.demo-separator {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.popover-content {
  padding: 1rem;
  background-color: var(--background-secondary);
  border-radius: 0.5rem;
  border: 1px solid var(--background-tertiary);
}

// Responsive Design
@media (max-width: 768px) {
  .showcase-header {
    padding: 1.5rem 1rem;

    .header-content {
      flex-direction: column;
      align-items: flex-start;
      gap: 1rem;

      .header-left {
        h1 {
          font-size: 1.5rem;
        }

        p {
          font-size: 0.875rem;
        }
      }

      .header-right {
        width: 100%;
        justify-content: flex-end;
      }
    }
  }

  .showcase-content {
    padding: 2rem 1rem;

    .showcase-section {
      .section-grid {
        grid-template-columns: 1fr;
      }
    }
  }
}
</style>
