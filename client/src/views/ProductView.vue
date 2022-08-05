<script lang="ts">
import { defineComponent } from "vue";

  export default defineComponent({
    data(){
      return {
        data:[]
      }
    },
    methods:{
      addItemToCart: function(product:any){
        fetch(`http://localhost:3000/basket/${localStorage.getItem("basketId")}`,{
          method:"POST",
          body:JSON.stringify(product)
        }).then(async (res:any) => {
          let res_json = await res.json()
          console.log(res_json)
          this.$emit('changeBasketData',res_json.data)
          window.location.reload()
        })
      }
    }
    ,
    mounted() {
      fetch("http://localhost:3000/product").then(async (res:any) => {
        let res_json = await res.json()
        this.data = res_json.data
      })
    }
  })
</script>

<template>
  <main>
    <div v-if="data.length>0">
    <div v-for="product in data" :key="(product as any).id">
    <van-card
        :num="1"
          :price="(product as any).price"
          currency="£"
          :title="(product as any).productName"
          desc="Description"
          centered
          thumb="123"
        >
        <template #footer>
          <van-button v-on:click="addItemToCart(product)" size="mini">Add to Cart</van-button>
        </template>
        </van-card>

    </div>
    </div>
    <div style="margin:50px; text-align:center;" v-else>
      <p>no products avaliable</p>
    </div>
  </main>
</template>
