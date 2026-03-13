<script setup lang="ts">
import { ref, computed } from "vue";

type SortDirection = "asc" | "desc" | null;

export interface Column {
  key: string;
  label: string;
  sortable?: boolean;
  width?: string;
  align?: "left" | "center" | "right";
}

interface Props {
  data: any[];
  columns: Column[];
  pageSize?: number;
  striped?: boolean;
  hoverable?: boolean;
  bordered?: boolean;
  size?: "sm" | "md" | "lg";
}

const props = withDefaults(defineProps<Props>(), {
  pageSize: 10,
  hoverable: true,
  size: "md",
});

const currentPage = ref(0);
const sortKey = ref<string | null>(null);
const sortDirection = ref<SortDirection>(null);

const sortedData = computed(() => {
  if (!sortKey.value || !sortDirection.value) {
    return props.data;
  }

  const sorted = [...props.data].sort((a: any, b: any) => {
    const aVal = a[sortKey.value!];
    const bVal = b[sortKey.value!];

    if (aVal === bVal) return 0;

    const comparison = aVal < bVal ? -1 : aVal > bVal ? 1 : 0;

    return sortDirection.value === "asc" ? comparison : -comparison;
  });

  return sorted;
});

const paginatedData = computed(() => {
  const start = currentPage.value * props.pageSize;
  const end = start + props.pageSize;
  return sortedData.value.slice(start, end);
});

const totalPages = computed(() => {
  return Math.ceil(sortedData.value.length / props.pageSize);
});

const handleSort = (key: string) => {
  const column = props.columns.find((col) => col.key === key);
  if (!column?.sortable) return;

  if (sortKey.value === key) {
    // Cycle through: asc -> desc -> null
    if (sortDirection.value === "asc") {
      sortDirection.value = "desc";
    } else if (sortDirection.value === "desc") {
      sortKey.value = null;
      sortDirection.value = null;
    }
  } else {
    sortKey.value = key;
    sortDirection.value = "asc";
  }
};

const getSortIcon = (key: string) => {
  if (sortKey.value !== key) return "↕";
  return sortDirection.value === "asc" ? "↑" : "↓";
};

const goToPage = (page: number) => {
  currentPage.value = Math.max(0, Math.min(page, totalPages.value - 1));
};
</script>

<template>
  <div :class="['data-table', `data-table--${size}`]">
    <div class="data-table__wrapper">
      <table
        :class="[
          'data-table__table',
          {
            'data-table__table--striped': striped,
            'data-table__table--bordered': bordered,
            'data-table__table--hoverable': hoverable,
          },
        ]"
      >
        <thead class="data-table__head">
          <tr>
            <th
              v-for="column in columns"
              :key="String(column.key)"
              :style="{ textAlign: column.align, width: column.width }"
              :class="{
                'data-table__header--sortable': column.sortable,
                'data-table__header--active': sortKey === String(column.key),
              }"
            >
              <button
                v-if="column.sortable"
                class="data-table__sort-button"
                @click="handleSort(String(column.key))"
              >
                {{ column.label }}
                <span class="data-table__sort-icon">
                  {{ getSortIcon(String(column.key)) }}
                </span>
              </button>
              <span v-else>{{ column.label }}</span>
            </th>
          </tr>
        </thead>

        <tbody class="data-table__body">
          <tr
            v-for="(row, index) in paginatedData"
            :key="index"
            class="data-table__row"
          >
            <td
              v-for="column in columns"
              :key="String(column.key)"
              :style="{ textAlign: column.align }"
              class="data-table__cell"
            >
              {{ row[column.key] }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <div v-if="totalPages > 1" class="data-table__pagination">
      <button
        class="data-table__pagination-button"
        @click="goToPage(currentPage - 1)"
        :disabled="currentPage === 0"
      >
        Previous
      </button>

      <div class="data-table__pagination-info">
        Page
        <span class="data-table__page-number">{{ currentPage + 1 }}</span> of
        <span class="data-table__page-total">{{ totalPages }}</span>
      </div>

      <button
        class="data-table__pagination-button"
        @click="goToPage(currentPage + 1)"
        :disabled="currentPage === totalPages - 1"
      >
        Next
      </button>
    </div>

    <div class="data-table__footer">
      <span class="data-table__count">
        Showing {{ paginatedData.length }} of {{ sortedData.length }} results
      </span>
    </div>
  </div>
</template>

<style scoped lang="scss">
.data-table {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  width: 100%;
  color: var(--foreground);
}

.data-table__wrapper {
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  overflow: hidden;
  box-shadow: var(--shadow-sm);
}

.data-table__table {
  width: 100%;
  border-collapse: collapse;
  font-size: 0.875rem;

  &--striped tbody tr:nth-child(even) {
    background-color: var(--background-muted);
  }

  &--bordered {
    border: 1px solid var(--border);

    td,
    th {
      border: 1px solid var(--border);
    }
  }

  &--hoverable tbody tr {
    transition: background-color 0.15s ease;

    &:hover {
      background-color: var(--background-muted);
    }
  }
}

.data-table__head {
  background-color: var(--background);
  border-bottom: 2px solid var(--border);
}

.data-table__head th {
  padding: 1rem;
  text-align: left;
  font-weight: 600;
  color: var(--foreground-secondary);
  vertical-align: middle;

  &.data-table__header--sortable {
    user-select: none;
  }

  &.data-table__header--active {
    background-color: var(--primary-light);
    color: var(--primary);
  }
}

.data-table__sort-button {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  border: none;
  background-color: transparent;
  color: inherit;
  cursor: pointer;
  font-weight: inherit;
  font-size: inherit;
  padding: 0;
  outline: none;
  transition: color 0.15s ease;

  &:hover {
    color: var(--primary);
  }

  &:focus-visible {
    outline: 2px solid var(--primary);
    outline-offset: 2px;
  }
}

.data-table__sort-icon {
  font-size: 0.75rem;
  opacity: 0.7;
  transition: opacity 0.15s ease;
}

.data-table__body {
  background-color: var(--popover);
}

.data-table__row {
  border-bottom: 1px solid var(--border);
  transition: background-color 0.15s ease;
}

.data-table__cell {
  padding: 1rem;
  vertical-align: middle;
}

.data-table--sm .data-table__cell,
.data-table--sm .data-table__head th {
  padding: 0.5rem 0.75rem;
  font-size: 0.8rem;
}

.data-table--lg .data-table__cell,
.data-table--lg .data-table__head th {
  padding: 1.25rem;
  font-size: 1rem;
}

.data-table__pagination {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 1rem;
  padding: 1rem;
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  background-color: var(--background);
}

.data-table__pagination-button {
  padding: 0.5rem 1rem;
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
  background-color: var(--popover);
  color: var(--foreground);
  cursor: pointer;
  font-size: 0.875rem;
  font-weight: 500;
  transition:
    background-color 0.15s ease,
    color 0.15s ease;
  outline: none;

  &:hover:not(:disabled) {
    background-color: var(--primary);
    color: var(--primary-foreground);
    border-color: var(--primary);
  }

  &:focus-visible:not(:disabled) {
    outline: 2px solid var(--primary);
    outline-offset: 2px;
  }

  &:disabled {
    color: var(--foreground-disabled);
    cursor: not-allowed;
    opacity: 0.5;
  }
}

.data-table__pagination-info {
  font-size: 0.875rem;
  color: var(--foreground-secondary);
  white-space: nowrap;
}

.data-table__page-number,
.data-table__page-total {
  font-weight: 600;
  color: var(--foreground);
}

.data-table__footer {
  padding: 0.75rem 1rem;
  text-align: center;
  font-size: 0.8rem;
  color: var(--foreground-muted);
}

.data-table__count {
  font-weight: 500;
}
</style>
