import { defineStore } from "pinia"
import { ref } from "vue"


export const useUserIsAuth = defineStore('userIsAuth', () => {
  const userIsAuth = ref(false)

  return{userIsAuth}
}, {persist: true})