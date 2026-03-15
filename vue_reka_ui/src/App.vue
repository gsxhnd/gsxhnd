<script setup lang="ts">
import { ref, onBeforeMount } from "vue";

type Theme = "dark" | "light";

const currentTheme = ref<Theme>("dark");

const setTheme = (theme: Theme) => {
  currentTheme.value = theme;
  const doc = document.documentElement;
  doc.dataset.theme = "neutral";
  doc.dataset.colorSchema = theme;
  localStorage.setItem("theme", theme);
};

onBeforeMount(() => {
  const savedTheme = localStorage.getItem("theme") as Theme | null;
  const theme = savedTheme || "dark";
  setTheme(theme);
});

// 提供全局主题切换函数
const toggleTheme = () => {
  setTheme(currentTheme.value === "dark" ? "light" : "dark");
};

// 暴露给全局
(window as any).__theme = {
  current: currentTheme,
  set: setTheme,
  toggle: toggleTheme,
};
</script>

<template>
  <router-view />
</template>
