import consts from "@/utils/consts";
import lightweightRestful from "vue-lightweight_restful";

export default {
    async add_node(parent, name) {
        console.log(`add_node: parent=${parent}, name=${name}`)
        await lightweightRestful.api.post(consts.api.v1.node.child, null, {
            parent: parent,
            name: name,
            node_type: consts.model.node.node_type.node,
        })
    },
    async update_node(id, data) {
        console.log("update_node")
        await lightweightRestful.api.put(consts.api.v1.node.node_item(id), null, data)
    },
    async remove_node(id) {
        console.log("remove_node")
        await lightweightRestful.api.delete(consts.api.v1.node.child_item(id), null)
    },
    async move_node(id, before_id, parent_id) {
        console.log("move_node")
        let data = {
            before_id: before_id,
            parent_id: parent_id,
        }
        await lightweightRestful.api.post(consts.api.v1.node.child_item(id), null, data)
    }
}