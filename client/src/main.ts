import './assets/main.scss'

import { createApp, provide, h } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'
import vuetify from '@/plugins/vuetify'

import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import { DefaultApolloClient } from '@vue/apollo-composable'
import { apolloClient } from '@/apollo'

const app = createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient)
  },
  render: () => h(App)
})
const pinia = createPinia()

pinia.use(piniaPluginPersistedstate)

app.use(pinia)
app.use(vuetify)
app.use(router)

app.mount('#app')
