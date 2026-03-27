import { createRouter, createWebHistory } from 'vue-router'
import LandingPage from '@/pages/LandingPage.vue'
import GuestInvoicePage from '@/pages/GuestInvoicePage.vue'
import LoginPage from '@/pages/LoginPage.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    // Guest routes (no auth required)
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
    {
      path: '/login',
      name: 'login',
      component: LoginPage,
    },

    // Authorized routes (require JWT)
    {
      path: '/dashboard',
      name: 'dashboard',
      component: () => import('@/pages/DashboardPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/companies',
      name: 'companies',
      component: () => import('@/pages/CompanyListPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/companies/new',
      name: 'company-create',
      component: () => import('@/pages/CompanyFormPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/companies/:id/edit',
      name: 'company-edit',
      component: () => import('@/pages/CompanyFormPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/clients',
      name: 'clients',
      component: () => import('@/pages/ClientListPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/clients/new',
      name: 'client-create',
      component: () => import('@/pages/ClientFormPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/clients/:id/edit',
      name: 'client-edit',
      component: () => import('@/pages/ClientFormPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/invoices',
      name: 'invoices',
      component: () => import('@/pages/InvoiceListPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/invoices/new',
      name: 'invoice-create',
      component: () => import('@/pages/AuthInvoicePage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/invoices/:id',
      name: 'invoice-detail',
      component: () => import('@/pages/InvoiceDetailPage.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/invoices/:id/edit',
      name: 'invoice-edit',
      component: () => import('@/pages/AuthInvoicePage.vue'),
      meta: { requiresAuth: true },
    },
  ],
})

router.beforeEach((to) => {
  if (to.meta.requiresAuth) {
    const token = localStorage.getItem('jwt_token')
    if (!token) {
      return { name: 'login', query: { redirect: to.fullPath } }
    }
  }
  if (to.name === 'login') {
    const token = localStorage.getItem('jwt_token')
    if (token) {
      return { name: 'dashboard' }
    }
  }
})

export default router
