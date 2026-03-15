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
  <div class="showcase-section">
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

      &.full-width {
        grid-column: 1 / -1;
      }
    }
  }
}
</style>
