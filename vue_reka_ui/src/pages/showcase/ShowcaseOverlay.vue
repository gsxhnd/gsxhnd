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
  <div class="showcase-section">
    <h2 class="section-title">{{ t("showcase.sections.overlay") }}</h2>
    <div class="section-grid">
      <!-- Dialog -->
      <Card class="component-card">
        <template #header>
          <h3>{{ t("showcase.components.dialog") }}</h3>
        </template>
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
            <Button @click="sheetOpen = false">{{ t("common.close") }}</Button>
          </template>
        </Sheet>
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

.popover-content {
  padding: 1rem;
  background-color: var(--background-secondary);
  border-radius: 0.5rem;
  border: 1px solid var(--background-tertiary);
}
</style>
