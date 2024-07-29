<script setup lang="ts">
import { useUserIsAuth } from '@/stores/userIsAuth';
// import { ref } from 'vue';

const userAuthStore = useUserIsAuth()


// const menu = ref(false)

const unlog = async () =>{
  userAuthStore.userIsAuth = false
  // перезагрузить?
}
</script>

<template>
  <div>
    <v-menu
      location="bottom"
    >
      <template v-slot:activator="{props}">
        <v-btn 
          rounded="xl"
          v-bind="props"
        >
          <v-avatar image="/src/assets/icon.png"></v-avatar>
          romurs <!--Никнейм -->
        </v-btn>
      </template>
      <v-card min-width="200" class="mt-3 ml-1" rounded="lg">
        <v-divider></v-divider>

        <v-list>
          <v-list-item>
            <v-btn 
            variant="text"
            :active="false"
            to="/profile"
            >Профиль</v-btn>
          </v-list-item>

          <v-list-item>
            <v-btn
            variant="text"
            :active="false"
            to="/"
            @click="unlog"
            >Выйти</v-btn>
          </v-list-item>
        </v-list>

        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            variant="text"
          >
            Закрыть
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-menu>
  </div>
</template>