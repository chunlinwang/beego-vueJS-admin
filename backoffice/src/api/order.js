import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/orders/list',
    method: 'get',
    params: query
  })
}

export function updateOrder(id, data) {
  return request({
    url: '/orders/update/'+id,
    method: 'put',
    data
  })
}
