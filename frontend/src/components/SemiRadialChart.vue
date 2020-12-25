<template>
  <v-container>
    <apexchart type="radialBar" :height="height" :options="chartOptions" :series="percentage"></apexchart>
  </v-container>
</template>

<script>
export default {
  name: "SemiRadialChart.vue",
  props: {
    percentage: Array,
    height: {
      type: String,
      default: function() {
        return "250";
      }
    }
  },
  data: () => ({
    chartOptions: {
      chart: {
        sparkline: {
          enabled: true
        },
        type: 'radialBar',
      },
      plotOptions: {
        radialBar: {
          startAngle: -90,
          endAngle: 90,
          track: {
            background: "#e7e7e7",
            strokeWidth: '97%',
          },
          dataLabels: {
            name: {
              show: false
            },
            value: {
              offsetY: -2,
              fontSize: '22px'
            }
          }
        },
      },
      grid: {
        padding: {
          top: -10
        }
      },
    }
  }),
}
</script>

<style scoped>

</style>
