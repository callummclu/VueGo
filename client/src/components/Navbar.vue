<script lang="ts">
import { defineComponent } from 'vue';
import { useRoute } from 'vue-router';

export default defineComponent({
    name:"Nav",
    data(){
      const route = useRoute()
      return {
        title:route
      }
    }
})
</script>


<template>
    <van-nav-bar 
        :title="'VueGo - ' + title.name?.toString()" 
    />
</template>