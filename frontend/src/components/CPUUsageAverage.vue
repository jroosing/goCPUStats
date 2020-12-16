<template>
  <v-card class="me-8" elevation="2" tile>
    <v-card-title>{{ title }}</v-card-title>
    <v-card-subtitle>On {{ cores }} core(s) and {{ threads }} thread(s)</v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <SemiRadialChart :percentage="[percentage]"/>
    </v-card-text>
  </v-card>
</template>

<script>
import SemiRadialChart from "@/components/SemiRadialChart";

export default {
  name: "CPUUsageAverage",
  components: {
    SemiRadialChart,
  },
  props: {
    title: String,
    cores: Number,
    threads: Number,
    percentage: Number,
  }
}
</script>

<style scoped>

</style>
