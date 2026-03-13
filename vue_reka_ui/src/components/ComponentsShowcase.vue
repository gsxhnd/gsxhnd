<script setup lang="ts">
import { ref } from "vue";
import Button from "./Button.vue";
import Input from "./Input.vue";
import Label from "./Label.vue";
import Checkbox from "./Checkbox.vue";
import RadioGroup from "./RadioGroup.vue";
import Radio from "./Radio.vue";
import Select from "./Select.vue";
import Card from "./Card.vue";
import Separator from "./Separator.vue";
import Tabs from "./Tabs.vue";
import Alert from "./Alert.vue";
import Badge from "./Badge.vue";
import Progress from "./Progress.vue";
import Dialog from "./Dialog.vue";
import Popover from "./Popover.vue";
import Dropdown from "./Dropdown.vue";
import Sheet from "./Sheet.vue";
import Combobox from "./Combobox.vue";
import Command from "./Command.vue";
import Calendar from "./Calendar.vue";
import DataTable from "./DataTable.vue";
import type { SelectOption } from "./Select.vue";
import type { Tab } from "./Tabs.vue";
import type { DropdownItem } from "./Dropdown.vue";
import type { CommandItem } from "./Command.vue";
import type { Column } from "./DataTable.vue";

// Form state
const name = ref("");
const email = ref("");
const agreed = ref(false);
const role = ref("");
const selectedColor = ref("blue");

// Phase 2 state
const activeTab = ref("tab1");
const progressValue = ref(65);
const closeAlert = ref(false);

// Phase 3 state
const dialogOpen = ref(false);
const popoverOpen = ref(false);
const sheetOpen = ref(false);
const sheetSide = ref<"top" | "right" | "bottom" | "left">("right");

const colorOptions: SelectOption[] = [
  { label: "Red", value: "red" },
  { label: "Blue", value: "blue" },
  { label: "Green", value: "green" },
  { label: "Yellow", value: "yellow" },
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
  { id: "archive", label: "Archive" },
  { id: "delete", label: "Delete", checkbox: false },
];

// Stage 4 state
const selectedComboboxValue = ref("");
const commandOpen = ref(false);
const selectedDate = ref<Date | null>(null);

const comboboxOptions = [
  { label: "Vue", value: "vue", disabled: false },
  { label: "React", value: "react", disabled: false },
  { label: "Svelte", value: "svelte", disabled: false },
  { label: "Angular", value: "angular", disabled: false },
  { label: "Next.js", value: "next", disabled: false },
  { label: "Nuxt", value: "nuxt", disabled: false },
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
  {
    id: "open-file",
    label: "Open File",
    icon: "📂",
    shortcut: "Cmd+O",
    category: "File",
  },
  {
    id: "save-file",
    label: "Save File",
    icon: "💾",
    shortcut: "Cmd+S",
    category: "File",
  },
  {
    id: "cut",
    label: "Cut",
    icon: "✂️",
    shortcut: "Cmd+X",
    category: "Edit",
  },
  {
    id: "copy",
    label: "Copy",
    icon: "📋",
    shortcut: "Cmd+C",
    category: "Edit",
  },
  {
    id: "paste",
    label: "Paste",
    icon: "📌",
    shortcut: "Cmd+V",
    category: "Edit",
  },
  {
    id: "undo",
    label: "Undo",
    icon: "↶",
    shortcut: "Cmd+Z",
    category: "Edit",
  },
];

const tableData = [
  {
    id: 1,
    name: "Alice Johnson",
    email: "alice@example.com",
    status: "Active",
    joinDate: "2024-01-15",
  },
  {
    id: 2,
    name: "Bob Smith",
    email: "bob@example.com",
    status: "Active",
    joinDate: "2024-02-20",
  },
  {
    id: 3,
    name: "Carol White",
    email: "carol@example.com",
    status: "Inactive",
    joinDate: "2024-01-10",
  },
  {
    id: 4,
    name: "David Brown",
    email: "david@example.com",
    status: "Active",
    joinDate: "2024-03-05",
  },
  {
    id: 5,
    name: "Eve Davis",
    email: "eve@example.com",
    status: "Active",
    joinDate: "2024-02-14",
  },
  {
    id: 6,
    name: "Frank Miller",
    email: "frank@example.com",
    status: "Inactive",
    joinDate: "2024-01-25",
  },
  {
    id: 7,
    name: "Grace Lee",
    email: "grace@example.com",
    status: "Active",
    joinDate: "2024-03-10",
  },
  {
    id: 8,
    name: "Henry Wilson",
    email: "henry@example.com",
    status: "Active",
    joinDate: "2024-02-28",
  },
  {
    id: 9,
    name: "Iris Taylor",
    email: "iris@example.com",
    status: "Active",
    joinDate: "2024-01-30",
  },
  {
    id: 10,
    name: "Jack Anderson",
    email: "jack@example.com",
    status: "Inactive",
    joinDate: "2024-03-01",
  },
  {
    id: 11,
    name: "Kate Thomas",
    email: "kate@example.com",
    status: "Active",
    joinDate: "2024-02-05",
  },
  {
    id: 12,
    name: "Leo Martinez",
    email: "leo@example.com",
    status: "Active",
    joinDate: "2024-01-20",
  },
];

const tableColumns: Column[] = [
  { key: "id", label: "ID", sortable: true, width: "60px", align: "center" },
  { key: "name", label: "Name", sortable: true },
  { key: "email", label: "Email", sortable: true },
  { key: "status", label: "Status", sortable: true, align: "center" },
  { key: "joinDate", label: "Join Date", sortable: true, align: "center" },
];

const handleSubmit = () => {
  console.log({
    name: name.value,
    email: email.value,
    agreed: agreed.value,
    role: role.value,
    selectedColor: selectedColor.value,
  });
};

const handleCommandSelect = (itemId: string) => {
  console.log("Command selected:", itemId);
  commandOpen.value = false;
};
</script>

<template>
  <div class="showcase">
    <div class="container">
      <h1>Component Showcase - All 21 Components</h1>

      <!-- Buttons Section -->
      <section class="section">
        <h2>Button Component</h2>
        <div class="grid">
          <div class="item">
            <h3>Variants</h3>
            <div class="flex-gap">
              <Button variant="primary">Primary</Button>
              <Button variant="secondary">Secondary</Button>
              <Button variant="outline">Outline</Button>
              <Button variant="ghost">Ghost</Button>
              <Button variant="destructive">Destructive</Button>
            </div>
          </div>
          <div class="item">
            <h3>Sizes</h3>
            <div class="flex-gap">
              <Button size="sm" variant="primary">Small</Button>
              <Button size="md" variant="primary">Medium</Button>
              <Button size="lg" variant="primary">Large</Button>
              <Button size="xl" variant="primary">XL</Button>
            </div>
          </div>
          <div class="item">
            <h3>States</h3>
            <div class="flex-gap">
              <Button variant="primary">Normal</Button>
              <Button variant="primary" disabled>Disabled</Button>
              <Button variant="primary" loading>Loading</Button>
              <Button variant="primary" icon-only>→</Button>
            </div>
          </div>
        </div>
      </section>

      <!-- Input Section -->
      <section class="section">
        <h2>Input Component</h2>
        <div class="grid">
          <div class="item">
            <h3>Basic Input</h3>
            <div class="form-group">
              <Label for="name" required>Name</Label>
              <Input
                id="name"
                v-model="name"
                placeholder="Enter your name"
                type="text"
              />
            </div>
          </div>
          <div class="item">
            <h3>Email Input</h3>
            <div class="form-group">
              <Label for="email">Email</Label>
              <Input
                id="email"
                v-model="email"
                placeholder="your@email.com"
                type="email"
              />
            </div>
          </div>
          <div class="item">
            <h3>Input with Error</h3>
            <div class="form-group">
              <Label for="error">Error State</Label>
              <Input
                id="error"
                placeholder="This has an error"
                type="text"
                error
                error-message="This field is required"
              />
            </div>
          </div>
        </div>
      </section>

      <!-- Checkbox Section -->
      <section class="section">
        <h2>Checkbox Component</h2>
        <div class="form-group">
          <div class="checkbox-group">
            <div class="checkbox-item">
              <Checkbox v-model="agreed" id="agree" />
              <Label for="agree">I agree to the terms</Label>
            </div>
            <div class="checkbox-item">
              <Checkbox id="newsletter" />
              <Label for="newsletter">Subscribe to newsletter</Label>
            </div>
            <div class="checkbox-item">
              <Checkbox id="notifications" disabled />
              <Label for="notifications" disabled
                >Notifications (disabled)</Label
              >
            </div>
          </div>
        </div>
      </section>

      <!-- RadioGroup Section -->
      <section class="section">
        <h2>Radio Component</h2>
        <div class="form-group">
          <Label>Choose a role:</Label>
          <RadioGroup v-model="role" :name="'role'" orientation="vertical">
            <div class="radio-item">
              <Radio value="admin" name="role" />
              <Label for="admin">Administrator</Label>
            </div>
            <div class="radio-item">
              <Radio value="user" name="role" />
              <Label for="user">User</Label>
            </div>
            <div class="radio-item">
              <Radio value="guest" name="role" disabled />
              <Label for="guest" disabled>Guest (disabled)</Label>
            </div>
          </RadioGroup>
          <p v-if="role" class="info">Selected role: {{ role }}</p>
        </div>
      </section>

      <!-- Select Section -->
      <section class="section">
        <h2>Select Component</h2>
        <div class="form-group">
          <Label for="color">Choose a color:</Label>
          <Select
            id="color"
            v-model="selectedColor"
            :options="colorOptions"
            placeholder="Pick a color..."
          />
          <p v-if="selectedColor" class="info">
            Selected color: {{ selectedColor }}
          </p>
        </div>
      </section>

      <!-- Form Section -->
      <section class="section">
        <h2>Complete Form Example</h2>
        <form @submit.prevent="handleSubmit" class="form">
          <div class="form-group">
            <Label for="form-name" required>Full Name</Label>
            <Input
              id="form-name"
              v-model="name"
              placeholder="John Doe"
              required
            />
          </div>

          <div class="form-group">
            <Label for="form-email" required>Email</Label>
            <Input
              id="form-email"
              v-model="email"
              type="email"
              placeholder="john@example.com"
              required
            />
          </div>

          <div class="form-group">
            <Label for="form-color">Favorite Color</Label>
            <Select
              id="form-color"
              v-model="selectedColor"
              :options="colorOptions"
            />
          </div>

          <div class="form-group">
            <div class="checkbox-item">
              <Checkbox v-model="agreed" id="form-agree" required />
              <Label for="form-agree" required>I agree to the terms</Label>
            </div>
          </div>

          <div class="form-actions">
            <Button type="submit" variant="primary" size="lg">Submit</Button>
            <Button type="button" variant="outline" size="lg">Cancel</Button>
          </div>
        </form>
      </section>

      <!-- Phase 2: Card Section -->
      <section class="section">
        <h2>Card Component</h2>
        <div class="grid">
          <Card>
            <template #header>
              <div class="card-header-content">
                <h4>Basic Card</h4>
              </div>
            </template>
            <p>This is a simple card with header, body, and footer.</p>
            <template #footer>
              <Button size="sm" variant="primary">Action</Button>
              <Button size="sm" variant="outline">Cancel</Button>
            </template>
          </Card>

          <Card hoverable>
            <template #header>
              <h4>Hoverable Card</h4>
            </template>
            <p>Hover over this card to see the effect.</p>
          </Card>

          <Card clickable>
            <template #header>
              <h4>Clickable Card</h4>
            </template>
            <p>Click this card for interaction.</p>
          </Card>
        </div>
      </section>

      <!-- Separator & Tabs Section -->
      <section class="section">
        <h2>Separator & Tabs Components</h2>

        <div class="separator-demo">
          <h3>Separator Examples</h3>
          <div>
            <p>Content above</p>
            <Separator />
            <p>Content below</p>
          </div>
        </div>

        <Separator />

        <div class="tabs-demo">
          <h3>Tabs Component</h3>
          <Tabs v-model="activeTab" :tabs="tabs">
            <template #content-tab1>
              <div class="tab-content">
                <p>This is the overview tab content.</p>
                <p>You can put any content here.</p>
              </div>
            </template>
            <template #content-tab2>
              <div class="tab-content">
                <p>This is the details tab content.</p>
                <p>More information goes here.</p>
              </div>
            </template>
            <template #content-tab3>
              <div class="tab-content">
                <p>This is the settings tab content.</p>
                <p>Configuration options here.</p>
              </div>
            </template>
          </Tabs>
        </div>
      </section>

      <!-- Alert Section -->
      <section class="section">
        <h2>Alert Component</h2>
        <div class="alerts-grid">
          <Alert variant="default" title="Default Alert">
            This is a default alert message.
          </Alert>

          <Alert variant="success" title="Success Alert">
            Operation completed successfully!
          </Alert>

          <Alert variant="warning" title="Warning Alert">
            Please review this warning message carefully.
          </Alert>

          <Alert variant="destructive" title="Error Alert">
            An error occurred. Please try again.
          </Alert>

          <Alert variant="info" title="Info Alert">
            Here is some information for you.
          </Alert>

          <Alert
            v-if="!closeAlert"
            variant="success"
            title="Closeable Alert"
            closeable
            @close="closeAlert = true"
          >
            This alert can be dismissed.
          </Alert>
        </div>
      </section>

      <!-- Badge Section -->
      <section class="section">
        <h2>Badge Component</h2>
        <div class="badge-grid">
          <div class="badge-group">
            <h3>Variants</h3>
            <div class="flex-gap">
              <Badge variant="default">Default</Badge>
              <Badge variant="primary">Primary</Badge>
              <Badge variant="secondary">Secondary</Badge>
              <Badge variant="success">Success</Badge>
              <Badge variant="warning">Warning</Badge>
              <Badge variant="destructive">Destructive</Badge>
              <Badge variant="info">Info</Badge>
            </div>
          </div>

          <div class="badge-group">
            <h3>Sizes</h3>
            <div class="flex-gap">
              <Badge size="sm">Small</Badge>
              <Badge size="md">Medium</Badge>
              <Badge size="lg">Large</Badge>
            </div>
          </div>

          <div class="badge-group">
            <h3>Pill Style</h3>
            <div class="flex-gap">
              <Badge pill variant="primary">Pill</Badge>
              <Badge pill variant="success">Active</Badge>
              <Badge pill variant="destructive">Error</Badge>
            </div>
          </div>
        </div>
      </section>

      <!-- Progress Section -->
      <section class="section">
        <h2>Progress Component</h2>
        <div class="progress-grid">
          <div class="progress-item">
            <h3>Basic Progress</h3>
            <Progress :value="35" />
          </div>

          <div class="progress-item">
            <h3>With Label</h3>
            <Progress :value="progressValue" show-label />
            <div class="flex-gap" style="margin-top: 1rem">
              <Button
                size="sm"
                @click="progressValue = Math.max(0, progressValue - 10)"
              >
                -
              </Button>
              <Button
                size="sm"
                @click="progressValue = Math.min(100, progressValue + 10)"
              >
                +
              </Button>
            </div>
          </div>

          <div class="progress-item">
            <h3>Variants</h3>
            <div style="display: flex; flex-direction: column; gap: 1rem">
              <div>
                <p style="margin-bottom: 0.5rem; font-size: 0.875rem">
                  Default
                </p>
                <Progress :value="100" variant="default" />
              </div>
              <div>
                <p style="margin-bottom: 0.5rem; font-size: 0.875rem">
                  Success
                </p>
                <Progress :value="100" variant="success" />
              </div>
              <div>
                <p style="margin-bottom: 0.5rem; font-size: 0.875rem">
                  Warning
                </p>
                <Progress :value="100" variant="warning" />
              </div>
              <div>
                <p style="margin-bottom: 0.5rem; font-size: 0.875rem">
                  Destructive
                </p>
                <Progress :value="100" variant="destructive" />
              </div>
            </div>
          </div>

          <div class="progress-item">
            <h3>Striped & Animated</h3>
            <Progress :value="80" striped animated show-label />
          </div>

          <div class="progress-item">
            <h3>Sizes</h3>
            <div style="display: flex; flex-direction: column; gap: 1rem">
              <div>
                <p style="margin-bottom: 0.5rem; font-size: 0.875rem">Small</p>
                <Progress :value="60" size="sm" />
              </div>
              <div>
                <p style="margin-bottom: 0.5rem; font-size: 0.875rem">Medium</p>
                <Progress :value="60" size="md" />
              </div>
              <div>
                <p style="margin-bottom: 0.5rem; font-size: 0.875rem">Large</p>
                <Progress :value="60" size="lg" />
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Stage 4: Data Components -->
      <section class="section">
        <h2>Combobox Component</h2>
        <div class="flex-gap">
          <div class="form-group" style="min-width: 250px">
            <Label>Select Framework</Label>
            <Combobox
              v-model="selectedComboboxValue"
              :options="comboboxOptions"
              placeholder="Search frameworks..."
              searchable
              clearable
              size="md"
            />
            <p class="info">Selected: {{ selectedComboboxValue || "None" }}</p>
          </div>

          <Card>
            <template #header>
              <h4>Combobox Features</h4>
            </template>
            <ul>
              <li>Searchable dropdown</li>
              <li>Keyboard navigation (↑↓ keys)</li>
              <li>Clearable option</li>
              <li>Multiple size variants</li>
              <li>Disabled state support</li>
            </ul>
          </Card>
        </div>
      </section>

      <section class="section">
        <h2>Command Component</h2>
        <div class="flex-gap">
          <Button
            variant="primary"
            @click="commandOpen = true"
            style="width: fit-content"
          >
            Open Command Palette (Press Cmd+K)
          </Button>

          <Command
            :open="commandOpen"
            @update:open="commandOpen = $event"
            :items="commandItems"
            placeholder="Type a command..."
            @select="handleCommandSelect"
          />

          <Card>
            <template #header>
              <h4>Command Component Features</h4>
            </template>
            <ul>
              <li>Command palette interface</li>
              <li>Grouped commands by category</li>
              <li>Keyboard shortcuts display</li>
              <li>Icon support</li>
              <li>Full keyboard navigation</li>
            </ul>
          </Card>
        </div>
      </section>

      <section class="section">
        <h2>Calendar Component</h2>
        <div class="flex-gap">
          <Calendar
            v-model="selectedDate"
            :min-date="new Date(2024, 0, 1)"
            :max-date="new Date(2025, 11, 31)"
          />

          <Card>
            <template #header>
              <h4>Calendar Features</h4>
            </template>
            <div class="calendar-info">
              <p>
                <strong>Selected Date:</strong>
                {{ selectedDate ? selectedDate.toLocaleDateString() : "None" }}
              </p>
              <ul>
                <li>Month/Year navigation</li>
                <li>Date range restrictions</li>
                <li>Today highlight</li>
                <li>Full keyboard support</li>
                <li>Click to select date</li>
              </ul>
            </div>
          </Card>
        </div>
      </section>

      <section class="section">
        <h2>DataTable Component</h2>
        <DataTable
          :data="tableData"
          :columns="tableColumns"
          :page-size="5"
          striped
          hoverable
          bordered
          size="md"
        />
      </section>

      <!-- Phase 3: Dialog Section -->
      <section class="section">
        <h2>Dialog Component</h2>
        <div class="flex-gap">
          <Dialog
            v-model:open="dialogOpen"
            title="Sample Dialog"
            description="This is a dialog example"
          >
            <template #trigger>
              <Button variant="primary">Open Dialog</Button>
            </template>

            <p>This is the dialog content. You can put any content here.</p>
            <p>The dialog includes a title, description, and action buttons.</p>

            <template #footer>
              <Button size="sm" variant="outline" @click="dialogOpen = false"
                >Cancel</Button
              >
              <Button size="sm" variant="primary" @click="dialogOpen = false"
                >Confirm</Button
              >
            </template>
          </Dialog>

          <Dialog size="lg" title="Large Dialog">
            <template #trigger>
              <Button variant="secondary">Large Dialog</Button>
            </template>

            <p>This is a large dialog with more space for content.</p>
          </Dialog>
        </div>
      </section>

      <!-- Popover Section -->
      <section class="section">
        <h2>Popover Component</h2>
        <div class="flex-gap">
          <Popover v-model:open="popoverOpen">
            <template #trigger>
              <Button variant="outline">Open Popover</Button>
            </template>

            <h4 style="margin: 0 0 0.5rem 0; font-size: 1rem">Popover Title</h4>
            <p style="margin: 0">
              This is popover content. It can be positioned around the trigger
              element.
            </p>
          </Popover>

          <Popover side="top">
            <template #trigger>
              <Button variant="outline">Top Popover</Button>
            </template>

            <p style="margin: 0">Positioned above the trigger</p>
          </Popover>

          <Popover side="left">
            <template #trigger>
              <Button variant="outline">Left Popover</Button>
            </template>

            <p style="margin: 0">Positioned to the left</p>
          </Popover>
        </div>
      </section>

      <!-- Dropdown Section -->
      <section class="section">
        <h2>Dropdown Component</h2>
        <div class="flex-gap">
          <Dropdown
            :items="dropdownItems"
            @select="(id) => console.log('Selected:', id)"
          >
            <template #trigger>
              <Button variant="outline">Menu ▼</Button>
            </template>
          </Dropdown>

          <Dropdown
            :items="[
              { id: 'new', label: 'New' },
              { id: 'open', label: 'Open' },
              { id: 'separator1', separator: true },
              { id: 'save', label: 'Save' },
              { id: 'export', label: 'Export' },
            ]"
            @select="(id) => console.log('Selected:', id)"
          >
            <template #trigger>
              <Button variant="secondary">File ▼</Button>
            </template>
          </Dropdown>
        </div>
      </section>

      <!-- Sheet Section -->
      <section class="section">
        <h2>Sheet Component</h2>
        <div class="flex-gap">
          <Sheet
            v-model:open="sheetOpen"
            title="Sidebar Sheet"
            :side="sheetSide"
          >
            <template #trigger>
              <Button variant="primary">Open Sheet</Button>
            </template>

            <p>
              This is a sheet/drawer component that slides in from the edge.
            </p>
            <p>
              You can put any content here, including forms, navigation, etc.
            </p>

            <div
              style="
                margin-top: 1rem;
                display: flex;
                flex-direction: column;
                gap: 0.5rem;
              "
            >
              <Label>Position:</Label>
              <div style="display: flex; gap: 0.5rem">
                <Button
                  v-for="side in ['top', 'right', 'bottom', 'left'] as const"
                  :key="side"
                  size="sm"
                  :variant="sheetSide === side ? 'primary' : 'outline'"
                  @click="sheetSide = side"
                >
                  {{ side.charAt(0).toUpperCase() + side.slice(1) }}
                </Button>
              </div>
            </div>

            <template #footer>
              <Button size="sm" variant="outline" @click="sheetOpen = false"
                >Close</Button
              >
              <Button size="sm" variant="primary" @click="sheetOpen = false"
                >Done</Button
              >
            </template>
          </Sheet>

          <Sheet title="Quick Navigation" side="left">
            <template #trigger>
              <Button variant="outline">Navigation</Button>
            </template>

            <nav style="display: flex; flex-direction: column; gap: 0.5rem">
              <Button variant="ghost" style="justify-content: flex-start"
                >Home</Button
              >
              <Button variant="ghost" style="justify-content: flex-start"
                >About</Button
              >
              <Button variant="ghost" style="justify-content: flex-start"
                >Services</Button
              >
              <Button variant="ghost" style="justify-content: flex-start"
                >Contact</Button
              >
            </nav>
          </Sheet>
        </div>
      </section>
    </div>
  </div>
</template>

<style scoped lang="scss">
.showcase {
  padding: 2rem 1rem;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
}

h1 {
  font-size: 2rem;
  font-weight: 700;
  margin-bottom: 2rem;
  color: var(--foreground);
}

h2 {
  font-size: 1.5rem;
  font-weight: 600;
  margin-bottom: 1rem;
  color: var(--foreground-secondary);
}

h3 {
  font-size: 1rem;
  font-weight: 600;
  margin-bottom: 0.75rem;
  color: var(--foreground-secondary);
}

.section {
  margin-bottom: 3rem;
  padding: 2rem;
  background-color: var(--card);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
}

.grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.item {
  padding: 1.5rem;
  background-color: var(--background);
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
}

.flex-gap {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
  align-items: center;
}

.form-group {
  margin-bottom: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.checkbox-group,
.radio-item {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.checkbox-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.radio-item {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 0.5rem;
}

.info {
  font-size: 0.875rem;
  color: var(--foreground-secondary);
  margin-top: 0.5rem;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
  max-width: 500px;
}

.form-actions {
  display: flex;
  gap: 1rem;
  margin-top: 1rem;
}

.card-header-content h4 {
  margin: 0;
  font-size: 1rem;
}

.separator-demo {
  margin-bottom: 2rem;
}

.tabs-demo {
  margin-top: 1.5rem;
}

.tab-content {
  padding: 1rem;
  background-color: var(--background);
  border-radius: var(--radius-md);
  border: 1px solid var(--border);
}

.alerts-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1rem;
}

.badge-grid {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.badge-group {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.progress-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
}

.progress-item {
  padding: 1.5rem;
  background-color: var(--background);
  border: 1px solid var(--border);
  border-radius: var(--radius-md);
}

.progress-item h3 {
  margin-top: 0;
}

/* Phase 3 Styles */
.dialog-demo,
.popover-demo,
.dropdown-demo,
.sheet-demo {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.overlay-note {
  padding: 1rem;
  background-color: var(--background-muted);
  border-radius: var(--radius-md);
  font-size: 0.875rem;
  color: var(--foreground-secondary);
}

/* Stage 4 Styles */
.calendar-info {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.calendar-info p {
  margin: 0;
  font-size: 0.875rem;
}

.calendar-info ul {
  margin: 0.5rem 0;
  padding-left: 1.25rem;
  font-size: 0.875rem;
}

.calendar-info li {
  margin: 0.25rem 0;
}
</style>
