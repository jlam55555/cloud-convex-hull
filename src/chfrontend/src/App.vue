<template>
    <NavBar/>

    <router-view/>
</template>

<script lang="ts">
    import {defineComponent} from 'vue'
    import NavBar from "./components/NavBar.vue";
    import Dashboard from "./components/Dashboard.vue";
    import ThreeViewer from "./components/ThreeViewer.vue";

    export default defineComponent({
        name: 'App',
        components: {
            NavBar,
            Dashboard,
            ThreeViewer,
        }
    })
</script>

<style>
    html, body {
        padding: 0;
        margin: 0;
    }
</style>
