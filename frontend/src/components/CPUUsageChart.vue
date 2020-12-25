<template>
  <v-card class="me-8" elevation="2" tile>
    <v-card-title>{{ title }}</v-card-title>
    <v-card-subtitle>On {{ physicalCores }} core(s) and {{ logicalCores }} thread(s)</v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <SemiRadialChart :percentage="[percentage]"/>
    </v-card-text>
  </v-card>
</template>

<script>
import SemiRadialChart from "@/components/SemiRadialChart";

export default {
  name: "CPUUsageChart",
  components: {
    SemiRadialChart,
  },
  props: {
    title: String,
    physicalCores: Number,
    logicalCores: Number,
    percentage: Number,
  }
}
</script>

<style scoped>
  .v-card {
    height: 100%;
  }
</style>
