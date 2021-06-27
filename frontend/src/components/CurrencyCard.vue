<template>
  <base-dashboard-card :color1="color1" :color2="color2">
    <template #avatar>
      <v-icon color="white" size="40">{{icon}}</v-icon>
    </template>

    <template #content>
      <v-card-title>{{currency.balance}}</v-card-title>
      <v-card-subtitle v-if="currency.balance > 0 && currency.usd_change && currency.name !== 'USD'">
        ~{{currency.balance * currency.usd_change}}
      </v-card-subtitle>
    </template>

    <template #label>
      <div>{{currency.name}}</div>
    </template>

  </base-dashboard-card>
</template>

<script>
import BaseDashboardCard from "@/components/BaseDashboardCard";
export default {
  name: "CurrencyCard",
  components: {BaseDashboardCard},
  props:{
    icon: {
      type: String,
      required: true,
    },
    color1:{
      type: String,
      required: true,
    },
    color2:{
      type: String,
      required: true,
    },
    currency: {
      type: Object,
      required: true
    }
  }
}
</script>