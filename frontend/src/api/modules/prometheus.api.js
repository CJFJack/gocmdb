export default ({service, request, serviceForMock, requestForMock, mock, faker, tools}) => ({
    LIST_PROMETHEUS_NODES(data = {}) {
        // 接口请求
        return request({
            url: '/node/query',
            method: 'post',
            data
        })
    },
    DEL_PROMETHEUS_NODE(data = {}) {
        // 接口请求
        return request({
            url: '/node/delete',
            method: 'post',
            data
        })
    },
    ADD_PROMETHEUS_NODE(data = {}) {
        // 接口请求
        return request({
            url: '/node/add',
            method: 'post',
            data
        })
    },
    MODIFY_PROMETHEUS_NODE(data = {}) {
        // 接口请求
        return request({
            url: '/node/modify',
            method: 'post',
            data
        })
    },
    LIST_PROMETHEUS_JOBS(data = {}) {
        // 接口请求
        return request({
            url: '/job/query',
            method: 'post',
            data
        })
    },
    DEL_PROMETHEUS_JOB(data = {}) {
        // 接口请求
        return request({
            url: '/job/delete',
            method: 'post',
            data
        })
    },
    ADD_PROMETHEUS_JOB(data = {}) {
        // 接口请求
        return request({
            url: '/job/add',
            method: 'post',
            data
        })
    },
    MODIFY_PROMETHEUS_JOB(data = {}) {
        // 接口请求
        return request({
            url: '/job/modify',
            method: 'post',
            data
        })
    },
    LIST_PROMETHEUS_TARGET(data = {}) {
        // 接口请求
        return request({
            url: '/target/query',
            method: 'post',
            data
        })
    },
    DEL_PROMETHEUS_TARGET(data = {}) {
        // 接口请求
        return request({
            url: '/target/delete',
            method: 'post',
            data
        })
    },
    ADD_PROMETHEUS_TARGET(data = {}) {
        // 接口请求
        return request({
            url: '/target/add',
            method: 'post',
            data
        })
    },
    MODIFY_PROMETHEUS_TARGET(data = {}) {
        // 接口请求
        return request({
            url: '/target/modify',
            method: 'post',
            data
        })
    },
})

