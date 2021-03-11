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
})
