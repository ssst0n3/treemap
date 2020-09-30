import Router from "vue-router";
import TreemapItem from "@/components/TreemapItem";

export default new Router({
    routes: [
        {
            path: '/tree/:root_node_id',
            name: 'TreemapItem',
            component: TreemapItem
        }
    ]
})
