import Vue from 'vue'
import VueRouter from 'vue-router'

import CPUPage from "@/pages/cpu/CPUPage";
import Home from "@/pages/Home";

Vue.use(VueRouter);

const routes = [
    { component: Home, name: 'Home', path: '/'},
    { component: CPUPage, name: 'CPU', path: "/cpu"}
];

const router = new VueRouter({
    mode: 'abstract',
    routes
});

export default router;
