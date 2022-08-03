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
    },
    methods : {
        basketRedirect: function() {
            window.location.href = window.location.origin + "/basket"
        },
        homeRedirect: function() {
            window.location.href = window.location.origin
        }
    }
})
</script>


<template>
    <van-nav-bar 
        :title="'VueGo - ' + title.name" 
    />
</template>

<style scoped>
.navbar{
    margin-bottom: 20px;
    display: flex;
    justify-content: space-between;
}
</style>
