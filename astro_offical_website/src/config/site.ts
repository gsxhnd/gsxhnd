/**
 * 网站配置文件
 * 集中管理网站的固定信息和配置项
 */

import { useTranslations, type Language } from "../i18n/utils";

export interface NavigationItem {
  label: string;
  href: string;
  submenu?: NavigationSubItem[];
}

export interface NavigationSubItem {
  label: string;
  href: string;
}

export const SITE_CONFIG = {
  // 基本信息
  name: "Nagi Inc.",
  url: "https://stargazers.club",

  // SEO
  themeColor: "#222828",

  // 社交媒体链接
  social: {
    github: "https://github.com/",
    twitter: "https://twitter.com/",
    linkedin: "https://www.linkedin.com/",
    youtube: "https://youtube.com/",
    instagram: "https://instagram.com/",
  },

  // 布局配置
  layout: {
    headerHeight: "74px",
    headerHeightMobile: "60px",
  },

  // 分页配置
  pagination: {
    newsPageSize: 10,
  },
} as const;

/** Locale-aware navigation items */
export function getNavigation(lang: Language): NavigationItem[] {
  const t = useTranslations(lang);
  const prefix = `/${lang}`;
  return [
    { label: t("nav.about"), href: `${prefix}/about` },
    { label: t("nav.company"), href: `${prefix}/company` },
    {
      label: t("nav.services"),
      href: `${prefix}/services`,
      submenu: lang === "zh"
        ? [
          { label: "Web网站·应用开发", href: `${prefix}/services#web-development` },
          { label: "移动应用开发", href: `${prefix}/services#mobile-apps` },
          { label: "UI/UX设计", href: `${prefix}/services#uiux-design` },
          { label: "数字策略·咨询", href: `${prefix}/services#digital-strategy` },
        ]
        : [
          { label: "Website & App Development", href: `${prefix}/services#web-development` },
          { label: "Mobile App Development", href: `${prefix}/services#mobile-apps` },
          { label: "UI/UX Design", href: `${prefix}/services#uiux-design` },
          { label: "Digital Strategy & Consulting", href: `${prefix}/services#digital-strategy` },
        ],
    },
    { label: t("nav.works"), href: `${prefix}/works` },
    { label: t("nav.news"), href: `${prefix}/news` },
  ];
}

/** Locale-aware footer navigation */
export function getFooterNavigation(lang: Language) {
  const t = useTranslations(lang);
  const prefix = `/${lang}`;
  return {
    services: [
      { label: t("nav.services"), href: `${prefix}/services` },
      { label: t("nav.works"), href: `${prefix}/works` },
    ],
    resources: [
      { label: t("nav.news"), href: `${prefix}/news` },
    ],
    company: [
      { label: t("nav.about"), href: `${prefix}/about` },
      { label: t("nav.company"), href: `${prefix}/company` },
      { label: t("nav.contact"), href: `${prefix}/contact` },
      { label: lang === "zh" ? "个人信息保护方针" : "Privacy Policy", href: `${prefix}/privacy` },
    ],
  };
}

export type SiteConfig = typeof SITE_CONFIG;
