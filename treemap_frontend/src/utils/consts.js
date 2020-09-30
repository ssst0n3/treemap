const api_ = '/api'
const v1 = `${api_}/v1`
const api = {
    v1: {
        node: `${v1}/node`
    }
}
export default {
    BaseUrl: process.env.NODE_ENV === 'development' ? 'http://127.0.0.1:12000' : '/',
    api: {
        v1: {
            node: {
                root: `${api.v1.node}/root`,
                tree: (id) => `${api.v1.node}/root/${id}`
            }
        }
    },
    model: {
        node: {
            node_type: {
                root: 0,
                node: 1,
                leaf: 2
            }
        }
    }
}