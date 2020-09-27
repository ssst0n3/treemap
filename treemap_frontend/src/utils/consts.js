const api = {
    v1: '/api/v1',
}
export default {
    BaseUrl: process.env.NODE_ENV === 'development' ? 'http://127.0.0.1:12000' : '/',
    api: {
        v1: {
            node: api.v1 + '/node/'
        }
    }
}