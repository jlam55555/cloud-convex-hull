<template>
    <h1>Hello, world! from model</h1>

    <ThreeViewer :objUrl="objUrl"/>
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

        data() {
            return {
                objUrl: null,
            };
        },

        // get model
        async created() {
            // get presigned GET URL
            const response = await presignRequest({
                type: 'GET',
                key: this.$route.params.key,
            })
                .then(res => res.json());

            // load this file with three.js
            this.objUrl = response.url;
        },
    })
</script>

<style></style>
