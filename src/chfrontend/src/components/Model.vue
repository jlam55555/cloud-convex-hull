<template>
    <h1>Hello, world! from model</h1>

    <ThreeViewer :objUrl="'testing'"/>
</template>

<script lang="ts">
    import {defineComponent} from 'vue'
    import {presignRequest} from "../util/Api";
    import ThreeViewer from "./ThreeViewer.vue";

    export default defineComponent({
        name: 'Model',

        components: {
            ThreeViewer
        },

        // get model
        async created() {
            // get presigned GET URL
            const response = await presignRequest({
                type: 'GET',
                key: this.$route.params.key,
            })
                .then(res => res.json());

            // TODO: load this obj file with THREE.js
            console.log(response);
        },
    })
</script>

<style></style>
