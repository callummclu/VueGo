import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HomeView
    },
    {
      path: '/products',
      name: 'Products',
      component: () => import('../views/ProductView.vue')
    },
    {
      path: '/basket',
      name: 'Basket',
      component: () => import('../views/BasketView.vue')
    },
    {
      path: '/order/:id',
      name:'order',
      component: () => import('../views/CheckoutView.vue')
    }
  ]
})

export default router
