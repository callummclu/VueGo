<template>
<van-notice-bar
  :scrollable="false"
  text="This page is currently under construction."
/>
<p style="padding:20px">Featured Items</p>
<van-grid :column-num="3">
  <van-grid-item v-for="value in 6" :key="value" icon="photo-o" text="Text" />
</van-grid>

</template>
