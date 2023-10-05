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

const retype = ref('')

const { mutate } = useMutation(
  gql(`
    mutation Register($data: AuthData!) {
      register(data: $data)
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
async function register() {
  if (retype.value !== userData.value.password) {
    errorMessage.value = "Passwords aren't matching"
    snackbar.value = true
    return
  }
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
  router.push('/authentication/login')
}
</script>

<template>
  <VTextField variant="outlined" label="Login" v-model="userData.username" />
  <VTextField variant="outlined" label="Password" type="password" v-model="userData.password" />
  <VTextField variant="outlined" label="Retype password" type="password" v-model="retype" />
  <div class="d-flex justify-space-between">
    <VBtn
      min-width="120"
      height="40"
      variant="text"
      color="primary"
      @click="router.push('/authentication/login')"
    >
      Already have account?
    </VBtn>
    <VBtn color="secondary" min-width="120" height="40" @click="register"> Register </VBtn>
  </div>
  <VSnackbar color="error" v-model="snackbar" :timeout="2000">
    <div class="auth__snackbar">
      <VIcon icon="mdi-alert-circle" size="30" />
      <span>{{ errorMessage }}</span>
    </div>
  </VSnackbar>
</template>

<style scoped lang="scss"></style>
