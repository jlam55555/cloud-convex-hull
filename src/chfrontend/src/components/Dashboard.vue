<template>
    <h1>This is the dashboard.</h1>

    <p>TODO: list of a user's projects</p>

    <p>TODO: upload a file</p>
</template>

<script lang="ts">
    import {defineComponent} from 'vue';
    import {testRequest} from "../util/Api";

    export default defineComponent({
        name: 'Dashboard',
        setup: () => {
            console.log("Start");
            testRequest();
            console.log("End");
        }
    });
</script>

<style scoped>

</style>
