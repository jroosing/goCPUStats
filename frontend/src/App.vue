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
            <v-list-item-icon>
              <v-icon>
                mdi-home
              </v-icon>
            </v-list-item-icon>
            <v-list-item-title>
              Home
            </v-list-item-title>
          </v-list-item>
          <v-list-item to="/cpu">
            <v-list-item-icon>
              <v-icon v-if="is64">
                mdi-cpu-64-bit
              </v-icon>
              <v-icon v-else>
                mdi-cpu-32-bit
              </v-icon>
            </v-list-item-icon>
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
import { mapGetters } from "vuex";
import Wails from "@wailsapp/runtime";

export default {
  name: "app",
  data: () => {
    return {
      drawer: false,
    }
  },
  computed: {
    ...mapGetters('hwInfo', ['cpuArch']),
    is64: function() {
      console.log(this.cpuArch);
      if (this.cpuArch === undefined) {
        return true;
      } else if (this.cpuArch.includes("64")) {
        return true;
      } else {
        return false;
      }
    }
  },
  mounted: function () {
    this.$store.dispatch('hwInfo/getHardwareInfo');
    Wails.Events.On("hwUsage", hwUsage => {
      if (hwUsage) {
        this.$store.dispatch('hwInfo/updateHardwareUsages', hwUsage)
      }
    });
  },
};
</script>
