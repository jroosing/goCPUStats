<template>
  <v-card class="me-8" elevation="2" tile>
    <v-card-title>{{ title }}</v-card-title>
    <v-card-subtitle>{{ memUsed }} / {{ memTotal }}
    </v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <!-- Most operating systems calculate gigabyte as 1073741824 = 1024*1024*1024 -->
      <SemiRadialChart :percentage="[percentage.toPrecision(2)]"/>
    </v-card-text>
  </v-card>
</template>

<script>
import SemiRadialChart from "@/components/SemiRadialChart";
import humanByteSize from "@/util/byte-utils";

export default {
  name: "MemoryUsage",
  components: {
    SemiRadialChart,
  },
  props: {
    title: String,
    percentage: Number,
    used: Number,
    total: Number
  },
  data: () => {
    return {
      memUsed: '0',
      memTotal: '0',
    }
  },
  watch: {
    used: function(newVal) {
      this.memUsed = humanByteSize(newVal, false, 2, true);
    },
    total: function(newVal) {
      this.memTotal = humanByteSize(newVal, false, 2, true);
    }
  }
}
</script>

<style scoped>

</style>
