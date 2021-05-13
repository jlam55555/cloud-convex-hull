<template>
    <div ref="threeContainer"></div>
</template>

<script lang="ts">
    import {defineComponent} from 'vue'
    import * as THREE from 'three';

    // despite this weirdness, this is the canonical way:
    // https://www.npmjs.com/package/three-orbitcontrols
    import {OrbitControls} from "three/examples/jsm/controls/OrbitControls";

    export default defineComponent({
        name: 'Model',

        props: {
            objUrl: {
                type: String,
                required: true,
            },
        },

        mounted() {
            console.log('Got objUrl: ' + this.objUrl);

            // ref: https://threejs.org/docs/#manual/en
            const scene = new THREE.Scene();
            const camera = new THREE.PerspectiveCamera(75,
                window.innerWidth/window.innerHeight, 0.1, 1000);

            const renderer = new THREE.WebGLRenderer();
            renderer.setSize(window.innerWidth, window.innerHeight);
            (<HTMLElement>this.$refs.threeContainer).appendChild(renderer.domElement);

            const geometry = new THREE.BoxGeometry();
            const material = new THREE.MeshBasicMaterial({ color: 0x00ff00 });
            const cube = new THREE.Mesh(geometry, material);
            scene.add(cube);

            camera.position.z = 5;

            const controls = new OrbitControls(camera, renderer.domElement);

            function animate() {
                requestAnimationFrame(animate);
                renderer.render(scene, camera);
            }
            animate();
        }
    })
</script>

<style></style>
