<script lang="ts">
  export default {
    data(){
      return {
        data:[]
      }
    },
    mounted() {
      fetch(`http://localhost:3000/basket/${localStorage.getItem("basketId")}`).then(async (res:any) => {
        let res_json = await res.json()
        this.data = res_json.data
      })
    },
    methods:{
      removeFromBasket: function(product:any){
        fetch(`http://localhost:3000/basket/${localStorage.getItem("basketId")}/delete`,{
          method:"POST",
          body:JSON.stringify(product)
        }).then(async (res:any) => {
          let res_json = await res.json()
          console.log(res_json)
        })
      },
      checkout: function(){
        fetch(`http://localhost:3000/checkout/${localStorage.getItem("basketId")}`,{
          method:"POST"
        }).then(async (res:any) => {
          let res_json = await res.json()
          window.location.href = window.location.origin +"/order/"+ res_json.data.InsertedID
        })
        
      }
    }
  }
</script>

<template>

    <div  v-for="product in data.items" :key="product.product">
        <van-card
        :num="1"
          :price="product.price"
          currency="£"
          :title="product.productName"
          :quantity="product.quantity"
          desc="Description"
          centered

          thumb="123"
         >
        <template #footer>
          <van-button color="red" v-on:click="removeFromBasket(product)" size="mini">remove from basket</van-button>
        </template>
        </van-card>
    </div>

    <div style="text-align:center;" v-if="!data.items || data.items.length == 0">
      <p>Basket is empty</p>
    </div>

    <van-submit-bar style="margin-bottom:75px" :price="data.total * 100" button-text="Submit" @submit="checkout" currency="£" label="Total"/>

</template>
