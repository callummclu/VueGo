<script lang="ts">
  export default {
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
  }
</script>

<template>
  <main>
    <div v-if="data.length>0">
    <div v-for="product in data" :key="product.id">
    <van-card
        :num="1"
          :price="product.price"
          currency="£"
          :title="product.productName"
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
