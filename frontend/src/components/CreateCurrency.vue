<template>
  <v-dialog v-model="dialog" max-width="400px">
    <template #activator="{ attrs, on }">
      <span v-bind="attrs" v-on="on">
        <slot></slot>
      </span>
    </template>
    <v-card>
      <v-card-title>Add Currency</v-card-title>
      <v-card-text>
        <v-form>
          <v-text-field
          v-model="currency.name"
          maxlength="5"
          placeholder="Name"
          type="text"
          />

          <v-text-field
          v-model="currency.usd_change"
          maxlength="5"
          placeholder="change"
          type="number"
          />

        </v-form>

        <v-alert v-if="errors"  color="error">
          {{errors}}
        </v-alert>

      </v-card-text>
      <v-card-actions class="justify-end">
        <v-btn color="orange" @click="save" :loading="loading" :disabled="loading">
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
  name: "CreateCurrency",
  data: function(){
    return{
      dialog: false,
      loading: false,
      currency: {
        name: "",
        balance: 0,
        usd_change: 0
      },
      errors: null
    }
  },
  methods:{
    save: async function (){
      this.loading = true;
      try{
        this.currency.usd_change = Number.parseFloat( this.currency.usd_change );
        const resp = await window.backend.CurrencyController.Create( this.currency );
        console.log( resp )
        this.$store.dispatch("addCurrency", this.currency );
        this.dialog = false;
      }catch (err){
        console.log( err );
        this.errors = err;
      }finally {
        this.loading=false;
      }
    }
  }
}
</script>
