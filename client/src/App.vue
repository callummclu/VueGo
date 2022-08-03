<script setup lang="ts">
import { RouterView } from 'vue-router'
const active = ref(0);
</script>

<script lang="ts">
import Nav from "./components/Navbar.vue";
import { ref } from 'vue';

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
      },
      redirect: function(page:any){
        window.location.replace(window.location.origin+"/"+page)
      }
    },
    created(){
      this.checkBasketExists()
    }
}
</script>

<template>
  <Nav/>
  <RouterView />
  <van-tabbar v-model="active">
    <van-tabbar-item @click="redirect('')" icon="home-o">Home</van-tabbar-item>
    <van-tabbar-item @click="redirect('products')" icon="coupon-o">Products</van-tabbar-item>
    <van-tabbar-item @click="redirect('basket')" icon="cart-o" :badge="basketLength">Basket</van-tabbar-item>
    
  </van-tabbar>
  
</template>

<style scoped>

</style>
