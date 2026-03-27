import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '@/pages/LandingPage.vue'
import GuestInvoicePage from '@/pages/GuestInvoicePage.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'landing',
      component: LandingPage,
    },
    {
      path: '/invoice/new',
      name: 'guest-invoice',
      component: GuestInvoicePage,
    },
  ],
})

export default router
