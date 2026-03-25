import zh from "./locales/zh.json";
import en from "./locales/en.json";

export const languages = {
  zh: "中文",
  en: "English",
};

export const defaultLang = "zh";

export const ui = {
  zh,
  en,
} as const;

export type Language = keyof typeof ui;

/** Locale to date formatting locale mapping */
export const dateLocales: Record<Language, string> = {
  zh: "zh-CN",
  en: "en-US",
};

export function getLangFromUrl(url: URL): Language {
  const [, lang] = url.pathname.split("/");
  if (lang in ui) return lang as Language;
  return defaultLang;
}

export function useTranslations(lang: Language) {
  return function t(key: string): string {
    const keys = key.split(".");
    let value: any = ui[lang];

    for (const k of keys) {
      value = value?.[k];
    }

    return value || key;
  };
}

/**
 * Always prefix the locale, since prefixDefaultLocale is true.
 */
export function getLocalizedPath(path: string, lang: Language): string {
  return `/${lang}${path}`;
}

export function removeLocalePath(path: string): string {
  const [, maybeLang, ...rest] = path.split("/");
  if (maybeLang in ui) {
    const basePath = rest.join("/");
    return basePath ? `/${basePath}` : "/";
  }
  return path;
}
