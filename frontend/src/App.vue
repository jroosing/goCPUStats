<template>
  <v-app id="inspire">
    <v-app-bar app fixed clipped-left>
      <v-toolbar-title
          style="width: 300px"
          class="ml-0 pl-4"
      >
        <span class="hidden-sm-and-down">Joey</span>
      </v-toolbar-title>
    </v-app-bar>
    <v-main>
      <HWInfo />
    </v-main>
    <v-footer app fixed>
      <span style="margin-left:1em">&copy; Joey Roosing</span>
    </v-footer>
  </v-app>
</template>

<script>
import HWInfo from "./components/HWInfo.vue";

import "./assets/css/main.css";
export default {
  name: "app",
  components: {
    HWInfo
  }
};
</script>
