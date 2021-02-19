import { getAgency } from '@/api/agency'
import _ from 'lodash'

const agency = {
  state: {
    info: {
      address: '',
      city: '',
      zip_code: '',
      phone: '',
      email: '',
      schedules: '',
      is_open: ''
    }
  },

  mutations: {
    SET_INFO: (state, info) => {
      state.info = info
    }
  },

  actions: {
    // 用户名登录
    getAgency({ commit }) {
      return new Promise((resolve, reject) => {
        getAgency().then(response => {
          _.map(response.data.schedules, function(s) {
            if (s.begin_hour === '0001-01-01T00:00:00Z') {
              s.begin_hour = null
            }
            if (s.end_hour === '0001-01-01T00:00:00Z') {
              s.end_hour = null
            }
          })

          const data = response.data

          commit('SET_INFO', data)

          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },
  }

}

export default agency
