<template>
    <p>{{status}}</p>
    <div ref="threeContainer"></div>
</template>

<script lang="ts">
    import {defineComponent} from 'vue'
    import * as THREE from 'three';

    // despite this weirdness, this is the canonical way:
    // https://www.npmjs.com/package/three-orbitcontrols
    import {OrbitControls} from "three/examples/jsm/controls/OrbitControls";
    import {OBJLoader} from "three/examples/jsm/loaders/OBJLoader";
    import {BufferGeometryUtils} from "three/examples/jsm/utils/BufferGeometryUtils";

    export default defineComponent({
        name: 'Model',

        data() {
            return {
                status: 'Waiting for object...',
                scene: <any>null,
            };
        },

        props: {
            objUrl: {
                type: String,
                required: true,
            },
        },

        watch: {
            // on prop change
            objUrl: function(newVal, oldVal) {
                if (!newVal) {
                    return;
                }

                // load object from obj file
                const loader = new OBJLoader();
                loader.load(newVal, obj => {
                    for (const child of obj.children) {
                        let geometry = (<any>child).geometry;

                        // prevent duplicate vertices from affecting the centering
                        geometry = BufferGeometryUtils.mergeVertices(geometry);

                        // make the mesh in a more convenient location
                        geometry.center();

                        const material = new THREE.MeshLambertMaterial({});
                        const mesh = new THREE.Mesh(geometry, material);

                        this.$data.scene.add(mesh);
                    }
                }, xhr => {
                    this.$data.status = ((xhr.loaded)/(xhr.total)*100).toFixed(2) + '% loaded';
                }, err => {
                    console.error(err);
                });
            }
        },

        mounted() {
            // ref: https://threejs.org/docs/#manual/en
            const scene = new THREE.Scene();
            const camera = new THREE.PerspectiveCamera(75, 1, 0.1, 1000);

            // set scene to state
            this.$data.scene = scene;

            const renderer = new THREE.WebGLRenderer();
            renderer.setSize(500, 500);
            (<HTMLElement>this.$refs.threeContainer).appendChild(renderer.domElement);

            camera.position.set(20, 20, 20);

            // allow user to scroll/pan/rotate with mouse/touch
            const controls = new OrbitControls(camera, renderer.domElement);

            // add lighting
            const light = new THREE.AmbientLight(0x404040);
            scene.add(light);

            const spotlight = new THREE.SpotLight(0xffffff, 1);
            spotlight.position.set(30, 30, 30);
            spotlight.target.position.set(0, 0, 0);
            scene.add(spotlight);

            function animate() {
                requestAnimationFrame(animate);
                renderer.render(scene, camera);
            }
            animate();
        }
    })
</script>

<style></style>
