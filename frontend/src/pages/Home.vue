<template>
  <v-container>
    <v-row>
      <v-col>
        <CPUUsageChart title="CPU Usage (avg)" height="250" :physical-cores="cpuInfo.physicalCores" :logical-cores="cpuInfo.logicalCores" :percentage="cpuLoad.usedPercent" />
      </v-col>
      <v-col>
        <MemoryUsageChart title="Memory Usage" height="250" :used="memoryLoad.used" :total="memoryLoad.total" :percentage="memoryLoad.usedPercent" />
      </v-col>
      <v-col>
        <MemoryUsageChart title="Swap Memory Usage" height="250" :used="swapMemoryLoad.used" :total="swapMemoryLoad.total" :percentage="swapMemoryLoad.usedPercent" />
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <CPUThreadUsagesChart title="CPU Usage" :logical-cores="cpuInfo.logicalCores" height="250"/>
      </v-col>
    </v-row>
    <v-row>
      <v-col>
        <HostInfoCard title="Host" subtitle="Host information" />
      </v-col>
      <v-col>
        <CPUInfoCard title="CPU" subtitle="CPU Information" />
      </v-col>
      <v-col>
        <TemperatureInfoCard title="Temperatures" subtitle="Hardware temperature information" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import TemperatureInfoCard from "@/components/TemperatureInfoCard";
import CPUInfoCard from "@/components/CPUInfoCard";
import HostInfoCard from "@/components/HostInfoCard";
import CPUThreadUsagesChart from "@/components/CPUThreadUsagesChart";
import MemoryUsageChart from "@/components/MemoryUsageChart";
import CPUUsageChart from "@/components/CPUUsageChart";
import {mapState} from "vuex";
import Wails from "@wailsapp/runtime";

export default {
  name: "Home",
  components: {
    TemperatureInfoCard,
    CPUInfoCard,
    HostInfoCard,
    CPUThreadUsagesChart,
    MemoryUsageChart,
    CPUUsageChart
  },
  computed: {
    ...mapState('hwInfo', ['cpuInfo', 'cpuLoad', 'memoryLoad', 'swapMemoryLoad']),
  },
  mounted: function () {
    this.$store.dispatch('hwInfo/getHardwareInfo');
    Wails.Events.On("hwUsage", hwUsage => {
      if (hwUsage) {
        this.$store.dispatch('hwInfo/updateHardwareUsages', hwUsage)
      }
    });
  },
}
</script>

<style scoped>

</style>
