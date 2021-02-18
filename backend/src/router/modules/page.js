/** When your routing table is too long, you can split it into small modules**/

import Layout from '@/views/layout/Layout'

const pageRouter = {
  path: '/page',
  component: Layout,
  redirect: '/page/menu1/menu1-1',
  name: 'Page',
  meta: {
    title: 'page',
    icon: 'list'
  },
  children: [
    {
      path: 'create',
      component: () => import('@/views/page/create'),
      name: 'CreatePage',
      meta: { title: 'createPage', icon: 'edit' }
    },
    {
      path: 'edit/:id(\\S+)',
      component: () => import('@/views/page/edit'),
      name: 'EditPage',
      meta: { title: 'editPage', noCache: true },
      hidden: true
    },
    {
      path: 'list',
      component: () => import('@/views/page/list'),
      name: 'PageList',
      meta: { title: 'pageList', icon: 'list' }
    }
  ]
}

export default pageRouter
