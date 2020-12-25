import 'core-js/stable';
import 'regenerator-runtime/runtime';
import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/dist/vuetify.min.css';

import router from './router';

import Vue from 'vue';
import Vuetify from 'vuetify';
import VueApexCharts from 'vue-apexcharts'
import store from "./store/store";

Vue.use(Vuetify);
Vue.use(VueApexCharts)
Vue.component('apexchart', VueApexCharts)

import App from './App.vue';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		store,
		vuetify: new Vuetify({
			icons: {
				iconfont: 'mdi'
			},
			theme: {
				dark: true
			}
		}),
		render: h => h(App),
		router,
		mounted() {
			this.$router.replace('/')
		},
	}).$mount("#app");
});
