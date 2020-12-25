<template>
  <v-container>
    <v-row>
      <v-col>
        <CPUThreadUsagesChart title="CPU Usage" :logical-cores="cpuInfo.logicalCores" height="250"/>
      </v-col>
      <v-col>
        <CPUUsageChart title="CPU Usage (avg)" height="250" :physical-cores="cpuInfo.physicalCores" :logical-cores="cpuInfo.logicalCores" :percentage="cpuLoad.usedPercent" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <CPUInfoCard title="CPU" subtitle="CPU Information" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import CPUThreadUsagesChart from "@/components/CPUThreadUsagesChart";
import CPUUsageChart from "@/components/CPUUsageChart";
import CPUInfoCard from "@/components/CPUInfoCard";

import {mapState} from "vuex";

export default {
  name: "CPUPage",
  components: {
    CPUThreadUsagesChart,
    CPUUsageChart,
    CPUInfoCard,
  },
  computed: {
    ...mapState('hwInfo', ['cpuInfo', 'cpuLoad']),
  },
}
</script>

<style scoped>

</style>
