<template>
  <v-row>
    <v-col cols="12">
      <dashboard-divider title="Wallets" icon="mdi-wallet-outline" icon_color="orange">
        <template #actions>
          <create-currency @created="addCurrency">
            <v-btn icon color="orange">
              <v-icon>mdi-plus</v-icon>
            </v-btn>
          </create-currency>
        </template>
      </dashboard-divider>
    </v-col>
    <v-col cols="12" sm="6" md="4" lg="3" xl="3" v-for="( currency, i ) in currencies" :key="currency.name">
      <currency-card
      :currency="currency"
      :icon="card_design[i].icon"
      :color1="card_design[i].first"
      :color2="card_design[i].second"
      @newTransaction="updateCurrencyAmount"
      @deleted="deleteCurrency"
      ></currency-card>
    </v-col>
  </v-row>
</template>

<script>
import CurrencyCard from "@/components/CurrencyCard";
import DashboardDivider from "@/components/DashboardDivider";
import CreateCurrency from "@/components/CreateCurrency";
export default {
  name: "CurrencyCards",
  components: {CreateCurrency, DashboardDivider, CurrencyCard },
  data: function(){
    return{
      currencies: [],
      card_design: [
          { first: "#f3a183", second: '#ec6f66', icon: 'mdi-cash' },
          { first: "#43cea2", second: '#185a9d', icon: 'mdi-currency-usd' },
          { first: "#a8e063", second: '#56ab2f', icon: 'mdi-bitcoin' },
          { first: "orange", second: 'blue', icon: 'mdi-currency-usd' },
          { first: "orange", second: 'blue', icon: 'mdi-currency-usd' },
          { first: "orange", second: 'blue', icon: 'mdi-currency-usd' },
          { first: "orange", second: 'blue', icon: 'mdi-currency-usd' },
      ]
    }
  },
  methods:{
    updateCurrencyAmount: function( transaction ){
      console.log( 'emit received' );
      console.log(transaction)
      this.$emit( 'transaction_created', transaction)
      for( let i = 0; i < this.currencies.length; i++ ){
        if( this.currencies[i].ID === transaction.currency_id ){
          this.currencies[ i ].balance += ( transaction.amount / 100.0 );
          break;
        }
      }
    },
    addCurrency: function( newCurrency ){
      this.currencies.push( newCurrency )
    },
    deleteCurrency: function( currency ){
      for( let index = 0; index < this.currencies.length; index ++ ){
        if( this.currencies[ index ].ID === currency.ID ){
          this.currencies.splice( index, 1 );
          return
        }
      }
    }
  },
  mounted(){
    window.backend.CurrencyController.GetAll().then( resp =>{
      this.currencies = resp;
    })
  }
}
</script>