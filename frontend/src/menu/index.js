import { uniqueId } from 'lodash'

/**
 * @description 给菜单数据补充上 path 字段
 * @description https://github.com/d2-projects/d2-admin/issues/209
 * @param {Array} menu 原始的菜单数据
 */
function supplementPath (menu) {
  return menu.map(e => ({
    ...e,
    path: e.path || uniqueId('d2-menu-empty-'),
    ...e.children ? {
      children: supplementPath(e.children)
    } : {}
  }))
}

export const menuHeader = supplementPath([
  { path: '/index', title: '首页', icon: 'home' },
  // {
  //   title: '页面',
  //   icon: 'folder-o',
  //   children: [
  //     { path: '/page1', title: '页面 api.conf' },
  //     { path: '/page2', title: '页面 2' },
  //     { path: '/page3', title: '页面 3' }
  //   ]
  // }
])

export const menuAside = supplementPath([
  { path: '/index', title: '首页', icon: 'home' },
  { path: '/users', title: '用户管理', icon: 'user' },
  {
    title: '云管理',
    icon: 'cloud',
    children: [
      { path: '/cloud/platform_management', title: '云平台管理', icon: 'cube' },
      { path: '/cloud/virtual_machine', title: '云主机管理', icon: 'server' },
    ]
  },
  {
    title: 'Prometheus',
    icon: 'eercast',
    children: [
      { path: '/prometheus/node', title: 'Node', icon: 'linode' },
      { path: '/prometheus/job', title: 'Job', icon: 'linode' },
      { path: '/prometheus/target', title: 'Target', icon: 'linode' },
    ]
  }
])
