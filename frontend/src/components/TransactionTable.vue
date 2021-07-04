<template>
  <div>
    <v-toolbar dense flat>
      <v-text-field
      dense
      outlined
      prepend-inner-icon="mdi-magnify"
      placeholder="Search..."
      v-model="search"
      hide-details
      clearable
      />
      <v-spacer></v-spacer>

      <v-menu offset-y close-on-click :close-on-content-click="false" >
        <template #activator="{on}">
          <v-btn v-on="on" color="orange" class="white--text">
            <v-icon left>mdi-calendar</v-icon>
            Select Date
            {{date_filter}}
          </v-btn>
        </template>

        <v-date-picker v-model="date_filter" type="month">
        </v-date-picker>

      </v-menu>

    </v-toolbar>
    <v-toolbar dense flat>
      <v-chip-group>
        <v-chip
        color="purple"
        pill
        class="white--text"
        v-for="currency of $store.getters.getCurrencies"
        :key="currency.ID"
        @click="search=currency.name"
        >
          {{ currency.name }}
        </v-chip>
      </v-chip-group>
      <v-spacer></v-spacer>
      <template v-for="item in total_by_currency">
        <span :key="item[0]">
          {{item[1]/100.0}} <b>{{item[0]}}</b>
        </span>
        <v-divider vertical :key="item[ 0 ] + '_divider'" class="mx-3"></v-divider>
      </template>
    </v-toolbar>

    <v-data-table
    :items="filterTransactions"
    :headers="headers"
    >
      <template #no-data>
        <v-alert v-if="err != null" type="error">
          {{ err }}
        </v-alert>
        <div v-else>
          No data
        </div>
      </template>

      <template #item.ID="{item}">
        <v-btn fab small :color="(parseFloat(item.amount) < 0) ? 'red' : 'green'" class="white--text">
          {{item.ID}}
        </v-btn>
      </template>

      <template #item.CreatedAt="{item}">
        <v-menu max-width="290" min-width="275" open-on-hover close-on-click :close-on-content-click="false" >
          <template #activator="{on}">
            <div v-on="on">
              {{ ( new Date(item.CreatedAt) ).toLocaleDateString().replaceAll('/', '-') }} <b>-</b>
              {{ ( new Date(item.CreatedAt) ).toLocaleTimeString() }}
            </div>
          </template>
          <template #default>
            <v-date-picker
            :value="getGeneralDate(item.CreatedAt)"
            readonly
            >
            </v-date-picker>
          </template>
        </v-menu>
      </template>

      <template #item.amount="{item}">
        {{ ( parseFloat( item.amount ) / 100.0).toFixed(2) }}
        <b>{{$store.getters.getCurrencyFromId(item.currency_id).name}}</b>
      </template>

    </v-data-table>
  </div>
</template>

<script>
export default {
  name: "TransactionTable",
  data: ()=>({
    transactions: [],
    err: null,
    search: null,
    date_filter: null,
    headers:[
      {
        text: 'Id',
        value: 'ID'
      },
      {
        text: "Amount",
        value: "amount"
      },
      {
        text: "Reason",
        value: "reason"
      },
      {
        text: "Description",
        value: "description"
      },
      {
        text: "Date",
        value: "CreatedAt"
      }
    ],
    total_by_currency: new Map()
  }),
  props:{
    currency_id:{
      type: Number,
      required: false,
      default: 0
    }
  },
  methods:{
    getGeneralDate: function( date_string ){
      let d = new Date(date_string);
      return d.getFullYear() + '-' + d.getMonth() + '-' + d.getDate();
    },
    loadTransactions: function loadTransactions(){
      this.loading = true;
      this.err = null;
      window.backend.TransactionController.List( this.currency_id )
        .then( resp=>{
          this.transactions = resp;
          console.log( resp );
          this.loading = false;
        })
        .catch( err=>{
          this.err = err;
          this.transactions = [];
        })
    }
  },
  watch:{
    currency_id: function(){
      this.loadTransactions()
    },
    date_filter: function(newVal){
      this.search = newVal;
    },
    search: function(){
      this.total_by_currency.clear()
      let filtered = this.filterTransactions;
      // fil the map which calculates the total amount by currency on current filter
      for( let item of filtered ){
        let currency_name = this.$store.getters.getCurrencyFromId(item.currency_id).name
        if( this.total_by_currency.has( currency_name )){
          this.total_by_currency.set(
              currency_name,
              this.total_by_currency.get( currency_name ) + parseFloat( item.amount )
          );
        }else{
          this.total_by_currency.set( currency_name, parseFloat( item.amount ) );
        }
      }
    }
  },
  computed:{
    filterTransactions: function (){
      if( this.search !== null && this.search !== undefined && this.search !== '' ){

        return this.transactions.filter(item => (
            item.ID.toString() === this.search ||
            item.amount.toString() === this.search ||
            item.reason.toString().indexOf(this.search) !== -1 ||
            item.description.toString().indexOf(this.search) !== -1 ||
            item.CreatedAt.toString().indexOf(this.search) !== -1 ||
            this.$store.getters.getCurrencyFromId(item.currency_id).name === this.search
        ));
      }
      return this.transactions;
    }
  },
  mounted() {
    this.loadTransactions();
  }
}
</script>

<style scoped>

</style>