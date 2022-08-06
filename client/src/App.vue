<script setup lang="ts">
import { ref } from 'vue';

import { RouterView,useRoute } from 'vue-router'
</script>

<script lang="ts">
import Nav from "./components/Navbar.vue";

export default {
    data(){
      return {
        basketData:[],
        basketLength: "0"
      }
    },
    components: {
        Nav
    },
    methods:{
      getBasketQuantity: function(data:any):string {
        let totalItems = 0

        data.items.forEach((item:any) => {
          totalItems += item.quantity
        });

        

        return totalItems.toString()
      },
      checkBasketExists: function() {
        let basketId = localStorage.getItem("basketId")

        if(basketId){

          fetch(`http://localhost:3000/basket/${basketId}`)
            .then(async (res:any) => {
              let res_json = await res.json()

              if('message' in res_json){
                this.createNewBasket()
              }
              
              if('data' in res_json){
                this.basketData = res_json.data
                this.basketLength = this.getBasketQuantity(res_json.data)
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
      },
      redirect: function(page:any){
        window.location.replace(window.location.origin+"/"+page)
      },
      changeBasket: function(data:any){
        console.log("here")
        this.basketData = data.items.length.toString()
      }
    },
    created(){
      this.checkBasketExists()
    },

}
</script>

<template>
  <Nav/>
  <RouterView />
  <van-tabbar>
    <van-tabbar-item name="Home" @click="redirect('')" icon="home-o">Home</van-tabbar-item>
    <van-tabbar-item name="Products" @click="redirect('products')" icon="coupon-o">Products</van-tabbar-item>
    <van-tabbar-item name="Basket" @click="redirect('basket')" icon="cart-o" :badge="basketLength">Basket</van-tabbar-item>
    
  </van-tabbar>
  
</template>

<style scoped>

</style>
