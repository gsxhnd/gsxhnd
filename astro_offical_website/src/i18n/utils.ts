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

export function getLangFromUrl(url: URL): Language {
  const [, lang] = url.pathname.split("/");
  if (lang in ui) return lang as Language;
  return defaultLang;
}

export function useTranslations(lang: Language) {
  return function t(key: string) {
    const keys = key.split(".");
    let value: any = ui[lang];
    
    for (const k of keys) {
      value = value?.[k];
    }
    
    return value || key;
  };
}

export function getLocalizedPath(path: string, lang: Language): string {
  if (lang === defaultLang) {
    return path;
  }
  return `/${lang}${path}`;
}

export function removeLocalePath(path: string): string {
  const [, maybeLang, ...rest] = path.split("/");
  if (maybeLang in ui) {
    return `/${rest.join("/")}`;
  }
  return path;
}
