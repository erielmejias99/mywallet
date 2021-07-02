<template>
  <v-sheet
      rounded
      elevation="3"
      :style="'background: -webkit-linear-gradient(left,' + color1 + ',' +color2 + ')'"
      class="dash-card"
      @click="$emit( 'click' )"
  >
    <div class="d-flex justify-space-between pa-2">
      <div class="mt-1">
        <h2 style="color: white">
          <slot name="content">
            <v-progress-circular color="white" indeterminate></v-progress-circular>
          </slot>
        </h2>
<!--        <div style="color: white" class="font-weight-bold">-->
<!--            <slot name="label">-->
<!--                Loading...-->
<!--            </slot>-->

<!--        </div>-->
      </div>

      <v-avatar size="80" tile color="transparent">
        <slot name="avatar"></slot>
      </v-avatar>
    </div>
    <div class="text-right pa-2">
      <slot name="actions">
      </slot>
    </div>
  </v-sheet>
</template>

<script>
export default {
  name: "BaseDashboardCard",
  props:{
    color1: {
      type: String,
      default: 'lightgreen'
    },
    color2: {
      type: String,
      default: 'green'
    },
  },
}
</script>

<style scoped>

.dash-card{
  cursor: pointer;
  transition: .5s all !important;
}

.dash-card:hover{
  box-shadow: 0 0 20px #999999 !important;
  transition: .5s all !important;
}

</style>
