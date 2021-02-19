/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/views/layout/Layout'

const promoCodeRouter = {
  path: '/promoCode',
  component: Layout,
  redirect: '/promoCode/menu1/menu1-1',
  name: 'PromoCode',
  meta: {
    title: 'promoCode',
    icon: 'list'
  },
  children: [
    {
      path: 'create',
      component: () => import('@/views/promoCode/create'),
      name: 'CreatePromoCode',
      meta: { title: 'createPromoCode', icon: 'edit' }
    },
    {
      path: 'edit/:id(\\S+)',
      component: () => import('@/views/promoCode/edit'),
      name: 'EditPromoCode',
      meta: { title: 'editPromoCode', noCache: true },
      hidden: true
    },
    {
      path: 'list',
      component: () => import('@/views/promoCode/list'),
      name: 'PromoCodeList',
      meta: { title: 'promoCodeList', icon: 'list' }
    }
  ]
}

export default promoCodeRouter
