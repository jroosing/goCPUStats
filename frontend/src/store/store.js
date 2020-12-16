import Vue from 'vue';
import Vuex from 'vuex';

import HWInfo from './modules/hwInfo.js';

Vue.use(Vuex);

export const store = new Vuex.Store({
    strict: false,
    modules: {
        hwInfo: HWInfo
    }
});

export default store;
