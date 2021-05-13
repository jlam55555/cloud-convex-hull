<template>
    <h1>This is the dashboard.</h1>

    <p>TODO: list of a user's projects</p>

    <p>TODO: upload a file</p>

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

                const {key, url: putUrl} = await presignRequest({type: 'PUT'})
                    .then(res => res.json());

                await putRequest(this.file, putUrl);

                this.$router.push({path: '/model/' + key});
            },
        },
    });
</script>

<style scoped>

</style>
