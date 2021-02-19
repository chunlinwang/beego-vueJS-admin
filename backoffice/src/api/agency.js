import request from '@/utils/request'
import qs from 'qs'

export function save(data) {
  return request({
    url: '/agency/update',
    method: 'put',
    data: qs.stringify(data)
  })
}

export function getAgency() {
  return request({
    url: '/agency/info',
    method: 'get'
  })
}
