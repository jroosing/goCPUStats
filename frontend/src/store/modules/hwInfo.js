export const SET_HW_STATS = 'SET_HW_STATS';
export const SET_HW_USAGES = 'SET_HW_USAGES';

const getDefaultState = () => {
    return {
        hostInfo: {},
        cpuInfo: {},
        cpuThreadCount: 0,
        os: "",
        arch: "",
        cpuThreadUsages: [{data: [0]}],

        cpuAvg: 0,
        mem: {
            usedPercent: 0,
            total: 0,
            used: 0,
        },
        swap: {
            usedPercent: 0,
            total: 0,
            used: 0,
        },
        temp: [
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0},
            {sensorTemperature: 0}
        ]
    };
};

const state = getDefaultState();

const mutations = {
    [SET_HW_STATS] (state, payload) {
        state.cpuInfo = payload.cpuInfo[0];
        state.cpuThreadCount = payload.cpuThreadCount;
        state.os = payload.operatingSystem;
        state.arch = payload.architecture;
        state.hostInfo = payload.hostInfo;
    },
    [SET_HW_USAGES] (state, payload) {
        state.cpuAvg = payload.cpuAvg;
        state.mem = payload.mem;
        state.swap = payload.swap;
        state.temp = payload.temp;

        let series = state.cpuThreadUsages;
        state.cpuThreadUsages = [{data: [0]}];
        payload.cpuThreadUsages.forEach((core, index) => {
            if (series[index] === undefined || (index === 0 && series[index].name !== "Thread 1")) {
                series[index] = { name: 'Thread ' + (index + 1), data: [] };
            }
            if (series.length > 120) {
                series = [{ data: [0] }];
            }
            series[index].data.push(core);
        });
        state.cpuThreadUsages = series;
    },
};

const actions = {
    updateHardwareUsages({commit}, payload) {
        commit(SET_HW_USAGES, payload);
    },
    async getHardwareInfo(context) {
        try {
            const result = await window.backend.Stats.GetHardwareInfo();
            context.commit(SET_HW_STATS, result);
        } catch (e) {
            return e;
        }
    }
};

export default {
    namespaced: true,
    state,
    getters: {},
    actions,
    mutations,
}
