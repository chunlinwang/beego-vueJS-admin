import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/deliverycity/list',
    method: 'get',
    params: query
  })
}

export function fetchDeliveryCity(id) {
  return request({
    url: '/deliverycity/'+id,
    method: 'get',
  })
}

export function createDeliveryCity(data) {
  return request({
    url: '/deliverycity/',
    method: 'post',
    data
  })
}

export function updateDeliveryCity(id, data) {
  return request({
    url: '/deliverycity/update/'+id,
    method: 'put',
    data
  })
}
