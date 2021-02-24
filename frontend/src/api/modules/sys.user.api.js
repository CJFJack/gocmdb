import { find, assign } from 'lodash'

const users = [
  { username: 'admin', password: 'admin', uuid: 'admin-uuid', name: 'Admin' },
  { username: 'editor', password: 'editor', uuid: 'editor-uuid', name: 'Editor' },
  { username: 'user1', password: 'user1', uuid: 'user1-uuid', name: 'User1' }
]

export default ({ service, request, serviceForMock, requestForMock, mock, faker, tools }) => ({
  /**
   * @description 登录
   * @param {Object} data 登录携带的信息
   */
  SYS_USER_LOGIN (data = {}) {
    // 接口请求
    return request({
      url: '/auth/login',
      method: 'post',
      data
    })
  },
  SYS_USER_LOGOUT (data = {}) {
    // 接口请求
    return request({
      url: '/auth/logout',
      method: 'post',
      data
    })
  }
})
