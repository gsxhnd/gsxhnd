/**
 * Apply saved theme immediately on page load to prevent FOUC (Flash of Unstyled Content)
 * This script must run before React hydration
 */
(function () {
  const savedTheme = localStorage.getItem("theme");
  const prefersDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
  const theme = savedTheme || (prefersDark ? "dark" : "light");

  document.documentElement.setAttribute("data-theme", theme);
})();
