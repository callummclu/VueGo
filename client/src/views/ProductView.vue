<script lang="ts">
  export default {
    data(){
      return {
        data:[]
      }
    },
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
    <div style="display:flex; align-items:center; justify-content: space-between;" v-for="product in data" :key="product.id">
      <div style="display:flex; align-items:center;gap:25px;">
        <p>£{{ product.price }}</p>
      <h3>{{ product.productName.length>0 ? product.productName : "undefined" }}</h3>
      </div>
      <button>Add to Cart</button>
    </div>
  </main>
</template>
