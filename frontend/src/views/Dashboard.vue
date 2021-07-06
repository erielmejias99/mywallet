<template>
  <v-container>
    <currency-cards @transaction_created="addTransToTable" ></currency-cards>
    <v-row>
      <v-col cols="12">
        <transaction-table ref="trans"></transaction-table>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import CurrencyCards from "@/layout/CurrencyCards";
import TransactionTable from "@/components/TransactionTable";

export default {
  name: "Dashboard",
  components: { CurrencyCards, TransactionTable },
  methods:{
    addTransToTable: function( trans ){
      console.log( 'here' );
      console.log( trans );
      this.$refs.trans.addTransaction( trans );
    }
  }
}
</script>
