<script setup lang="ts">
import { useUserStore } from '@/stores'

const userStore = useUserStore()
</script>

<template>
  <VAppBar flat>
    <VContainer class="mx-auto d-flex align-center justify-center">
      <VBtn>Aboba</VBtn>
      <VSpacer />
      <VBtn>{{ userStore.username }}</VBtn>
    </VContainer>
  </VAppBar>
  <VMain>
    <RouterView />
  </VMain>
</template>

<style scoped lang="scss"></style>
