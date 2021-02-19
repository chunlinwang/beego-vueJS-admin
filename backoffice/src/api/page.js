import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/pages/list',
    method: 'get',
    params: query
  })
}

export function fetchPage(id) {
  return request({
    url: '/pages/'+id,
    method: 'get',
  })
}

export function createPage(data) {
  return request({
    url: '/pages/',
    method: 'post',
    data
  })
}

export function updatePage(id, data) {
  return request({
    url: '/pages/update/'+id,
    method: 'put',
    data
  })
}
