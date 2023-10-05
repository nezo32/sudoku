<script setup lang="ts">
import { provide } from 'vue'
import { computed, ref } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

const header = computed(() => route.name)

const loading = ref(false)

provide('authLoading', loading)
</script>

<template>
  <VContainer class="auth-base">
    <VCard :loading="loading" class="auth-base__main pa-6" rounded="xl" elevation="4">
      <VCardTitle class="mb-6 text-h4">{{ header }}</VCardTitle>
      <RouterView @isLoading="(v: boolean) => (loading = v)" />
    </VCard>
  </VContainer>
</template>

<style scoped lang="scss">
.auth-base {
  max-width: 500px !important;
  height: 100%;

  display: flex;
  justify-content: center;
  align-items: center;

  &__main {
    width: 100%;
  }
}
:global(.auth__snackbar) {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 10px;
}
</style>
