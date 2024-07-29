<template>
    <v-card class="mx-auto px-6 py-8" color="" max-width="344">
      <v-form
        v-model="form"
        @submit.prevent="onSubmit"
      >
        <v-text-field
          v-model="email"
          :readonly="loading"
          :rules="[rules.emailRequired, rules.emailValid]"
          class="mb-2"
          label="Email"
          clearable
        ></v-text-field>

        <v-text-field
          :append-inner-icon="visible ? 'mdi-eye-off' : 'mdi-eye'"
          :rules="[rules.passRequired]"
          :type="visible ? 'text' : 'password'"
          v-model="password"
          density="default"
          placeholder="Enter your password"
          prepend-inner-icon="mdi-lock-outline"
          @click:append-inner="visible = !visible"
      ></v-text-field>
      
        <br>

        <v-btn
          :disabled="!form"
          :loading="loading"
          color="success"
          size="large"
          type="submit"
          variant="elevated"
          block
        >
          Log In
        </v-btn>
      </v-form>
    </v-card>
    <br>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useUserIsAuth } from '@/stores/userIsAuth';
import { useRouter } from 'vue-router';

const userAuthStore = useUserIsAuth()

const router = useRouter()

const form = ref(false)
const email = ref(null)
const password = ref(null)
const loading = ref(false)
const visible = ref(false)
const rules = ref({
  passRequired: (value : string ) => !!value || 'Required.',
  emailRequired: (v: string) => !!v || 'Field is required',
  emailValid: (v: string) => (/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/).test(v) ||'Invalid email address!'
})

const onSubmit = async () => {
  if (!form.value) return

  loading.value = true

  try{
  // запрос в БД
  // получение данных
  // запись данных
  userAuthStore.userIsAuth = true
  router.push('/')

  loading.value = false
  }
  catch{
    // выводить ошибку
  }
}
</script>