<script setup lang="ts">
import { ref } from 'vue';

const headerLinks = [
  {text:'Играть', href: '/game'},
  {text: 'Мои Рекорды', href: '/records'},
  // {text:'Profile', href: ''},
  // {text: 'Updates', href: ''},
]

const userAuth = ref(false)
</script>

<template>
  <v-app-bar flat color="#2a71f2">
    <v-container class="mx-auto d-flex align-center justify-center">
      <v-btn
        v-for="link in headerLinks"
        :key="link.text"
        :text="link.text"
        :href="link.href"
        variant="text"
        
      ></v-btn>

      <v-spacer></v-spacer>
      
      <v-btn  v-if="!userAuth" href="/auth/login" variant="tonal">Войти</v-btn>
      <v-avatar
        v-if="userAuth"
        class="me-4 "
        color="grey-darken-1"
        size="40"
      ></v-avatar>
    </v-container>
  </v-app-bar>
</template>

<style scoped lang="scss">
</style>