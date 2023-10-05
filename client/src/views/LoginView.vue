<script setup lang="ts">
import { gql } from '@/generated'
import { useMutation } from '@vue/apollo-composable'
import { ref, watch } from 'vue'
import { useRouter } from 'vue-router'

const emits = defineEmits<{
  (event: 'isLoading', value: boolean): void
}>()

const router = useRouter()

const snackbar = ref(false)

const errorMessage = ref<string>('')

const userData = ref({
  username: '',
  password: ''
})

const { mutate } = useMutation(
  gql(`
    mutation Login($data: AuthData!) {
      login(data: $data)
    }
  `)
)

watch(
  userData,
  () => {
    snackbar.value = false
  },
  { deep: true }
)

async function login() {
  emits('isLoading', true)
  try {
    await mutate({ data: userData.value })
  } catch (err) {
    const msg = (err as Error).message
    errorMessage.value = msg
    snackbar.value = true
  } finally {
    emits('isLoading', false)
  }
  router.push('/')
}
</script>

<template>
  <VTextField variant="outlined" label="Login" v-model="userData.username" />
  <VTextField variant="outlined" label="Password" type="password" v-model="userData.password" />
  <div class="d-flex justify-end">
    <VBtn
      min-width="120"
      height="40"
      variant="text"
      color="primary"
      @click="router.push('/authentication/register')"
    >
      Sign up
    </VBtn>
  </div>
  <div class="d-flex justify-space-between mt-2">
    <VBtn min-width="120" height="40" variant="text" color="primary">Forgot password</VBtn>
    <VBtn color="secondary" min-width="120" height="40" @click="login"> Log in </VBtn>
  </div>
  <VSnackbar color="error" v-model="snackbar" :timeout="2000">
    <div class="auth__snackbar">
      <VIcon icon="mdi-alert-circle" size="30" />
      <span>{{ errorMessage }}</span>
    </div>
  </VSnackbar>
</template>

<style scoped lang="scss"></style>
