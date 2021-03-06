import lightweightRestful from "vue-lightweight_restful";
import consts from "@/utils/consts";

export default {
    async add_node(parent, name, toast_options) {
        console.log(`add_node: parent=${parent}, name=${name}`)
        let response = await lightweightRestful.api.post(consts.api.v1.node.child, null, {
            parent: parent,
            name: name,
            node_type: consts.model.node.node_type.node,
        }, toast_options)
        return response.id
    },
    async update_node(id, action, data, toast_options) {
        console.log("update_node:", action)
        let params = {
            action: action
        }
        await lightweightRestful.api.put(consts.api.v1.node.node_item(id), params, data, toast_options)
    },
    async update_name(id, name, toast_options) {
        let data = {
            name: name
        }
        await this.update_node(id, consts.model.node.action.action_update_node_name, data, toast_options)
    },
    async update_node_content(id, content_id, content_type, toast_options) {
        let data = {
            content_id: content_id,
            content_type: content_type
        }
        await this.update_node(id, consts.model.node.action.action_update_node_content, data, toast_options)
    },
    async remove_node(id, toast_options) {
        console.log("remove_node")
        await lightweightRestful.api.delete(consts.api.v1.node.child_item(id), null, toast_options)
    },
    async move_node(id, before_id, parent_id, toast_options) {
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
        let data = {
            before_id: before_id,
            parent: parent_id,
        }
        await this.update_node(id, action, data, toast_options)
    },
}