<script setup lang="ts">
import { computed } from 'vue';
import { useSiteTheme } from '@/stores/siteTheme';
import { useUserIsAuth } from '@/stores/userIsAuth';
import UserMenu from '@/components/Header/UserMenu.vue';

const themeStore = useSiteTheme()
const userAuthStore = useUserIsAuth()

const headerLinks = computed(() => [
  {text:'Главная', to: '/'},
  {text: 'Мои Рекорды', to: '/records',},
])
</script>

<template>
  <v-app-bar flat absolute color="#40A9FF">
    <v-container class="mx-auto d-flex align-center  text-white">
      <v-btn
        v-for="(link, i) of headerLinks"
        :key="i"
        :text="link.text"
        variant="text"
        :active="false"
        :to="link.to"
      ></v-btn>

      <v-spacer></v-spacer>

      <v-switch
        v-model="themeStore.themeIsLight"
        class="me-10"
        hide-details
        style="color: white;"
        true-icon="mdi-white-balance-sunny"
        false-icon="mdi-weather-night"
        ></v-switch>
      <v-btn  v-if="!userAuthStore.userIsAuth" href="/auth/login" variant="tonal" color="white"
      >Войти</v-btn>
      <UserMenu v-if="userAuthStore.userIsAuth"/>
    </v-container>
  </v-app-bar>
</template>