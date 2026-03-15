<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import Button from "@/components/Button.vue";
import Card from "@/components/Card.vue";
import Command from "@/components/Command.vue";
import Calendar from "@/components/Calendar.vue";
import DataTable from "@/components/DataTable.vue";
import type { CommandItem } from "@/components/Command.vue";
import type { Column } from "@/components/DataTable.vue";

interface Props {
  selectedComponent?: string;
}

defineProps<Props>();

const { t } = useI18n();

const commandOpen = ref(false);
const selectedDate = ref<Date | null>(null);

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
</script>

<template>
  <div class="component-demo">
    <!-- Command -->
    <div
      v-if="!selectedComponent || selectedComponent === 'command'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.command") }}</h3>
        </template>
        <div class="demo-content">
          <Button @click="commandOpen = !commandOpen">Open Command</Button>
          <Command
            v-if="commandOpen"
            :items="commandItems"
            placeholder="Search commands..."
          />
          <p class="demo-description">
            Command component for command palette and keyboard shortcuts
          </p>
        </div>
      </Card>
    </div>

    <!-- Calendar -->
    <div
      v-if="!selectedComponent || selectedComponent === 'calendar'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.calendar") }}</h3>
        </template>
        <div class="demo-content">
          <Calendar v-model="selectedDate" />
          <p class="demo-description">
            Calendar component for date selection and display
          </p>
        </div>
      </Card>
    </div>

    <!-- Data Table -->
    <div
      v-if="!selectedComponent || selectedComponent === 'dataTable'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.dataTable") }}</h3>
        </template>
        <div class="demo-content">
          <DataTable :columns="tableColumns" :data="tableData" />
          <p class="demo-description">
            DataTable component for displaying structured data with sorting and
            filtering
          </p>
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
</style>
