<script setup lang="ts">
import { provide, ref, watchEffect } from "vue"
import { themeKey, type Theme } from "@/composables/useTheme"

const props = withDefaults(
  defineProps<{
    defaultTheme?: Theme
    storageKey?: string
  }>(),
  {
    defaultTheme: "system",
    storageKey: "vite-ui-theme",
  },
)

// Read persisted theme from localStorage, fall back to defaultTheme prop
const theme = ref<Theme>(
  (localStorage.getItem(props.storageKey) as Theme) ?? props.defaultTheme,
)

// watchEffect replaces React's useEffect — re-runs whenever `theme` changes
watchEffect(() => {
  const root = document.documentElement

  root.classList.remove("light", "dark")

  if (theme.value === "system") {
    const systemTheme = window.matchMedia("(prefers-color-scheme: dark)").matches
      ? "dark"
      : "light"
    root.classList.add(systemTheme)
    return
  }

  root.classList.add(theme.value)
})

function setTheme(newTheme: Theme) {
  localStorage.setItem(props.storageKey, newTheme)
  theme.value = newTheme
}

// provide replaces React's Context.Provider — any descendant can inject via useTheme()
provide(themeKey, { theme, setTheme })
</script>

<template>
  <slot />
</template>
