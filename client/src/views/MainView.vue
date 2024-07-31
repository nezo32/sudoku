<script setup lang="ts">
import Welcome from '@/components/Main/Welcome.vue';
import Rules from '@/components/Main/Rules.vue';
</script>

<template>
  <div>
    <v-container>
      <v-card
        class="align-centerfill-height mx-auto"
        color="#72befc"
        width="900"
      >
        <Welcome/>
        <Rules/>
      </v-card>
    </v-container>
  </div>
</template>