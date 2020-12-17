export const SET_HW_STATS = 'SET_HW_STATS';
export const SET_HW_USAGES = 'SET_HW_USAGES';

const getDefaultState = () => {
    return {
        cpuInfo: {},
        hostInfo: {},

        cpuLoad: {
            usedPercent: 0,
            usedPerLogicalCorePercent: [],
        },
        memoryLoad: {
            total: 0,
            used: 0,
            usedPercent: 0,
        },
        swapMemoryLoad: {
            total: 0,
            used: 0,
            usedPercent: 0,
        },

        temps: {},
    };
};

const state = getDefaultState();

const mutations = {
    [SET_HW_STATS] (state, payload) {
        state.cpuInfo = payload.cpu;
        state.hostInfo = payload.host;
    },
    [SET_HW_USAGES] (state, payload) {
        state.cpuLoad = payload.cpuLoad;
        state.memoryLoad = payload.memLoad;
        state.swapMemoryLoad = payload.swapMemLoad;
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
