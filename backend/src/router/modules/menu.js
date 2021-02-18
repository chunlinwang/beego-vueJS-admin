/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/views/layout/Layout'

const menuRouter = {
  path: '/menu',
  component: Layout,
  redirect: '/menu/menu1/menu1-1',
  name: 'Menu',
  meta: {
    title: 'menu',
    icon: 'list'
  },
  children: [
    {
      path: 'create',
      component: () => import('@/views/menu/create'),
      name: 'CreateMenu',
      meta: { title: 'createMenu', icon: 'edit' }
    },
    {
      path: 'edit/:id(\\S+)',
      component: () => import('@/views/menu/edit'),
      name: 'EditArticle',
      meta: { title: 'editMenu', noCache: true },
      hidden: true
    },
    {
      path: 'list',
      component: () => import('@/views/menu/list'),
      name: 'ArticleList',
      meta: { title: 'menuList', icon: 'list' }
    }
  ]
}

export default menuRouter
