import { HardwareInfo, HardwareLoad } from '@/proto/sys_pb';
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
        temperatures: {
            cpu: {
                temperature: 0
            },
            disk: {
                temperature: 0
            },
            gpu: {
                temperature: 0
            },
            memory: {
                temperature: 0
            }
        },
        cpuArch: "amd64",
    };
};

const state = getDefaultState();

const mutations = {
    [SET_HW_STATS] (state, payload) {
        try {
            let hwInfo = HardwareInfo.deserializeBinary(payload).toObject();
            state.cpuInfo = hwInfo.cpu;
            state.cpuArch = hwInfo.cpu.architecture;
            state.hostInfo = hwInfo.host;
        } catch (err) {
            console.error(err);
        }
    },
    [SET_HW_USAGES] (state, payload) {
        try {
            let hwLoad = HardwareLoad.deserializeBinary(payload).toObject();
            state.cpuLoad = hwLoad.cpu;
            state.memoryLoad = hwLoad.mem;
            state.swapMemoryLoad = hwLoad.swapMem;
            state.temperatures = hwLoad.temps;
        } catch (err) {
            console.error(err);
        }
    },
};

const actions = {
    updateHardwareUsages({commit}, payload) {
        commit(SET_HW_USAGES, payload);
    },
    async getHardwareInfo(context) {
        try {
            const result = await window.backend.Stats.GetHardwareInfoProto();
            context.commit(SET_HW_STATS, result);
        } catch (e) {
            return e;
        }
    }
};

export default {
    namespaced: true,
    state,
    getters: {
        cpuArch: state => state.cpuArch
    },
    actions,
    mutations,
}
