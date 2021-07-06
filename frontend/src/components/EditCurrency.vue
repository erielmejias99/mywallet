<template>
  <v-dialog v-model="dialog" max-width="400px">
    <template #activator="{ attrs, on }">
      <span v-bind="attrs" v-on="on">
        <slot></slot>
      </span>
    </template>
    <v-card>
      <v-card-title>Edit USDChange</v-card-title>
      <v-card-text>
        <v-form>
          <v-text-field
          v-model="currency.usd_change"
          maxlength="5"
          placeholder="USD Change"
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
  name: "EditCurrency",
  data: function(){
    return{
      dialog: false,
      loading: false,
      errors: null
    }
  },
  props:{
    currency: {
      type: Object,
      required: true
    }
  },
  methods:{
    save: async function (){
      this.loading = true;
      try{
        await window.backend.CurrencyController.Update( this.currency.ID, parseFloat(this.currency.usd_change)  );
        this.$emit("input",this.currency );
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
