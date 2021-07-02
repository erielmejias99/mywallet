<template>
  <v-row>
    <v-col cols="12" sm="6" md="4" lg="3" xl="3" v-for="( currency, i ) in $store.getters.getCurrencies" :key="currency.name">
      <currency-card
      :currency="currency"
      :icon="card_design[i].icon"
      :color1="card_design[i].first"
      :color2="card_design[i].second"
      ></currency-card>
    </v-col>
  </v-row>
</template>

<script>
import CurrencyCard from "@/components/CurrencyCard";
export default {
  name: "CurrencyCards",
  components: { CurrencyCard },
  data: function(){
    return{
      card_design: [
          { first: "purple", second: 'darkblue', icon: 'mdi-cash' },
          { first: "orange", second: 'red', icon: 'mdi-currency-usd' },
          { first: "blue", second: 'darkblue', icon: 'mdi-bitcoin' },
          { first: "orange", second: 'blue', icon: 'mdi-currency-usd' },
          { first: "orange", second: 'blue', icon: 'mdi-currency-usd' },
          { first: "orange", second: 'blue', icon: 'mdi-currency-usd' },
          { first: "orange", second: 'blue', icon: 'mdi-currency-usd' },
      ]
    }
  },
}
</script>