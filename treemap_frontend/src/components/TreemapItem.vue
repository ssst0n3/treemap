<template>
    <div>
        <js-mind :key="refresh_key" :values="mind" :height="height" ref="js_mind"></js-mind>
    </div>
</template>

<script>
    import lightweightRestful from 'vue-lightweight_restful'
    import consts from "@/utils/consts";

    export default {
        name: "TreemapItem",
        data: function () {
            return {
                add_queue: {},
                refresh_key: 1,
                root_node_id: 0,
                height: '500px',
                mind: {
                    "meta": {"name": "treemap", "author": "ssst0n3@gmail.com", "version": "0.1"},
                    "format": "node_tree",
                    "data": {
                        "id": 1,
                        "topic": "init",
                        "children": []
                    },
                },
                tree: {},
            }
        },
        async created() {
            this.root_node_id = this.$route.params.root_node_id
            await this.refresh()
        },
        methods: {
            jm_listener(type, data) {
                if (type === consts.jm.type.edit) {
                    switch (data.evt) {
                        case "insert_node_after": {
                            this.insert_node(data)
                            break
                        }
                        case "insert_node_before": {
                            this.insert_node(data)
                            break
                        }
                        case "add_node": {
                            if (data.data[2] === 'New Node') {
                                this.add_queue[data.data[1]] = data.data
                            } else {
                                let parent = data.data[0]
                                let name = data.data[2]
                                this.add_node(parent, name)
                            }
                            break
                        }
                        case "update_node": {
                            let id = data.data[0]
                            if (id in this.add_queue) {
                                let parent = this.add_queue[id][0]
                                let name = data.data[1]
                                this.add_node(parent, name)
                                delete this.add_queue[id]
                            } else {
                                this.update_node()
                            }
                            break
                        }
                        case "remove_node": {
                            this.remove_node()
                            break
                        }
                    }
                    console.log(data)
                }
            },
            insert_node(data) {
                let neighbour = data.data[0]
                let id = data.data[1]
                console.log(`insert: neighbour=${neighbour}`)
                let parent = this.$refs.js_mind.jm.mind.nodes[neighbour].parent.id
                data.data[0] = parent
                if (data.data[2] === 'New Node') {
                    this.add_queue[id] = data.data
                } else {
                    let name = data.data[2]
                    this.add_node(parent, name)
                }
            },
            add_node(parent, name) {
                console.log(`add_node: parent=${parent}, name=${name}`)
                lightweightRestful.api.post(consts.api.v1.node.child, null, {
                    parent: parent,
                    name: name,
                    node_type: consts.model.node.node_type.node,
                })
            },
            update_node() {
                console.log("update_node")
            },
            remove_node() {
                console.log("remove_node")
            },
            async refresh() {
                this.tree = await lightweightRestful.api.get(consts.api.v1.node.tree(this.root_node_id))
                this.mind.data = this.load(this.tree)
                this.refresh_key++
                this.$nextTick(() => {
                    this.$refs.js_mind.jm.add_event_listener(this.jm_listener);
                })
            },
            get_jm_data() {
                return this.$refs.js_mind.jm.get_data()
            },
            load(node) {
                let format = {
                    'id': node.id,
                    'topic': node.name,
                    'children': []
                }
                if (node.Sub) {
                    node.Sub.forEach(
                        child => {
                            format.children.push(this.load(child))
                        }
                    )
                }
                return format
            }
        }
    }
</script>

<style scoped>

</style>