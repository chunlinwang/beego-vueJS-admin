import request from '@/utils/request'

export function fetchList(query) {
  return request({
    url: '/menus/list',
    method: 'get',
    params: query
  })
}

export function fetchMenu(id) {
  return request({
    url: '/menus/' + id,
    method: 'get'
  })
}

export function createMenu(data) {
  return request({
    url: '/menus/',
    method: 'post',
    data
  })
}

export function updateMenu(id, data) {
  return request({
    url: '/menus/update/' + id,
    method: 'put',
    data: data
  })
}

export function getMenus(query) {
  return request({
    url: '/menus',
    method: 'get',
    params: query
  })
}
