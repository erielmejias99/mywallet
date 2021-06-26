<template>
  <v-row v-if="status==='ok'">
    <v-col cols="12" sm="6" md="4" lg="3" xl="3" v-for="currency in currencies" :key="currency.name">
      <currency-card :currency="currency"></currency-card>
    </v-col>
  </v-row>
  <div v-else-if="status==='loading'" class="text-center">
    Loading...
  </div>
  <div v-else class="red--text">
    Error occurred...
  </div>
</template>

<script>
import CurrencyCard from "@/components/CurrencyCard";
export default {
  name: "CurrencyCards",
  components: {CurrencyCard},
  data: function (){
    return{
      currencies: [],
      status: 'loading'
    }
  },
  methods:{
    getCurrencies: async function (){
      this.status = 'loading'
      try{
        this.currencies =  await window.backend.CurrencyRep.GetAll();
        this.status = 'ok'
      }catch( err ){
        this.status = 'error';
        console.log( err );
      }
    }
  },
  mounted() {
    this.getCurrencies();
  }
}
</script>