<script setup lang="ts">
import { useI18n } from 'vue-i18n';
import { useTheme } from '@/composables/useTheme';

const { t, locale } = useI18n();
const { currentTheme, toggleTheme } = useTheme();

const switchLanguage = (lang: string) => {
  locale.value = lang;
};
</script>

<template>
  <header class="showcase-header">
    <div class="header-content">
      <h1>{{ t('showcase.title') }}</h1>
      <div class="header-controls">
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

        <button class="theme-btn" @click="toggleTheme">
          {{ currentTheme === 'dark' ? '☀️' : '🌙' }}
          {{ currentTheme === 'dark' ? t('showcase.lightMode') : t('showcase.darkMode') }}
        </button>
      </div>
    </div>
  </header>
</template>

<style scoped lang="scss">
.showcase-header {
  background-color: var(--background-secondary);
  border-bottom: 1px solid var(--background-tertiary);
  padding: 1.5rem 2rem;
  position: sticky;
  top: 0;
  z-index: 100;

  .header-content {
    max-width: 1600px;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
    gap: 2rem;

    h1 {
      margin: 0;
      font-size: 1.875rem;
      font-weight: 700;
      color: var(--foreground);
    }

    .header-controls {
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

@media (max-width: 768px) {
  .showcase-header {
    padding: 1rem;

    .header-content {
      flex-direction: column;
      align-items: flex-start;
      gap: 1rem;

      h1 {
        font-size: 1.5rem;
      }

      .header-controls {
        width: 100%;
        justify-content: flex-start;
      }
    }
  }
}

@media (max-width: 480px) {
  .showcase-header {
    .header-content {
      h1 {
        font-size: 1.25rem;
      }

      .header-controls {
        flex-direction: column;
        gap: 0.75rem;

        .control-group {
          width: 100%;
          justify-content: center;
        }

        .theme-btn {
          width: 100%;
        }
      }
    }
  }
}
</style>
