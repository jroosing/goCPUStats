<template>
  <v-app id="inspire">
    <v-app-bar app fixed clipped-left>
      <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
      <v-toolbar-title
          style="width: 300px"
          class="ml-0 pl-4"
      >
        <span class="hidden-sm-and-down">Hardware Statistics</span>
      </v-toolbar-title>
    </v-app-bar>
    <v-navigation-drawer v-model="drawer" absolute temporary>
      <v-list nav dense>
        <v-list-item-group>
          <v-list-item to="/">
            <v-list-item-title>
              Home
            </v-list-item-title>
          </v-list-item>
          <v-list-item to="/cpu">
            <v-list-item-title>
              CPU
            </v-list-item-title>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-navigation-drawer>
    <v-main>
      <router-view></router-view>
    </v-main>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; Joey Roosing</span>
    </v-footer>
  </v-app>
</template>

<script>
import "./assets/css/main.css";
export default {
  name: "app",
  data: () => {
    return {
      drawer: false,
    }
  },
};
</script>
