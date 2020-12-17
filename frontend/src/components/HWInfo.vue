<template>
  <v-app>
    <v-container>
      <v-row>
        <v-col>
          <CPUUsageAverage title="CPU Usage (avg)" :cores="cpuInfo.cores" :threads="cpuThreadCount" :percentage="cpuAvg" />
        </v-col>
        <v-col>
          <MemoryUsage title="Memory Usage" :used="mem.used" :total="mem.total" :percentage="mem.usedPercent" />
        </v-col>
        <v-col>
          <MemoryUsage title="Swap Memory Usage" :used="swap.used" :total="swap.total" :percentage="swap.usedPercent" />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <CPUPerCore title="CPU Usage" ref="CPUPerCoreRef" :threads="cpuThreadCount"/>
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
import CPUUsageAverage from "@/components/CPUUsageAverage";
import Wails from '@wailsapp/runtime';
import {mapState} from "vuex";
import MemoryUsage from "@/components/MemoryUsage";
import CPUPerCore from "@/components/CPUPerCore";
import HostInfoCard from "@/components/HostInfoCard";
import CPUInfoCard from "@/components/CPUInfoCard";
import TemperatureInfoCard from "@/components/TemperatureInfoCard";

export default {
  components: {
    TemperatureInfoCard,
    CPUInfoCard,
    HostInfoCard,
    CPUPerCore,
    MemoryUsage,
    CPUUsageAverage
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
