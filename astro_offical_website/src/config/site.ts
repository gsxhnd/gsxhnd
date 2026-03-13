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
  name: "CName 公司",
  title: "CName 公司 | 设计体验，创造未来",
  description: "通过设计和技术的力量，挖掘品牌本质，创造新的价值",
  url: "https://stargazers.club",
  lang: "zh",

  // SEO
  themeColor: "#222828",

  // 公司信息
  company: {
    name: "CName Inc.",
    nameZh: "CName 公司",
    established: "2019年4月1日",
    address: "123123\n123123",
    phone: "+81-3-XXXX-XXXX",
    phoneDisplay: "+81-3-XXXX-XXXX",
    email: "xxx@xxx",
    emailDisplay: "xxx@xxx",
    businessHours: "工作日 9:00 - 18:00",
    businessHoursNote: "(周末及节假日休息)",
    businessHoursFull: "周一至周五 9:00 - 18:00（周末和假日休息）",
    ceo: "123",
    employees: "35名（截至2024年12月末）",
    capital: "1万",
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
      label: "关于我们",
      href: "./about",
    },
    {
      label: "公司概况",
      href: "./company",
    },
    {
      label: "服务",
      href: "./services",
      submenu: [
        {
          label: "网站与应用开发",
          href: "./services#web-development",
        },
        {
          label: "移动应用开发",
          href: "./services#mobile-apps",
        },
        {
          label: "UI/UX 设计",
          href: "./services#uiux-design",
        },
        {
          label: "数字战略",
          href: "./services#digital-strategy",
        },
      ],
    },
    {
      label: "作品案例",
      href: "./works",
    },
    {
      label: "新闻动态",
      href: "./news",
    },
  ] as NavigationItem[],
    {
      label: "关于我们",
      href: "./about",
    },
    {
      label: "公司概况",
      href: "./company",
    },
    {
      label: "服务",
      href: "./services",
      submenu: [
        {
          label: "网站与应用开发",
          href: "./services#web-development",
        },
        {
          label: "移动应用开发",
          href: "./services#mobile-apps",
        },
        {
          label: "UI/UX设计",
          href: "./services#uiux-design",
        },
        {
          label: "数字战略",
          href: "./services#digital-strategy",
        },
      ],
    },
    {
      label: "作品案例",
      href: "./works",
    },
    {
      label: "新闻动态",
      href: "./news",
    },
  ],

  // Footer 导航
  footerNavigation: {
    services: [
      { label: "服务", href: "./services" },
      { label: "作品案例", href: "./works" },
    ],
    resources: [
      { label: "新闻动态", href: "./news" },
      { label: "关于我们", href: "./about" },
    ],
    company: [
      { label: "公司概况", href: "./company" },
      { label: "联系我们", href: "./contact" },
    ],
    legal: [{ label: "隐私政策", href: "./privacy" }],
  },

  // 业务内容
  businessContent: [
    "Web网站・Web应用开发",
    "移动应用开发",
    "UI/UX设计",
    "品牌咨询",
    "数字营销支持",
  ],

  // 资格认证
  certifications: [
    "信息安全管理系统认证 (ISO 27001)",
    "Google Partner",
    "AWS Partner",
  ],

  // 布局配置
  layout: {
    headerHeight: "74px",
    headerHeightMobile: "60px",
  },

  // 分页配置
  pagination: {
    newsPageSize: 2,
  },
} as const;

export type SiteConfig = typeof SITE_CONFIG;
export type { NavigationItem, NavigationSubItem };
