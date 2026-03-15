import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";

const routes: RouteRecordRaw[] = [
  {
    path: "/",
    redirect: "/showcase",
  },
  {
    path: "/showcase",
    component: () => import("@/pages/Showcase.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
