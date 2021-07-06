<template>
  <v-app>

    <v-main>
      <dashboard></dashboard>
    </v-main>

    <v-footer fixed class="justify-space-between">
      <div class="text-right">
        Mady with <v-icon color="red">mdi-heart</v-icon>
        by <a href="https://erielmejias99.herokuapp.com" target="_blank">erielmejias99@gmail.com</a>
      </div>
    </v-footer>
  </v-app>
</template>

<script>
import Dashboard from "./views/Dashboard";
export default {
  components: {Dashboard},
}
</script>