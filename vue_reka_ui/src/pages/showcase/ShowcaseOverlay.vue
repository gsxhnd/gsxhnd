<script setup lang="ts">
import { ref } from "vue";
import { useI18n } from "vue-i18n";
import Button from "@/components/Button.vue";
import Card from "@/components/Card.vue";
import Dialog from "@/components/Dialog.vue";
import Popover from "@/components/Popover.vue";
import Dropdown from "@/components/Dropdown.vue";
import Sheet from "@/components/Sheet.vue";
import type { DropdownItem } from "@/components/Dropdown.vue";

interface Props {
  selectedComponent?: string;
}

defineProps<Props>();

const { t } = useI18n();

const dialogOpen = ref(false);
const popoverOpen = ref(false);
const sheetOpen = ref(false);

const dropdownItems: DropdownItem[] = [
  { id: "edit", label: "Edit" },
  { id: "duplicate", label: "Duplicate" },
  { id: "separator1", separator: true },
  { id: "delete", label: "Delete" },
];
</script>

<template>
  <div class="component-demo">
    <!-- Dialog -->
    <div
      v-if="!selectedComponent || selectedComponent === 'dialog'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.dialog") }}</h3>
        </template>
        <div class="demo-content">
          <Button @click="dialogOpen = true">{{ t("common.confirm") }}</Button>
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
          <p class="demo-description">
            Dialog component for displaying modular content and confirmations
          </p>
        </div>
      </Card>
    </div>

    <!-- Popover -->
    <div
      v-if="!selectedComponent || selectedComponent === 'popover'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.popover") }}</h3>
        </template>
        <div class="demo-content">
          <Popover v-model="popoverOpen">
            <Button @click="popoverOpen = !popoverOpen">Open Popover</Button>
            <template #content>
              <div class="popover-content">
                <p>This is popover content</p>
              </div>
            </template>
          </Popover>
          <p class="demo-description">
            Popover component for displaying contextual floating content
          </p>
        </div>
      </Card>
    </div>

    <!-- Dropdown -->
    <div
      v-if="!selectedComponent || selectedComponent === 'dropdown'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.dropdown") }}</h3>
        </template>
        <div class="demo-content">
          <Dropdown :items="dropdownItems">
            <Button>Menu</Button>
          </Dropdown>
          <p class="demo-description">
            Dropdown component for displaying menu items with actions
          </p>
        </div>
      </Card>
    </div>

    <!-- Sheet -->
    <div
      v-if="!selectedComponent || selectedComponent === 'sheet'"
      class="demo-container"
    >
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.sheet") }}</h3>
        </template>
        <div class="demo-content">
          <Button @click="sheetOpen = true">{{ t("common.close") }}</Button>
          <Sheet v-model="sheetOpen" title="Sheet Panel">
            <p>This is the sheet content</p>
            <template #footer>
              <Button @click="sheetOpen = false">{{
                t("common.close")
              }}</Button>
            </template>
          </Sheet>
          <p class="demo-description">
            Sheet component for side panel content display
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

.popover-content {
  padding: 1rem;
  background-color: var(--background-secondary);
  border-radius: 0.5rem;
  border: 1px solid var(--background-tertiary);
}
</style>
