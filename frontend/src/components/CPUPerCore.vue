<template>
  <v-card class="me-8" elevation="2" tile>
    <v-card-title>CPU Usage</v-card-title>
    <v-card-subtitle>CPU Usage over all cores / threads ({{coreCount}} in total)</v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <apexchart type="line" height="350" ref="chart" :options="chartOptions" :series="series"></apexchart>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "CPUPerCore"
}
</script>

<style scoped>

</style>
