<template>
  <v-app>

    <v-main>
      <router-view></router-view>
    </v-main>

    <v-footer fixed class="justify-space-between">
      <div class="text-right">
        Mady with <v-icon color="red">mdi-heart</v-icon>
        by <a href="https://erielmejias99.herokuapp.com" target="_blank">erielmejias99@gmail.com</a>
      </div>
      <v-menu
      top
      open-on-hover
      max-width="500"
      min-width="500"
      max-height="400"
      offset-y
      >
        <template #activator="{on}">
          <div style="max-width: 500px; max-height: 20px;" v-on="on">
            <v-icon color="orange">
              mdi-message-alert-outline
            </v-icon>
            <span v-if="$store.getters.areLogs">
              {{ $store.getters.lastLog.toString().substring(0, 400)}}
            </span>
            <span v-else>
              Logs...
            </span>
          </div>
        </template>

        <v-sheet class="pa-2" >
          <div class="d-flex justify-space-between">
            <span class="font-weight-bold title orange--text">Logs:</span>
            <v-btn color="gray" @click="$store.dispatch('clearLogs')" text small>
              Clear All
            </v-btn>
          </div>
          <v-divider></v-divider>
          <template v-if="$store.getters.areLogs">
            <div v-for="( item, i ) in $store.getters.getLogs" :key="i">
              <div class="pa-1">
                {{item}}
              </div>
              <v-divider class="mb-1"></v-divider>
            </div>
          </template>
          <template v-else>
            <div class="pa-12 text-center">
              No Logs...
            </div>
          </template>
        </v-sheet>
      </v-menu>
    </v-footer>
  </v-app>
</template>

<script>
  export default {
    data: function(){ return {
        navItems: [
          { icon: 'mdi-view-dashboard', label: "Dashboard", to: { name: "Dashboard" } },
          { icon: 'mdi-bank-transfer', label: "Transaction", to: { name: "Transaction" } },
        ]
      }
    },
    async mounted(){
      try{
        await window.backend.CurrencyController.GetAll()
        this.$store.dispatch( 'log', "Currencies loaded." )
      }catch(error){
        this.$store.dispatch( 'log', error )
        console.log( error );
      }
    }
  }
</script>