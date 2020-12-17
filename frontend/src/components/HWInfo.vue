<template>
  <v-app>
    <v-container>
      <v-row>
        <v-col>
          <CPUUsageChart title="CPU Usage (avg)" :cores="cpuInfo.cores" :threads="cpuThreadCount" :percentage="cpuAvg" />
        </v-col>
        <v-col>
          <MemoryUsage title="Memory Usage" :used="mem.used" :total="mem.total" :percentage="mem.usedPercent" />
        </v-col>
        <v-col>
          <MemoryUsageChart title="Swap Memory Usage" :used="swap.used" :total="swap.total" :percentage="swap.usedPercent" />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <CPUThreadUsagesChart title="CPU Usage" ref="CPUPerCoreRef" :threads="cpuThreadCount"/>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <HostInfoCard title="Host" subtitle="Host information" />
        </v-col>
        <v-col>
          <CPUInfoCard title="CPU" subtitle="CPU Information" :threads="cpuThreadCount"/>
        </v-col>
        <v-col>
          <TemperatureInfoCard title="Temperatures" subtitle="Hardware temperature information" />
        </v-col>
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
import TemperatureInfoCard from "@/components/TemperatureInfoCard";

export default {
  components: {
    TemperatureInfoCard,
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
    ...mapState('hwInfo', ['cpuInfo', 'cpuAvg', 'cpuThreadCount','mem', 'swap']),
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
