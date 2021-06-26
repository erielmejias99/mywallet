<template>
  <v-app>
    <v-navigation-drawer app permanent expand-on-hover clipped>
      <v-list>

        <v-list-item>
          <v-list-item-avatar>
            <v-avatar color="red">
              <v-img :src="require('./assets/wallpaper2.jpg')"></v-img>
            </v-avatar>
          </v-list-item-avatar>
          <v-list-item-content>
            <v-list-item-title>Eriel Mejias</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
        <v-list-item v-for="item in navItems" :key="item.label" link :to="item.to">
          <v-list-item-icon>
            <v-icon color="orange">{{ item.icon }}</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{item.label}}</v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <router-view></router-view>
    </v-main>

    <v-footer fixed>
      Eriel Mejias
      <v-icon>mdi-dashboard</v-icon>
    </v-footer>
  </v-app>
</template>

<script>
  export default {
    data: function(){ return {
        navItems: [
          { icon: 'mdi-view-dashboard', label: "Dashboard", to: { name: "Dashboard" } },
          { icon: 'mdi-bank-transfer', label: "Transaction", to: { name: "Transaction" } },
        ]
      }
    }
  }
</script>