<template>
    <h1>Hello, world! from model</h1>

    <ThreeViewer :objUrl="objUrl"/>

    <button @click="downloadObjFile()">Download .obj file</button>
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
                objUrl: '',
            };
        },

        methods: {
            async getObjUrl() {
                return (await presignRequest({
                    type: 'GET',
                    key: <string>this.$route.params.key,
                }).then(res => res.json())).url;
            },

            async downloadObjFile() {
                // cannot use a static link element because the presigned links
                // expire, so have to manually dynamically generate one and
                // click it; see https://stackoverflow.com/a/49917066/2397327
                const a = document.createElement('a');
                a.href = await this.getObjUrl();
                a.download = a.href.split('/').pop() + '.obj';

                document.body.appendChild(a);
                a.click();
                document.body.removeChild(a);
            }
        },

        // get model
        async created() {
            // load this file with three.js
            this.objUrl = await this.getObjUrl();
        },
    })
</script>

<style></style>
