import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/promocodes/list',
    method: 'get',
    params: query
  })
}

export function fetchPage(id) {
  return request({
    url: '/promocodes/'+id,
    method: 'get',
  })
}

export function createPage(data) {
  return request({
    url: '/promocodes/',
    method: 'post',
    data
  })
}

export function updatePage(id, data) {
  return request({
    url: '/promocodes/update/'+id,
    method: 'put',
    data
  })
}
