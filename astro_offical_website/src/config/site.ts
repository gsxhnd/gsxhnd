/**
 * 网站配置文件
 * 集中管理网站的固定信息和配置项
 */

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
  title: "Nagi Inc. | 体験をデザインし、未来を創造する",
  description: "デザインとテクノロジーの力で、ブランドの本質を引き出し、新しい価値を生み出します",
  url: "https://stargazers.club",
  lang: "zh",

  // SEO
  themeColor: "#222828",

  // 公司信息
  company: {
    name: "Nagi Inc.",
    nameJa: "株式会社Nagi",
    established: "2015年4月1日",
    address: "〒150-0001\n東京都渋谷区神宮前4-25-4 原宿ビル3F",
    phone: "+81-3-XXXX-XXXX",
    phoneDisplay: "+81-3-XXXX-XXXX",
    email: "info@nagi.co.jp",
    emailDisplay: "info@nagi.co.jp",
    businessHours: "平日 9:00 - 18:00",
    businessHoursNote: "(土日祝日は休業)",
    businessHoursFull: "営業時間: 平日 9:00 - 18:00 (土日祝日は休業)",
    ceo: "桐島 紗都",
    ceoTitle: "代表取締役社長",
    employees: "28名（2025年1月現在）",
    capital: "3,000万円",
  },

  // 社交媒体链接
  social: {
    github: "https://github.com/",
    twitter: "https://twitter.com/",
    linkedin: "https://www.linkedin.com/",
    youtube: "https://youtube.com/",
    instagram: "https://instagram.com/",
  },

  // 导航菜单
  navigation: [
    {
      label: "私たちについて",
      href: "./about",
    },
    {
      label: "会社概要",
      href: "./company",
    },
    {
      label: "サービス",
      href: "./services",
      submenu: [
        {
          label: "Webサイト・アプリケーション開発",
          href: "./services#web-development",
        },
        {
          label: "モバイルアプリ開発",
          href: "./services#mobile-apps",
        },
        {
          label: "UI/UXデザイン",
          href: "./services#uiux-design",
        },
        {
          label: "デジタル戦略・コンサルティング",
          href: "./services#digital-strategy",
        },
      ],
    },
    {
      label: "制作実績",
      href: "./works",
    },
    {
      label: "お知らせ",
      href: "./news",
    },
  ] as NavigationItem[],

  // Footer 导航
  footerNavigation: {
    services: [
      { label: "サービス", href: "./services" },
      { label: "制作実績", href: "./works" },
    ],
    resources: [
      { label: "お知らせ", href: "./news" },
      { label: "よくあるご質問", href: "./services#faq" },
    ],
    company: [
      { label: "私たちについて", href: "./about" },
      { label: "会社概要", href: "./company" },
      { label: "お問い合わせ", href: "./contact" },
      { label: "個人情報保護方針", href: "./privacy" },
    ],
  },

  // 业务内容
  businessContent: [
    "デジタルプロダクトのデザイン・開発",
    "ブランディング戦略の企画・実行",
    "UI/UXコンサルティング",
    "Webサイト・アプリケーションの企画・制作",
    "デジタルマーケティング支援",
  ],

  // 资格认证
  certifications: [
    "情報セキュリティマネジメントシステム認証 (ISO 27001)",
    "Google Partner",
    "AWS Partner",
  ],

  // 公司沿革
  history: [
    { year: "2015", month: "4月", description: "東京都渋谷区にて株式会社Nagi設立" },
    { year: "2016", month: "3月", description: "大手企業向けブランディングサービス開始" },
    { year: "2017", month: "8月", description: "従業員数10名を突破、オフィス拡張移転" },
    { year: "2018", month: "5月", description: "UI/UXコンサルティング事業を本格展開" },
    { year: "2019", month: "10月", description: "グッドデザイン賞受賞（コーポレートサイト制作実績）" },
    { year: "2020", month: "6月", description: "リモートワーク体制の全社導入、働き方改革推進" },
    { year: "2021", month: "9月", description: "デジタルマーケティング支援サービス開始" },
    { year: "2022", month: "11月", description: "従業員数20名を突破、パートナー企業との協業体制強化" },
    { year: "2023", month: "3月", description: "AI技術を活用したデザインツール開発プロジェクト始動" },
    { year: "2024", month: "4月", description: "創立10周年、ブランドリニューアル実施" },
    { year: "2025", month: "1月", description: "従業員数28名、新規事業領域への展開を加速" },
  ],

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

export type SiteConfig = typeof SITE_CONFIG;
export type { NavigationItem, NavigationSubItem };
