const api_ = '/api'
const v1 = `${api_}/v1`
const api = {
    v1: {
        node: `${v1}/node`
    }
}
export default {
    BaseUrl: process.env.NODE_ENV === 'development' ? 'http://127.0.0.1:12000' : '/',
    jm: {
       type: {
           edit: 3
       }
    },
    api: {
        v1: {
            node: {
                root: `${api.v1.node}/root`,
                tree: (id) => `${api.v1.node}/root/${id}`,
                child: `${api.v1.node}/child`
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