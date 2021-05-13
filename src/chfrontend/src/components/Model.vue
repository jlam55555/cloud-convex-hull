<template>
    Viewing model: {{$route.params.key || 'None'}}

    <ThreeViewer :objUrl="objUrl"/>

    <button @click="downloadObjFile()">Download .obj file</button>

    <button @click="calculateConvexHull()">Calculate convex hull</button>
</template>

<script lang="ts">
    import {defineComponent} from 'vue'
    import {convexHullRequest, presignRequest} from "../util/Api";
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
                    key: <string>this.$route.params.key + '.obj',
                }).then(res => res.json())).url;
            },

            async downloadObjFile() {
                // cannot use a static link element because the presigned links
                // expire, so have to manually dynamically generate one and
                // click it; see https://stackoverflow.com/a/49917066/2397327
                const a = document.createElement('a');
                a.href = await this.getObjUrl();
                a.download = <string>a.href.split('/').pop();

                document.body.appendChild(a);
                a.click();
                document.body.removeChild(a);
            },

            async calculateConvexHull() {
                const {key} = await convexHullRequest(this.$route.params.key)
                    .then(res => res.json());

                this.$router.push({
                    path: '/model/' + key.split('.')[0]
                });
            },
        },

        watch: {
            async $route(newVal, oldVal) {
                this.objUrl = await this.getObjUrl();
            },
        },

        // get model
        async created() {
            // load this file with three.js
            this.objUrl = await this.getObjUrl();
        },
    })
</script>

<style></style>
