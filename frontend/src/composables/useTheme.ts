import { type InjectionKey, type Ref, inject } from "vue"

export type Theme = "dark" | "light" | "system"

export interface ThemeProviderState {
  theme: Ref<Theme>
  setTheme: (theme: Theme) => void
}

export const themeKey: InjectionKey<ThemeProviderState> = Symbol("theme")

export function useTheme() {
  const context = inject(themeKey)

  if (!context) {
    throw new Error("useTheme must be used within a ThemeProvider")
  }

  return context
}
