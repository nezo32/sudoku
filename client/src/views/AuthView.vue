<template>
  <v-app id="inspire" v-theme>
    <MainHeader/>
    <v-main class="bg-grey-lighten-3">
      <v-container class="mx-auto d-flex align-center justify-center">
        <v-card
          class="align-centerfill-height mx-auto"
          min-height="auto"
          width="50vh"
          rounded="lg"
          color="#5daff0"
          >
          <div class="text-center">
            <v-img src="../assets/logo.png" height="100"></v-img>
            <h1 class="text-h4 font-weight-bold">Судоку</h1>
          </div>
            <v-container class="mx-auto d-flex align-center justify-center">
              <v-btn variant="text"><RouterLink to="/auth/login" class="text-decoration-none text-white">Log In</RouterLink></v-btn>
              <v-btn  variant="text"><RouterLink to="/auth/register" class="text-decoration-none text-white">Sign in</RouterLink></v-btn>
            </v-container>
          <RouterView/>
        </v-card>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup>



</script>