<template>
  <base-dashboard-card :color1="color1" :color2="color2">
    <template #avatar>
      <span class="title white--text">
        <b>
          {{currency.name.toUpperCase()}}
        </b>
      </span>
<!--      <v-icon color="white" size="40">{{icon}}</v-icon>-->
    </template>

    <template #content>
      <v-card-title>{{currency.balance}}</v-card-title>
      <v-card-subtitle>

        <template v-if="currency.balance > 0 && currency.usd_change && currency.name !== 'USD'">
        &eqsim; <b>{{parseFloat( currency.balance * currency.usd_change ).toFixed(2)}}</b> USD
        </template>
        <template v-else>
          &nbsp;None
        </template>

      </v-card-subtitle>
    </template>
    <template #actions>
      <create-transaction :currency="currency" negative >
        <v-btn x-small fab color="white" class="mx-1">
          <v-icon small color="red">mdi-minus</v-icon>
        </v-btn>
      </create-transaction>

      <create-transaction :currency="currency" :negative="false" @created="createdTransaction">
        <v-btn x-small fab color="white" class="mx-1">
          <v-icon small color="green">mdi-plus</v-icon>
        </v-btn>
      </create-transaction>

      <v-btn x-small fab color="white" class="mx-1">
        <v-icon small @click="deleteCurrency( currency )" color="red">mdi-delete</v-icon>
      </v-btn>

      <edit-currency :currency.sync="currency">
        <v-btn x-small fab color="white" class="mx-1">
          <v-icon small color="primary">mdi-pen</v-icon>
        </v-btn>
      </edit-currency>

    </template>

  </base-dashboard-card>
</template>

<script>
import BaseDashboardCard from "@/components/BaseDashboardCard";
import CreateTransaction from "@/components/CreateTransaction";
import EditCurrency from "@/components/EditCurrency";
export default {
  name: "CurrencyCard",
  components: {EditCurrency, CreateTransaction, BaseDashboardCard},
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
  },
  methods:{
    createdTransaction: function(trans){
      console.log( "createdTransaction");
      console.log( trans );
      this.$emit( 'newTransaction', trans)
    },
    deleteCurrency: async function(currency){
      const resp = await confirm("Are you sure? Currency " + currency.name + " will be deleted and all its transactions");
      if( resp ){
        try{
          console.log( currency );
          await window.backend.CurrencyController.Delete( currency.ID );
          this.$emit( 'deleted', currency );
        }catch (err){
          alert( err );
        }
      }
    }
  }
}
</script>