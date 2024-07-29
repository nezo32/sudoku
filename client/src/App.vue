<template>
  <v-app :theme="theme">
    <Header/>
    <v-main class="overflow-hidden">
      <router-view />
    </v-main>
  </v-app>
</template>

<script lang="ts" setup>
import { computed } from 'vue';
import Header from './components/Header/Header.vue';
import { useSiteTheme } from '@/stores/siteTheme';

const store = useSiteTheme()

const theme = computed(() =>  store.themeIsLight ? "light" : "dark")
</script>
