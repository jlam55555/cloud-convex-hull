import { createApp } from 'vue'
import App from './App.vue'
import Dashboard from "./components/Dashboard.vue";
import Model from "./components/Model.vue";
import {createRouter, createWebHistory} from "vue-router";
import PageNotFound from "./components/PageNotFound.vue";

const routes = [
    {path: '/', component: Dashboard},
    {path: '/model/:key', component: Model},

    // see: https://stackoverflow.com/a/40194152/2397327
    {path: '/:pathMatch(.*)*', component: PageNotFound},
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

const app = createApp(App);
app.use(router);
app.mount('#app');
