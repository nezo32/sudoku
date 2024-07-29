<template>
  <v-container>
    <v-card
      class="align-centerfill-height mx-auto"
      min-height="auto"
      width="50vh"
      max-width="390"
      rounded="lg"
      style="background: radial-gradient(circle, rgba(93,175,240,1) 13%, rgba(155,208,251,1) 94%);"
      >
      <div class="text-center">
        <v-img src="../assets/logo.png" height="100"></v-img>
        <h1 class="text-h4 font-weight-bold" style="color: white">Судоку</h1>
      </div>
        <v-container class="mx-auto d-flex align-center justify-center">
          <v-btn variant="text"><RouterLink to="/auth/login" class="text-decoration-none text-white">Log In</RouterLink></v-btn>
          <v-btn  variant="text"><RouterLink to="/auth/register" class="text-decoration-none text-white">Sign in</RouterLink></v-btn>
        </v-container>
      <RouterView/>
    </v-card>
  </v-container>
</template>

<script setup>

</script>