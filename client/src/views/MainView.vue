<script setup lang="ts">
</script>

<template>
  <div>
    <v-container>
      <v-card
      class="align-centerfill-height mx-auto"
      max-width="900"
      height="400"
      rounded="lg"
      color="#279eff"
      >
      <v-row>
        <v-col>
          <v-container>
            <v-img src="../assets/logo.png" height="100"></v-img>
            <div class="text-center">
              <h1 class="text-h5 font-weight" style="color: white">Добро пожаловать в</h1>
              <h1 class="text-h3 font-weight-bold" style="color: white">Судоку</h1>
              <v-container class="mt-10 mx-auto">
                <v-btn rounded color="success" size="x-large"
                >Игать</v-btn>
                <div>
                  <br>
                  
                </div>
              </v-container>
            </div>
          </v-container> 
        </v-col>
        <v-col>
          <v-container>
            <v-img src="../assets/gameExample.png" height="350"></v-img>
          </v-container>
        </v-col>
    </v-row>
      </v-card>  
    </v-container>
  </div>
</template>