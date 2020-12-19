<template>
  <v-card class="me-8" elevation="2" tile>
    <v-card-title>{{ title }}</v-card-title>
    <v-card-subtitle>CPU Usage over all cores / threads ({{ logicalCores }} in total)</v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <apexchart type="line" height="350" ref="chart" :options="chartOptions" :series="series"></apexchart>
    </v-card-text>
  </v-card>
</template>

<script>
import {mapState} from "vuex";

export default {
  name: "CPUThreadUsagesChart",
  components: {},
  props: {
    logicalCores: Number,
    title: String,
  },
  computed: {
    ...mapState('hwInfo', ['cpuLoad']),
  },
  data: () => {
    return {
      series: [{data: [0]}],
      chartOptions: {
        chart: {
          id: 'realtime',
          type: 'line',
          animations: {
            enabled: false,
          },
          toolbar: {
            show: false
          },
          zoom: {
            enabled: false
          }
        },
        dataLabels: {
          enabled: false
        },
        stroke: {
          curve: 'smooth'
        },
        tooltip: {
          enabled: false,
        },
        markers: {
          size: 0
        },
        xaxis: {
          type: "numeric",
          range: 60,
          labels: {
            show: false,
          },
        },
        yaxis: {
          max: 100,
        },
        legend: {
          show: true
        },
      },
    }
  },
  watch: {
    'cpuLoad.usedPerLogicalCorePercentList': {
      deep: true,
      handler(newValues) {
        let seriesState = this.series;
        newValues.forEach((core, index) => {
            if (seriesState[index] === undefined || (index === 0 && seriesState[index].name !== "Thread 1")) {
              seriesState[index] = { name: 'Thread ' + (index + 1), data: [] };
            }
            if (seriesState.length > 120) {
              seriesState = [{ data: [0] }];
            }
          seriesState[index].data.push(core);
        });
        this.series = seriesState;
        this.$refs.chart.updateSeries(this.series);
      }
    }
  }
}
</script>

<style scoped>

</style>
