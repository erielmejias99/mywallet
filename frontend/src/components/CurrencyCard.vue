<template>
  <v-card>
    <div>{{currency.name}}</div>
    <v-card-title>{{currency.balance}}</v-card-title>
  </v-card>
</template>

<script>
export default {
  name: "CurrencyCard",
  props:{
    currency: {
      type: Object,
      required: true
    }
  }
}
</script>