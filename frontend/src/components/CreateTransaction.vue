<template>
  <v-dialog v-model="dialog" max-width="400">
    <template #activator="{attrs,on}">
      <span v-on="on" v-bind:class="attrs">
        <slot></slot>
      </span>
    </template>
    <v-card>
      <v-card-title>Create Transaction</v-card-title>
      <v-card-text>
        <v-form>

          <v-text-field
          type="number"
          placeholder="Amount"
          v-model="transaction.amount"
          >
            <template #prepend-inner>
              <v-icon color="red" v-if="negative">
                mdi-minus
              </v-icon>
              <v-icon color="green" v-else>
                mdi-plus
              </v-icon>
            </template>
          </v-text-field>
          <div>
            <div class="text-center grey--text"><small>The rest:</small></div>
            <div class="text-center title">
              <template v-if="transaction.amount !== undefined && transaction.amount !== null">
                {{ newAmount }}
              </template>
              <template v-else>
                {{currency.balance}}
              </template>
              {{currency.name}}
            </div>
          </div>

          <v-text-field
          placeholder="Reason"
          v-model="transaction.reason"
          :counter="50"
          maxlength="100"
          prepend-inner-icon="mdi-account-question-outline"
          >
          </v-text-field>

          <v-textarea
          placeholder="Short Description"
          v-model="transaction.description"
          >
          </v-textarea>

          <v-alert v-if="err"  color="error">
            {{err}}
          </v-alert>

        </v-form>
      </v-card-text>

      <v-card-actions class="justify-end">
        <v-btn color="orange" @click="create" :loading="loading" :disabled="loading">
          Save
        </v-btn>

        <v-btn class="ml-2 white gray--tex" shaped @click="dialog=false">
          Cancel
        </v-btn>
      </v-card-actions>

    </v-card>
  </v-dialog>
</template>

<script>
export default {
  name: "CreateTransaction",
  data:()=>({
    dialog: false,
    loading: false,
    err: null,
    transaction:{
      amount: null,
      reason: '',
      description: ''
    }
  }),
  props:{
    currency:{
      type: Object,
      required: true,
    },
    negative: {
      type: Boolean,
      required: true
    }
  },
  methods:{
    create: async function(){
      try{
        let transaction = {
          currency_id: this.currency.ID,
          amount: parseInt( parseFloat( this.transaction.amount ) * 100 ),
          reason: this.transaction.reason,
          description: this.transaction.description
        }
        if( this.negative ){
          transaction.amount = -Math.abs( transaction.amount )
        }else{
          transaction.amount = Math.abs( transaction.amount )
        }
        const resp = await confirm( "Create: " + JSON.stringify(transaction) );
        if( resp ){
          this.$store.dispatch('log', "Try Create Transaction: " + JSON.stringify(transaction) );
          await window.backend.TransactionController.Create( transaction )
          this.dialog = false;
          this.$store.dispatch('log', "Transaction Created." + JSON.stringify(transaction) );
          this.$store.dispatch( 'updateBalance', transaction )
        }
      }catch( error ){
        this.err = error;
        await this.$store.dispatch('log', "Create Transaction Error: " + error );
      }finally {
        this.loading = false;
      }
    }
  },
  computed: {
    newAmount: function(){
      if( this.negative ){
        return parseFloat( this.currency.balance ) - Math.abs( this.transaction.amount )
      }else{
        return parseFloat( this.currency.balance ) + Math.abs( this.transaction.amount )
      }
    }
  }
}
</script>

<style scoped>

</style>