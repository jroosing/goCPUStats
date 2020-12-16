<template>
  <v-card class="me-8" elevation="2" tile>
    <v-card-title>{{ title }}</v-card-title>
    <v-card-subtitle>CPU Usage over all cores / threads ({{ threads }} in total)</v-card-subtitle>
    <v-divider></v-divider>
    <v-card-text>
      <apexchart type="line" height="350" ref="chart" :options="chartOptions" :series="chartSeries"></apexchart>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  name: "CPUPerCore",
  components: {},
  props: {
    threads: Number,
    title: String,
    series: Array,
  },
  data: () => {
    return {
      chartSeries: [{data: [0]}],
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
  methods: {
    updateSeries(series) {
      this.chartSeries = series;
      this.$refs.chart.updateSeries(this.chartSeries);
    }
  }
}
</script>

<style scoped>

</style>
