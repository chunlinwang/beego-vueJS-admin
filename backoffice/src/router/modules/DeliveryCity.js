/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/views/layout/Layout'

const deliveryCityRouter = {
  path: '/deliveryCity',
  component: Layout,
  redirect: '/deliveryCity/menu1/menu1-1',
  name: 'DeliveryCity',
  meta: {
    title: 'deliveryCity',
    icon: 'list'
  },
  children: [
    {
      path: 'create',
      component: () => import('@/views/deliveryCity/create'),
      name: 'CreateDeliveryCity',
      meta: { title: 'createDeliveryCity', icon: 'edit' }
    },
    {
      path: 'edit/:id(\\S+)',
      component: () => import('@/views/deliveryCity/edit'),
      name: 'EditDeliveryCity',
      meta: { title: 'editDeliveryCity', noCache: true },
      hidden: true
    },
    {
      path: 'list',
      component: () => import('@/views/deliveryCity/list'),
      name: 'DeliveryCityList',
      meta: { title: 'deliveryCityList', icon: 'list' }
    }
  ]
}

export default deliveryCityRouter
