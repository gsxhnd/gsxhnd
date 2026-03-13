<script setup lang="ts">
import { ref, computed } from "vue";

type DateValue = Date | null;

interface Props {
  modelValue?: DateValue;
  disabled?: boolean;
  minDate?: Date;
  maxDate?: Date;
}

interface Emits {
  (e: "update:modelValue", value: DateValue): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const currentDate = ref(new Date());
const selected = computed({
  get: () => props.modelValue,
  set: (value: DateValue) => emit("update:modelValue", value),
});

const currentMonth = computed(() => currentDate.value.getMonth());
const currentYear = computed(() => currentDate.value.getFullYear());

const monthName = computed(() => {
  const months = [
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
  ];
  return months[currentMonth.value];
});

const daysInMonth = computed(() => {
  return new Date(currentYear.value, currentMonth.value + 1, 0).getDate();
});

const firstDayOfMonth = computed(() => {
  return new Date(currentYear.value, currentMonth.value, 1).getDay();
});

const days = computed(() => {
  const result: (number | null)[] = [];

  // Add empty cells for days before month starts
  for (let i = 0; i < firstDayOfMonth.value; i++) {
    result.push(null);
  }

  // Add days of month
  for (let i = 1; i <= daysInMonth.value; i++) {
    result.push(i);
  }

  return result;
});

const isDateDisabled = (day: number | null) => {
  if (props.disabled) return true;
  if (!day) return true;

  const date = new Date(currentYear.value, currentMonth.value, day);

  if (props.minDate && date < props.minDate) return true;
  if (props.maxDate && date > props.maxDate) return true;

  return false;
};

const isDateSelected = (day: number | null) => {
  if (!selected.value || !day) return false;

  return (
    day === selected.value.getDate() &&
    currentMonth.value === selected.value.getMonth() &&
    currentYear.value === selected.value.getFullYear()
  );
};

const isDateToday = (day: number | null) => {
  if (!day) return false;

  const today = new Date();
  return (
    day === today.getDate() &&
    currentMonth.value === today.getMonth() &&
    currentYear.value === today.getFullYear()
  );
};

const selectDate = (day: number | null) => {
  if (!(day && day > 0)) return;
  if (isDateDisabled(day)) return;
  selected.value = new Date(currentYear.value, currentMonth.value, day);
};

const previousMonth = () => {
  currentDate.value = new Date(currentYear.value, currentMonth.value - 1, 1);
};

const nextMonth = () => {
  currentDate.value = new Date(currentYear.value, currentMonth.value + 1, 1);
};

const previousYear = () => {
  currentDate.value = new Date(currentYear.value - 1, currentMonth.value, 1);
};

const nextYear = () => {
  currentDate.value = new Date(currentYear.value + 1, currentMonth.value, 1);
};
</script>

<template>
  <div class="calendar">
    <div class="calendar__header">
      <div class="calendar__nav">
        <button
          class="calendar__button calendar__button--nav"
          @click="previousYear"
          :disabled="disabled"
        >
          ‹‹
        </button>
        <button
          class="calendar__button calendar__button--nav"
          @click="previousMonth"
          :disabled="disabled"
        >
          ‹
        </button>
      </div>

      <div class="calendar__title">
        <span>{{ monthName }} {{ currentYear }}</span>
      </div>

      <div class="calendar__nav">
        <button
          class="calendar__button calendar__button--nav"
          @click="nextMonth"
          :disabled="disabled"
        >
          ›
        </button>
        <button
          class="calendar__button calendar__button--nav"
          @click="nextYear"
          :disabled="disabled"
        >
          ››
        </button>
      </div>
    </div>

    <div class="calendar__weekdays">
      <div
        v-for="day in ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']"
        :key="day"
        class="calendar__weekday"
      >
        {{ day }}
      </div>
    </div>

    <div class="calendar__grid">
      <button
        v-for="(day, index) in days"
        :key="index"
        :class="[
          'calendar__day',
          {
            'calendar__day--empty': !day,
            'calendar__day--selected': isDateSelected(day),
            'calendar__day--today': isDateToday(day),
            'calendar__day--disabled': isDateDisabled(day),
          },
        ]"
        :disabled="isDateDisabled(day)"
        @click="selectDate(day)"
      >
        {{ day }}
      </button>
    </div>
  </div>
</template>

<style scoped lang="scss">
.calendar {
  background-color: var(--popover);
  border: 1px solid var(--border);
  border-radius: var(--radius-lg);
  padding: 1rem;
  width: fit-content;
  box-shadow: var(--shadow-md);
}

.calendar__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 1.5rem;
  gap: 1rem;
}

.calendar__nav {
  display: flex;
  gap: 0.5rem;
}

.calendar__button {
  border: none;
  background-color: transparent;
  color: var(--popover-foreground);
  cursor: pointer;
  padding: 0.5rem;
  border-radius: var(--radius-md);
  transition: background-color 0.15s ease;
  outline: none;
  font-weight: 600;

  &:hover:not(:disabled) {
    background-color: var(--background-muted);
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

  &--nav {
    padding: 0.5rem 0.75rem;
    font-size: 0.875rem;
  }
}

.calendar__title {
  font-size: 1rem;
  font-weight: 600;
  color: var(--popover-foreground);
  min-width: 160px;
  text-align: center;
}

.calendar__weekdays {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 0.5rem;
  margin-bottom: 0.5rem;
}

.calendar__weekday {
  width: 32px;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 0.75rem;
  font-weight: 600;
  color: var(--foreground-muted);
}

.calendar__grid {
  display: grid;
  grid-template-columns: repeat(7, 1fr);
  gap: 0.5rem;
}

.calendar__day {
  width: 32px;
  height: 32px;
  border: 1px solid transparent;
  border-radius: var(--radius-md);
  background-color: transparent;
  color: var(--popover-foreground);
  cursor: pointer;
  outline: none;
  transition:
    background-color 0.15s ease,
    border-color 0.15s ease,
    color 0.15s ease;
  font-size: 0.875rem;
  font-weight: 500;

  &:hover:not(.calendar__day--empty):not(.calendar__day--disabled) {
    background-color: var(--background-muted);
  }

  &:focus-visible:not(.calendar__day--empty):not(.calendar__day--disabled) {
    outline: 2px solid var(--primary);
    outline-offset: 2px;
  }

  &--empty {
    cursor: default;
  }

  &--selected {
    background-color: var(--primary);
    color: var(--primary-foreground);
    border-color: var(--primary);
    font-weight: 600;
  }

  &--today {
    border-color: var(--primary);
    color: var(--primary);
  }

  &--disabled {
    color: var(--foreground-disabled);
    cursor: not-allowed;
    opacity: 0.5;
  }
}
</style>
