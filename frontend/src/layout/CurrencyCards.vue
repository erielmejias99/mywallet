<template>
  <v-row>
    <v-col cols="12" sm="6" md="4" lg="3" xl="3" v-for="currency in $store.getters.getCurrencies" :key="currency.name">
      <currency-card :currency="currency" icon="mdi-currency-usd" color1="orange" color2="red"></currency-card>
    </v-col>
  </v-row>
</template>

<script>
import CurrencyCard from "@/components/CurrencyCard";
export default {
  name: "CurrencyCards",
  components: {CurrencyCard},
}
</script>