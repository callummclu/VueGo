<script lang="ts">
import { defineComponent } from 'vue';
import { useRoute } from 'vue-router';

export default defineComponent({
    name:"Nav",
    setup() {
        const route = useRoute()

        return route
    }
})
</script>

<template>
    <div class="navbar">
        <h1>V&GES - {{$route.name}}</h1>
    </div>
</template>

<style scoped>
.navbar{
    margin-bottom: 20px;
}
</style>
