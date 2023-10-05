<script setup lang="ts">
import { onMounted, watch, computed } from 'vue'
import { useTheme } from 'vuetify'
import { useThemeStore } from '@/stores'
import { storeToRefs } from 'pinia'
import { useRoute } from 'vue-router'
import MainBase from './components/bases/MainBase.vue'

const route = useRoute()

const theme = useTheme()
const themeStore = useThemeStore()
const { theme: currentTheme } = storeToRefs(themeStore)

const mainView = computed(() => !route.fullPath.includes('authentication'))

const mediaQuery = '(prefers-color-scheme: dark)'

watch(currentTheme, (n) => {
  currentThemeHandler(n)
  switch (n) {
    case undefined:
      themeStore.setSystemTheme()
      break
    case 'light':
      themeStore.setLightTheme()
      break
    case 'dark':
      themeStore.setDarkTheme()
      break
  }
})

onMounted(() => {
  if (currentTheme.value) {
    theme.global.name.value = currentTheme.value
  } else {
    currentThemeHandler(currentTheme.value)
  }
})

function currentThemeHandler(value: Theme | undefined) {
  if (!value) {
    setMediaListener()
    theme.global.name.value = getCurrentSystemTheme()
  } else {
    removeMediaListener()
    theme.global.name.value = value
  }
}
function mediaChangeHandler(ev: MediaQueryListEvent) {
  theme.global.name.value = ev.matches ? 'dark' : 'light'
}
function setMediaListener() {
  window.matchMedia(mediaQuery).addEventListener('change', mediaChangeHandler)
}
function removeMediaListener() {
  window.matchMedia(mediaQuery).removeEventListener('change', mediaChangeHandler)
}
function getCurrentSystemTheme() {
  return window.matchMedia(mediaQuery).matches ? 'dark' : 'light'
}
</script>

<template>
  <VApp>
    <MainBase v-if="mainView" />
    <RouterView v-else />
  </VApp>
</template>

<style scoped lang="scss"></style>
