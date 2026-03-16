import { ref } from 'vue';

export function useTheme() {
  const currentTheme = ref((window as any).__theme?.current?.value || 'dark');

  const toggleTheme = () => {
    const newTheme = currentTheme.value === 'dark' ? 'light' : 'dark';
    currentTheme.value = newTheme;
    (window as any).__theme?.set?.(newTheme);
  };

  return {
    currentTheme,
    toggleTheme,
  };
}
