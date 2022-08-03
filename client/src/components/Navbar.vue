<script lang="ts">
import { defineComponent } from 'vue';
import { useRoute } from 'vue-router';

export default defineComponent({
    name:"Nav",
    props:["itemsNumber"],
    setup() {
        const route = useRoute()
        return route
    },
})
</script>


<template>
    <div class="navbar">
        <h1>V&GES - {{$route.name}}</h1>
        <p><a href="basket">basket <span v-if="true">({{itemsNumber}})</span></a></p>
    </div>
</template>

<style scoped>
.navbar{
    margin-bottom: 20px;
    display: flex;
    justify-content: space-between;
}
</style>
