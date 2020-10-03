import Vue from 'vue'
import TreemapItem from "@/components/TreemapItem";
import CreateRootNode from "@/components/root_node/CreateRootNode";
import SelectRootNode from "@/components/root_node/SelectRootNode";
import jm from 'vue-jsmind'

Vue.use(jm)

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
}

export default b