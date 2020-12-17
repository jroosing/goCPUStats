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
              {{ temp[4].sensorTemperature }}&#8451;
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              GPU
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ temp[11].sensorTemperature }}&#8451;
            </v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title>
              Memory
            </v-list-item-title>
            <v-list-item-subtitle>
              {{ temp[14].sensorTemperature }}&#8451;
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
    ...mapState('hwInfo', ['temp']),
  },
}
</script>

<style scoped>

</style>
