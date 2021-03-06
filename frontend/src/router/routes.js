import layoutHeaderAside from '@/layout/header-aside'

// 由于懒加载页面太多的话会造成webpack热更新太慢，所以开发环境不使用懒加载，只有生产环境使用懒加载
const _import = require('@/libs/util.import.' + process.env.NODE_ENV)

/**
 * 在主框架内显示
 */
const frameIn = [
    {
        path: '/',
        redirect: {name: 'index'},
        component: layoutHeaderAside,
        children: [
            // 首页
            {
                path: 'index',
                name: 'index',
                meta: {
                    auth: true
                },
                component: _import('system/index')
            },
            // 用户管理
            {
                path: 'users',
                name: 'users',
                meta: {
                    title: '用户管理',
                    auth: true
                },
                component: _import('users')
            },
            // 云平台管理
            {
                path: 'cloud/platform_management',
                name: 'cloud_platform_management',
                meta: {
                    title: '云平台管理',
                    auth: true
                },
                component: _import('cloud/platform_management')
            },
            // 云主机管理
            {
                path: 'cloud/virtual_machine',
                name: 'cloud_virtual_machine',
                meta: {
                    title: '云主机管理',
                    auth: true
                },
                component: _import('cloud/virtual_machine')
            },
            // Prometheus - Node节点
            {
                path: 'prometheus/node',
                name: 'prometheus_node',
                meta: {
                    title: 'Prometheus - Node节点',
                    auth: true
                },
                component: _import('prometheus/node')
            },
            // Prometheus - Job
            {
                path: 'prometheus/job',
                name: 'prometheus_job',
                meta: {
                    title: 'Prometheus - Job',
                    auth: true
                },
                component: _import('prometheus/job')
            },
            // Prometheus - Target
            {
                path: 'prometheus/target',
                name: 'prometheus_target',
                meta: {
                    title: 'Prometheus - Target',
                    auth: true
                },
                component: _import('prometheus/target')
            },
            // 系统 前端日志
            {
                path: 'log',
                name: 'log',
                meta: {
                    title: '前端日志',
                    auth: true
                },
                component: _import('system/log')
            },
            // 刷新页面 必须保留
            {
                path: 'refresh',
                name: 'refresh',
                hidden: true,
                component: _import('system/function/refresh')
            },
            // 页面重定向 必须保留
            {
                path: 'redirect/:route*',
                name: 'redirect',
                hidden: true,
                component: _import('system/function/redirect')
            }
        ]
    }
]

/**
 * 在主框架之外显示
 */
const frameOut = [
    // 登录
    {
        path: '/login',
        name: 'login',
        component: _import('system/login')
    }
]

/**
 * 错误页面
 */
const errorPage = [
    {
        path: '*',
        name: '404',
        component: _import('system/error/404')
    }
]

// 导出需要显示菜单的
export const frameInRoutes = frameIn

// 重新组织后导出
export default [
    ...frameIn,
    ...frameOut,
    ...errorPage
]
