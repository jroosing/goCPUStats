export const SET_HW_STATS = 'SET_HW_STATS';
export const SET_HW_USAGES = 'SET_HW_USAGES';

const getDefaultState = () => {
    return {
        hostInfo: {},
        cpuInfo: {},
        coreCount: 0,
        os: "",
        arch: "",
        perCPUUsage: [{data: [0]}],

        usages: {
            CPUAverage: 0,
            PerCPUUsage: [],
            Mem: {
                usedPercent: 0,
                total: 0,
                used: 0,
            },
            Swap: {
                usedPercent: 0,
                total: 0,
                used: 0,
            },
            Temp: [
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
        },
    };
};

const state = getDefaultState();

const mutations = {
    [SET_HW_STATS] (state, payload) {
        const { OS, Arch, CoreCount, CPUInfo, HostInfo } = payload;
        if (CPUInfo.length > 0) {
            state.cpuInfo = CPUInfo[0];
        }
        state.coreCount = CoreCount;
        state.os = OS;
        state.arch = Arch;
        state.hostInfo = HostInfo;
    },
    [SET_HW_USAGES] (state, payload) {
        state.usages = payload;

        let series = state.perCPUUsage;
        state.perCPUUsage = [{data: [0]}];
        payload.PerCPUUsage.forEach((core, index) => {
            if (series[index] === undefined || (index === 0 && series[index].name !== "Thread 1")) {
                series[index] = { name: 'Thread ' + (index + 1), data: [] };
            }
            if (series.length > 120) {
                series = [{ data: [0] }];
            }
            series[index].data.push(core);
        });
        state.perCPUUsage = series;
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
