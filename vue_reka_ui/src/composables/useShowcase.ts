import { ref } from 'vue';
import { useI18n } from 'vue-i18n';

export interface ComponentItem {
  id: string;
  label: string;
}

export interface ComponentCategory {
  id: string;
  label: string;
  components: ComponentItem[];
}

export function useShowcase() {
  const { t } = useI18n();

  const selectedCategory = ref('form');
  const selectedComponent = ref('button');
  const expandedCategories = ref(new Set(['form']));

  const componentCategories: ComponentCategory[] = [
    {
      id: 'form',
      label: t('showcase.sections.form'),
      components: [
        { id: 'button', label: t('showcase.components.button') },
        { id: 'input', label: t('showcase.components.input') },
        { id: 'label', label: t('showcase.components.label') },
        { id: 'checkbox', label: t('showcase.components.checkbox') },
        { id: 'radio', label: t('showcase.components.radio') },
        { id: 'select', label: t('showcase.components.select') },
        { id: 'combobox', label: t('showcase.components.combobox') },
      ],
    },
    {
      id: 'feedback',
      label: t('showcase.sections.feedback'),
      components: [
        { id: 'badge', label: t('showcase.components.badge') },
        { id: 'alert', label: t('showcase.components.alert') },
        { id: 'progress', label: t('showcase.components.progress') },
        { id: 'tabs', label: t('showcase.components.tabs') },
      ],
    },
    {
      id: 'layout',
      label: t('showcase.sections.layout'),
      components: [
        { id: 'card', label: t('showcase.components.card') },
        { id: 'separator', label: t('showcase.components.separator') },
      ],
    },
    {
      id: 'overlay',
      label: t('showcase.sections.overlay'),
      components: [
        { id: 'dialog', label: t('showcase.components.dialog') },
        { id: 'popover', label: t('showcase.components.popover') },
        { id: 'dropdown', label: t('showcase.components.dropdown') },
        { id: 'sheet', label: t('showcase.components.sheet') },
      ],
    },
    {
      id: 'advanced',
      label: t('showcase.sections.advanced'),
      components: [
        { id: 'command', label: t('showcase.components.command') },
        { id: 'calendar', label: t('showcase.components.calendar') },
        { id: 'dataTable', label: t('showcase.components.dataTable') },
      ],
    },
  ];

  const selectComponent = (categoryId: string, componentId: string) => {
    selectedCategory.value = categoryId;
    selectedComponent.value = componentId;
  };

  const toggleCategory = (categoryId: string) => {
    if (expandedCategories.value.has(categoryId)) {
      expandedCategories.value.delete(categoryId);
    } else {
      expandedCategories.value.add(categoryId);
    }
  };

  return {
    selectedCategory,
    selectedComponent,
    expandedCategories,
    componentCategories,
    selectComponent,
    toggleCategory,
  };
}
