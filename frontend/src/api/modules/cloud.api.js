export default ({service, request, serviceForMock, requestForMock, mock, faker, tools}) => ({
    LIST_CLOUD_PLATFORM (data = {}) {
        // 接口请求
        return request({
            url: '/cloudplatform/query',
            method: 'post',
            data
        })
    },
    ADD_CLOUD_PLATFORM (data = {}) {
        // 接口请求
        return request({
            url: '/cloudplatform/add',
            method: 'post',
            data
        })
    },
    MODIFY_CLOUD_PLATFORM (data = {}) {
        // 接口请求
        return request({
            url: '/cloudplatform/modify',
            method: 'post',
            data
        })
    },
    DEL_CLOUD_PLATFORM (data = {}) {
        // 接口请求
        return request({
            url: '/cloudplatform/delete',
            method: 'post',
            data
        })
    },
    LIST_VIRTUAL_MACHINE (data = {}) {
        // 接口请求
        return request({
            url: '/virtualmachine/query',
            method: 'post',
            data
        })
    },
    START_VIRTUAL_MACHINE (data = {}) {
        // 接口请求
        return request({
            url: '/virtualmachine/start',
            method: 'post',
            data
        })
    },
    STOP_VIRTUAL_MACHINE (data = {}) {
        // 接口请求
        return request({
            url: '/virtualmachine/stop',
            method: 'post',
            data
        })
    },
    REBOOT_VIRTUAL_MACHINE (data = {}) {
        // 接口请求
        return request({
            url: '/virtualmachine/reboot',
            method: 'post',
            data
        })
    },
})
