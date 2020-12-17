<template>
  <v-app>
    <v-container>
      <v-row>
        <v-col>
          <CPUUsageChart title="CPU Usage (avg)" :physical-cores="cpuInfo.physicalCores" :logical-cores="cpuInfo.logicalCores" :percentage="cpuLoad.usedPercent" />
        </v-col>
        <v-col>
          <MemoryUsageChart title="Memory Usage" :used="memoryLoad.used" :total="memoryLoad.total" :percentage="memoryLoad.usedPercent" />
        </v-col>
        <v-col>
          <MemoryUsageChart title="Swap Memory Usage" :used="swapMemoryLoad.used" :total="swapMemoryLoad.total" :percentage="swapMemoryLoad.usedPercent" />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <CPUThreadUsagesChart title="CPU Usage" :logical-cores="cpuInfo.logicalCores" />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <HostInfoCard title="Host" subtitle="Host information" />
        </v-col>
        <v-col>
          <CPUInfoCard title="CPU" subtitle="CPU Information" />
        </v-col>
<!--        <v-col>-->
<!--          <TemperatureInfoCard title="Temperatures" subtitle="Hardware temperature information" />-->
<!--        </v-col>-->
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import CPUUsageChart from "@/components/CPUUsageChart";
import Wails from '@wailsapp/runtime';
import {mapState} from "vuex";
import MemoryUsageChart from "@/components/MemoryUsageChart";
import CPUThreadUsagesChart from "@/components/CPUThreadUsagesChart";
import HostInfoCard from "@/components/HostInfoCard";
import CPUInfoCard from "@/components/CPUInfoCard";
//import TemperatureInfoCard from "@/components/TemperatureInfoCard";

export default {
  components: {
    //TemperatureInfoCard,
    CPUInfoCard,
    HostInfoCard,
    CPUThreadUsagesChart,
    MemoryUsageChart,
    CPUUsageChart
  },
  data: () => {
    return {}
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
};
</script>
