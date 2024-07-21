<template>
    <v-card class="mx-auto px-6 py-8" color="grey-lighten-3" max-width="344">
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
        :rules="[rules.passRequired, rules.passMin]"
        :type="visible ? 'text' : 'password'"
        v-model="password"
        density="default"
        placeholder="Enter your password"
        prepend-inner-icon="mdi-lock-outline"
        variant="outlined"
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

const form = ref(false)
const email = ref(null)
const password = ref(null)
const loading = ref(false)
const visible = ref(false)
const rules = ref({
  passRequired: (value : string ) => !!value || 'Required.',
  passMin: (v : string) => v.length >= 8 || 'Min 8 characters',
  emailRequired: (v: string) => !!v || 'Field is required',
  emailValid: (v: string) => (/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/).test(v) ||'Invalid email address!'
})

const onSubmit = () => {
  if (!form.value) return

  loading.value = true

  setTimeout(() => (loading.value = false), 2000)
}
</script>