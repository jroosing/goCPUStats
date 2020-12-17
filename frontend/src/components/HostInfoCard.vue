<template>
  <v-card class="me-8" elevation="2" tile>
    <v-card-title>{{ title }}</v-card-title>
    <v-card-subtitle>{{ subtitle }}</v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <v-list dense>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Hostname
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ hostInfo.hostname }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Platform
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ hostInfo.platform }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Platform Family
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ hostInfo.platformFamily }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Platform Version
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ hostInfo.platformVersion }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Kernel Version
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ hostInfo.kernelVersion }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Kernel Architecture
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ hostInfo.kernelArch }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-card-text>
  </v-card>
</template>

<script>
import {mapState} from "vuex";

export default {
  name: "HostInfoCard",
  props: {
    title: String,
    subtitle: String
  },
  computed: {
    ...mapState('hwInfo', ['hostInfo']),
  },
}
</script>

<style scoped>

</style>
