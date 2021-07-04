<template>
  <div>
    <v-toolbar dense flat>

      <v-icon v-if="icon && icon.length" :color="icon_color" size="30px">
        {{icon}}
      </v-icon>

      <v-toolbar-title>
        {{title}}
      </v-toolbar-title>

      <v-spacer></v-spacer>

      <slot name="actions">
      </slot>

    </v-toolbar>
    <v-divider class="mb-2 grey--text" ></v-divider>
  </div>
</template>

<script>

export default {
  name: "DashboardDivider",
  props:{
    icon: {
      required: false,
      type: String
    },
    icon_color: {
      required: false,
      type: String,
      default: "black"
    },
    title:{
      required: true,
    }
  }
}

</script>
