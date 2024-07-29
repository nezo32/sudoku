<template>
  <v-card
    class="mx-auto px-6 py-8"
    max-width="344"
  >
    <v-form
      v-model="form"
      @submit.prevent="onSubmit"
    >
      <v-text-field
        v-model="nick"
        color="primary"
        label="Nickname"
        clearable
      ></v-text-field>

      <v-text-field
          v-model="email"
          :readonly="loading"
          :rules="[rules.emailRequired, rules.emailValid]"
          class="mb-2"
          label="Email"
          clearable
        ></v-text-field>

        <v-text-field
          :append-inner-icon="visiblePass1 ? 'mdi-eye-off' : 'mdi-eye'"
          :rules="[rules.passRequired, rules.passMin]"
          :type="visiblePass1 ? 'text' : 'password'"
          v-model="password1"
          density="default"
          placeholder="Enter your password"
          prepend-inner-icon="mdi-lock-outline"
          clearable
          @click:append-inner="visiblePass1 = !visiblePass1"
      ></v-text-field>

      <v-text-field
        :append-inner-icon="visiblePass2 ? 'mdi-eye-off' : 'mdi-eye'"
        :rules="[rules.passRequired, rules.passSimilarities]"
        :type="visiblePass2 ? 'text' : 'password'"
        v-model="password2"
        density="default"
        placeholder="Repeet your password"
        prepend-inner-icon="mdi-lock-outline"
        clearable
        @click:append-inner="visiblePass2 = !visiblePass2"
      ></v-text-field>
    </v-form>

    <v-divider></v-divider>

    <v-card-actions>
      <v-spacer></v-spacer>

      <v-btn c
      :disabled="!form"
        :loading="loading"
        color="success"
        size="large"
        type="submit"
        variant="elevated"
        block
      >
        Complete Registration

        <v-icon icon="mdi-chevron-right" end></v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
  <br>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const form = ref(false)
const email = ref(null)
const password1 = ref(null)
const password2 = ref(null)
const loading = ref(false)
const nick = ref("")
const visiblePass1 = ref(false)
const visiblePass2 = ref(false)
const rules = ref({
  passRequired: (value : string ) => !!value || 'Required.',
  passMin: (v : string) => v.length >= 8 || 'Min 8 characters',
  passSimilarities: (v: string) => v === password1.value || "Passwords don't match",
  emailRequired: (v: string) => !!v || 'Field is required',
  emailValid: (v: string) => (/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/).test(v) ||'Invalid email address!'
})

const onSubmit = () => {
  if (!form.value) return

  loading.value = true

  setTimeout(() => (loading.value = false), 2000)
}
</script>

<style scoped lang="scss">
</style>