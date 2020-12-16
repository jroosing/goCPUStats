<template>
  <v-app>
    <v-container>
      <v-row>
        <v-col>
          <v-card class="me-8" elevation="2" tile>
            <v-card-title>CPU Usage (avg)</v-card-title>
            <v-card-subtitle>On {{ cpuInfo.cores }} core(s) and {{ coreCount }} thread(s)</v-card-subtitle>
            <v-divider></v-divider>
            <v-card-text>
              <SemiRadialChart :percentage="[usages.CPUAverage]"/>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col>
          <v-card class="me-8" elevation="2" tile>
            <v-card-title>Memory Usage</v-card-title>
            <v-card-subtitle>{{ (usages.Mem.used / 1073741824).toPrecision(4) }} /
              {{ (usages.Mem.total / 1073741824).toPrecision(4) }} GB
            </v-card-subtitle>
            <v-divider></v-divider>
            <v-card-text>
              <!-- Most operating systems calculate gigabyte as 1073741824 = 1024*1024*1024 -->
              <SemiRadialChart :percentage="[usages.Mem.usedPercent.toPrecision(2)]"/>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col>
          <v-card class="me-8" elevation="2" tile>
            <v-card-title>Swap Memory Usage</v-card-title>
            <v-card-subtitle>{{ (usages.Swap.used / 1073741824).toPrecision(4) }} /
              {{ (usages.Swap.total / 1073741824).toPrecision(4) }} GB
            </v-card-subtitle>
            <v-divider></v-divider>
            <v-card-text>
              <SemiRadialChart :percentage="[usages.Swap.usedPercent.toPrecision(2)]"/>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-card class="me-8" elevation="2" tile>
            <v-card-title>CPU Usage</v-card-title>
            <v-card-subtitle>CPU Usage over all cores / threads ({{coreCount}} in total)</v-card-subtitle>
            <v-divider></v-divider>
            <v-card-text>
              <apexchart type="line" height="350" ref="chart" :options="chartOptions" :series="series"></apexchart>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <v-card class="me-8" elevation="2" tile>
            <v-card-title>Host</v-card-title>
            <v-card-subtitle>Host Information</v-card-subtitle>
            <v-divider></v-divider>
            <v-card-text>
              <v-list dense>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Hostname
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ hostInfo.hostname }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Platform
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ hostInfo.platform }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Platform Family
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ hostInfo.platformFamily }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Platform Version
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ hostInfo.platformVersion }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Kernel Version
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ hostInfo.kernelVersion }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Kernel Architecture
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ hostInfo.kernelArch }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col>
          <v-card class="me-8 wide" elevation="2" tile>
            <v-card-title>CPU</v-card-title>
            <v-card-subtitle>CPU Information</v-card-subtitle>
            <v-divider></v-divider>
            <v-card-text>
              <v-list dense>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Model
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ cpuInfo.modelName }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Total cores
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ cpuInfo.cores }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Total threads
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ coreCount }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Cache Size
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ cpuInfo.cacheSize }} MB
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Operating System
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ os }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Architecture
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ arch }}
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-card-text>
          </v-card>
        </v-col>
        <v-col>
          <v-card class="me-8 wide" elevation="2" tile>
            <v-card-title>Temperatures</v-card-title>
            <v-card-subtitle>Hardware temperature Information</v-card-subtitle>
            <v-divider></v-divider>
            <v-card-text>
              <v-list dense>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      CPU
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ usages.Temp[4].sensorTemperature }}&#8451;
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      GPU
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ usages.Temp[11].sensorTemperature }}&#8451;
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <v-list-item-title>
                      Memory
                    </v-list-item-title>
                    <v-list-item-subtitle>
                      {{ usages.Temp[14].sensorTemperature }}&#8451;
                    </v-list-item-subtitle>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-card-text>
          </v-card>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<script>
import SemiRadialChart from "@/components/SemiRadialChart";
import Wails from '@wailsapp/runtime';
import {mapState} from "vuex";

export default {
  components: {
    SemiRadialChart,
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
            show: true
          },
          zoom: {
            enabled: true
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
  computed: {
    ...mapState('hwInfo', ['hostInfo', 'cpuInfo', 'coreCount', 'os', 'arch', 'usages'])
  },
  mounted: function () {
    this.$store.dispatch('hwInfo/getHardwareInfo');
    Wails.Events.On("hwUsage", hwUsage => {
      if (hwUsage) {
        hwUsage.PerCPUUsage.forEach((core, index) => {
          if (this.series[index] === undefined || (index === 0 && this.series[index].name !== "Thread 1")) {
            this.series[index] = { name: 'Thread ' + (index + 1), data: [] };
          }
          if (this.series.length > 120) {
            this.series = [{ data: [0] }];
          }
          this.series[index].data.push(core);
          this.$refs.chart.updateSeries(this.series);
        });
        this.$store.dispatch('hwInfo/updateHardwareUsages', hwUsage)
      }
    });
  }
};
</script>
