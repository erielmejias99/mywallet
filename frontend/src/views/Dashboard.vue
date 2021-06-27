<template>
  <v-container>
    <dashboard-divider title="Wallets">
      <template #actions>
        <create-currency>
          <v-btn icon color="orange">
            <v-icon>mdi-plus</v-icon>
          </v-btn>
        </create-currency>
      </template>
    </dashboard-divider>
    <currency-cards></currency-cards>
  </v-container>
</template>

<script>
import CurrencyCards from "@/layout/CurrencyCards";
import DashboardDivider from "@/components/DashboardDivider";
import CreateCurrency from "@/components/CreateCurrency";
export default {
  name: "Dashboard",
  components: {CreateCurrency, DashboardDivider, CurrencyCards}
}
</script>
