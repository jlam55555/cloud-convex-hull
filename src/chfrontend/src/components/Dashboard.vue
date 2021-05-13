<template>
    <h1>Upload a file to begin</h1>
    <form>
        <input type="file" placeholder="Model file" @change="onFileChange($event)">
        <button @click="uploadFile($event)">Upload file</button>
    </form>
</template>

<script lang="ts">
    import {defineComponent} from 'vue';
    import {presignRequest, putRequest} from "../util/Api";

    export default defineComponent({
        name: 'Dashboard',
        setup() { },
        data() {
            return {
                file: null
            };
        },
        methods: {
            onFileChange(evt: any) {
                this.file = evt.target.files[0];
            },
            async uploadFile(evt: Event) {
                evt.preventDefault();

                if (!this.file) {
                    console.error('Must input a file to upload');
                    return;
                }

                // need to put to presigned url
                const {key, url: putUrl} = await presignRequest({type: 'PUT'})
                    .then(res => res.json());

                // make regular put request
                await putRequest(this.file, putUrl);

                // redirect to model page after completion
                this.$router.push({path: '/model/' + key.split('.')[0]});
            },
        },
    });
</script>

<style scoped>

</style>
