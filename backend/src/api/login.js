import request from '@/utils/request'
import qs from 'qs'

export function loginByUsername(username, password) {
  const data = {
    'username': username,
    'password': password
  }
  return request({
    url: '/users/login',
    method: 'post',
    data: qs.stringify(data)
  })
}

export function logout() {
  return request({
    url: '/users/logout',
    method: 'post'
  })
}

export function getUserInfo(token) {
  return request({
    url: '/users/info',
    method: 'get',
    params: { token }
  })
}

