<template>
  <v-card class="me-8 wide" elevation="2" tile>
    <v-card-title>CPU</v-card-title>
    <v-card-subtitle>CPU Information</v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <v-list dense>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Model
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ cpuInfo.modelName }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Total physical cores
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ cpuInfo.physicalCores }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Total logical cores
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ cpuInfo.logicalCores }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Cache Size
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ cpuInfo.cacheSize }} MB
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Mhz
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ cpuInfo.mhz }}
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Architecture
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ cpuInfo.architecture }}
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
  name: "CPUInfoCard",
  props: {
    title: String,
    subtitle: String,
    cpuThreadCount: Number,
  },
  computed: {
    ...mapState('hwInfo', ['cpuInfo']),
  },
}
</script>

<style scoped>

</style>
