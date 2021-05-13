<template>
    <NavBar/>

    <router-view/>
</template>

<script lang="ts">
    import {defineComponent} from 'vue'
    import NavBar from "./components/NavBar.vue";
    import Dashboard from "./components/Dashboard.vue";

    export default defineComponent({
        name: 'App',
        components: {
            NavBar,
            Dashboard
        }
    })
</script>

<style>
    html, body {
        padding: 0;
        margin: 0;
    }
</style>
