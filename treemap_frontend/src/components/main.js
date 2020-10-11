import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import TreemapItem from "@/components/TreemapItem";
import CreateRootNode from "@/components/root_node/CreateRootNode";
import SelectRootNode from "@/components/root_node/SelectRootNode";
import jm from 'vue-jsmind'
import consts from "@/utils/consts";
import util from "@/utils/util";
import api from "@/utils/api";
import lightweightRestful from 'vue-lightweight_restful'

util.InitRestfulClient(consts.BaseUrl)
Vue.use(jm)
Vue.use(BootstrapVue)
Vue.use(lightweightRestful)

const components = [
    TreemapItem,
    CreateRootNode,
    SelectRootNode,
]

const install = function (Vue) {
    components.forEach(c => {
        Vue.component(c.name, c);
    })
    // Vue.component(TreemapItem.name, TreemapItem)
    // Vue.component(CreateRootNode.name, CreateRootNode)
    // Vue.component(SelectRootNode.name, SelectRootNode)
}

if (typeof window !== 'undefined' && window.Vue) {
    install(window.Vue);
}

const b = {
    install,
    ...components,
    consts,
    util,
    api,
}

export default b