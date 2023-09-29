import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useThemeStore = defineStore(
  'theme',
  () => {
    const theme = ref<Theme>()

    function setLightTheme() {
      theme.value = 'light'
    }
    function setDarkTheme() {
      theme.value = 'dark'
    }
    function setSystemTheme() {
      theme.value = undefined
    }

    return { theme, setLightTheme, setDarkTheme, setSystemTheme }
  },
  { persist: true }
)
