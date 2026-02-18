<script setup lang="ts">
import { computed } from "vue"
import { Moon, Sun } from "lucide-vue-next"
import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { useTheme } from "@/composables/useTheme"

const { theme, setTheme } = useTheme()

// Derive whether the current effective theme is dark.
// For "system", we read the OS preference — this matches what ThemeProvider applies to the DOM.
const isDark = computed(() => {
  if (theme.value === "dark") return true
  if (theme.value === "light") return false
  return window.matchMedia("(prefers-color-scheme: dark)").matches
})
</script>

<template>
  <DropdownMenu>
    <DropdownMenuTrigger as-child>
      <Button variant="outline" size="icon">
        <!--
          Vue approach: drive icon visibility from the reactive `isDark` computed
          instead of Tailwind's dark: class variant, which would couple the icon
          state to the DOM class rather than the reactive theme ref.
        -->
        <Sun
          class="h-[1.2rem] w-[1.2rem] transition-all"
          :class="isDark ? 'scale-0 -rotate-90' : 'scale-100 rotate-0'"
        />
        <Moon
          class="absolute h-[1.2rem] w-[1.2rem] transition-all"
          :class="isDark ? 'scale-100 rotate-0' : 'scale-0 rotate-90'"
        />
        <span class="sr-only">Toggle theme</span>
      </Button>
    </DropdownMenuTrigger>
    <DropdownMenuContent align="end">
      <DropdownMenuItem @click="setTheme('light')">Light</DropdownMenuItem>
      <DropdownMenuItem @click="setTheme('dark')">Dark</DropdownMenuItem>
      <DropdownMenuItem @click="setTheme('system')">System</DropdownMenuItem>
    </DropdownMenuContent>
  </DropdownMenu>
</template>
