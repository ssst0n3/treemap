import lightweightRestful from "vue-lightweight_restful";
import consts from "@/utils/consts";

export default {
    async add_node(parent, name) {
        console.log(`add_node: parent=${parent}, name=${name}`)
        await lightweightRestful.api.post(consts.api.v1.node.child, null, {
            parent: parent,
            name: name,
            node_type: consts.model.node.node_type.node,
        })
    },
    async update_name(id, data) {
        console.log("update_node")
        let params = {
            action: consts.model.node.action.action_update_node_name
        }
        await lightweightRestful.api.put(consts.api.v1.node.node_item(id), params, data)
    },
    async update_node_content(id, content_id, content_type) {
        console.log("update_node_content")
        let params = {
            action: consts.model.node.action.action_update_node_content
        }
        let data = {
            content_id: content_id,
            content_type: content_type
        }
        await lightweightRestful.api.put(consts.api.v1.node.node_item(id), params, data)
    },
    async remove_node(id) {
        console.log("remove_node")
        await lightweightRestful.api.delete(consts.api.v1.node.child_item(id), null)
    },
    async move_node(id, before_id, parent_id) {
        console.log("move_node")

        let action = consts.model.node.action.action_move_node
        switch (before_id) {
            case "_first_": {
                action = consts.model.node.action.action_move_node_first
                before_id = 0
                break
            }
            case "_last_": {
                action = consts.model.node.action.action_move_node_last
                before_id = 0
                break
            }
        }

        let params = {
            action: action
        }
        let data = {
            before_id: before_id,
            parent: parent_id,
        }
        await lightweightRestful.api.put(consts.api.v1.node.node_item(id), params, data)
    }
}