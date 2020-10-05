import restful from "vue-lightweight_restful";


export default {
    InitRestfulClient(baseUrl) {
        restful.api.initClient(baseUrl)
    }
}