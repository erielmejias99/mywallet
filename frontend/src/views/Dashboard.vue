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
    <v-row>
      <v-col cols="12">
        <transaction-table></transaction-table>
      </v-col>
    </v-row>
    <v-row>
      {{$store.getters.getCurrencies}}
    </v-row>
  </v-container>
</template>

<script>
import CurrencyCards from "@/layout/CurrencyCards";
import DashboardDivider from "@/components/DashboardDivider";
import CreateCurrency from "@/components/CreateCurrency";
import TransactionTable from "@/components/TransactionTable";
export default {
  name: "Dashboard",
  components: {TransactionTable, CreateCurrency, DashboardDivider, CurrencyCards}
}
</script>
