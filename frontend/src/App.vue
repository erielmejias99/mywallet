<template>
  <v-app>

    <v-main>
      <router-view></router-view>
    </v-main>

    <v-footer fixed class="justify-end">
      <div class="text-right">
        Mady with <v-icon color="red">mdi-heart</v-icon>
        by <a href="https://erielmejias99.herokuapp.com" target="_blank">erielmejias99@gmail.com</a>
      </div>
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
        const resp = await window.backend.CurrencyController.GetAll()
        this.$store.dispatch("updateCurrencies", resp );
      }catch(error){
        console.log( error );
      }
    }
  }
</script>