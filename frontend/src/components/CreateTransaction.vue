<template>
  <v-dialog v-model="dialog">
    <template #activator="{attrs,on}">
      <span v-on="on" v-bind:class="attrs">
        <slot></slot>
      </span>
    </template>
  </v-dialog>
</template>

<script>
export default {
  name: "CreateTransaction",
  data:()=>({
    transaction:{
      amount: 0,
    }
  }),
  props:{
    negative: {
      type: Boolean,
      required: true
    }
  }
}
</script>

<style scoped>

</style>