<template>
  <v-data-table
  :items="transactions"
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

    <template #item.CreatedAt="{item}">
      <div>
        {{ ( new Date(item.CreatedAt) ).toLocaleDateString() }}<br>
        {{ ( new Date(item.CreatedAt) ).toLocaleTimeString() }}<br>
      </div>
    </template>

  </v-data-table>
</template>

<script>
export default {
  name: "TransactionTable",
  data: ()=>({
    transactions: [],
    err: null,
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
    ]
  }),
  props:{
    currency_id:{
      type: Number,
      required: false,
      default: 0
    }
  },
  methods:{
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
    }
  },
  mounted() {
    this.loadTransactions();
  }
}
</script>

<style scoped>

</style>