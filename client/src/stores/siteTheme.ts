import { defineStore } from "pinia"
import { ref } from "vue"


export const useSiteTheme = defineStore('siteTheme', () => {
  const themeIsLight = ref(false)
  
  return{themeIsLight}
}, {persist: true})