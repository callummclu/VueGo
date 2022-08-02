<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<script lang="ts">
import Navbar from "./components/Navbar.vue";

export default {
    components: {
        Navbar
    }
}
</script>

<template>
  <Navbar/>
  <RouterView />
</template>

<style scoped>

</style>
