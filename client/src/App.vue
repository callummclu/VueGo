<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<script lang="ts">
import Navbar from "./components/Navbar.vue";

export default {
    data(){
      return {
        basketData:[],
        basketLength: "0"
      }
    },
    components: {
        Navbar
    },
    methods:{
      checkBasketExists: function() {
        let basketId = localStorage.getItem("basketId")

        if(basketId){

          fetch(`http://localhost:3000/basket/${basketId}`)
            .then(async (res:any) => {
              let res_json = await res.json()
              if(res_json.data === "undefined"){
                this.createNewBasket()
              } else {
                this.basketData = res_json.data
                this.basketLength = ((this.basketData as any).items).length.toString()
              }
            })
        } else {
          this.createNewBasket()
        }
      },
      createNewBasket: function() {
        fetch("http://localhost:3000/basket",{
          method:"POST",
          body:JSON.stringify(
            {
              "items":[]
            }
          )
        }).then(async (res:any) => {
          let res_json = await res.json()
          localStorage.setItem("basketId",res_json.data.InsertedID)
        })
      }
    },
    created(){
      this.checkBasketExists()
    }
}
</script>

<template>

  <Navbar :items-number="basketLength"/>
  <RouterView />
</template>

<style scoped>

</style>
