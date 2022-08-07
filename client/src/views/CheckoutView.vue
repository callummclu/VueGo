<script lang="ts">import { routeLocationKey, useRoute } from "vue-router"


  export default {
    data(){
        const route = useRoute()
      return {
        data:[],
        orderId:route.params.id
      }
    },
    mounted() {
      const route = useRoute()
      fetch(`http://localhost:3000/order/${route.params.id}`).then(async (res:any) => {
        let res_json = await res.json()

        this.data = res_json.data
      })
    }
  }
</script>

<template>
  <main>
    <div style="text-align: center;margin-top:50px;margin-bottom: 50px;">
        <h2>Thanks for your order</h2>

        <p>id: {{orderId}}</p>
    </div>
    

    

    <div v-for="item in data.basket?.items">
    <van-card
        :num="item.quantity"
        :price="item.price"
        :title="item.productName"
        currency="£"
        :quantity="item.quantity"
        desc="Description"
        thumb="https://fastly.jsdelivr.net/npm/@vant/assets/ipad.jpeg"
        />
        
    </div>
  </main>
</template>
