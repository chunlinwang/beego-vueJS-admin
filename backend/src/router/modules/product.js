/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/views/layout/Layout'

const productRouter = {
  path: '/product',
  component: Layout,
  redirect: '/product/menu1/menu1-1',
  name: 'Product',
  meta: {
    title: 'product',
    icon: 'list'
  },
  children: [
    {
      path: 'create',
      component: () => import('@/views/product/create'),
      name: 'CreateProduct',
      meta: { title: 'createProduct', icon: 'edit' }
    },
    {
      path: 'edit/:id(\\S+)',
      component: () => import('@/views/product/edit'),
      name: 'EditProduct',
      meta: { title: 'editProduct', noCache: true },
      hidden: true
    },
    {
      path: 'list',
      component: () => import('@/views/product/list'),
      name: 'PageList',
      meta: { title: 'productList', icon: 'list' }
    }
  ]
}

export default productRouter
