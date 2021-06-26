<template>
  <v-container>
    <currency-cards></currency-cards>
  </v-container>
</template>

<script>
import CurrencyCards from "@/layout/CurrencyCards";
export default {
  name: "Dashboard",
  components: {CurrencyCards}
}
</script>
