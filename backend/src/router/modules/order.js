/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/views/layout/Layout'

const orderRouter = {
  path: '/order',
  component: Layout,
  // redirect: '/order/menu1/menu1-1',
  alwaysShow: false,
  name: 'Order',
  meta: {
    title: 'order',
    icon: 'shopping'
  },
  children: [
    {
      path: 'list',
      component: () => import('@/views/order/list'),
      name: 'Order',
      meta: { title: 'order', icon: 'shopping' }
    }
  ]
}

export default orderRouter
