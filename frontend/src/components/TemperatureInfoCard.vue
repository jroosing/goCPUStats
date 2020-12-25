<template>
  <v-card class="me-8 wide" elevation="2" tile>
    <v-card-title>Temperatures</v-card-title>
    <v-card-subtitle>Hardware temperature Information</v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <v-list dense>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              CPU
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ temperatures.cpu.temperature }}&#8451;
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              GPU
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ temperatures.gpu.temperature }}&#8451;
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Memory
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ temperatures.memory.temperature }}&#8451;
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Disk
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ temperatures.disk.temperature }}&#8451;
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
            </v-list-item-title>
            <v-list-item-subtitle>
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
            </v-list-item-title>
            <v-list-item-subtitle>
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
  name: "TemperatureInfoCard",
  props: {
    title: String,
    subtitle: String
  },
  computed: {
    ...mapState('hwInfo', ['temperatures']),
  },
}
</script>

<style scoped>
.v-card {
  height: 100%;
}
</style>
