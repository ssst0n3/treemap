<template>
  <div>
    <js-mind :key="refresh_key" :values="mind" :height="`${height}px`" ref="js_mind"/>
  </div>
</template>

<script>
import lightweightRestful from 'vue-lightweight_restful'
import consts from "@/utils/consts";
import api from "@/utils/api";

export default {
  name: "TreemapItem",
  props: {
    root_node_id: Number,
    height: Number,
  },
  data: function () {
    return {
      add_queue: {},
      refresh_key: 1,
      mind: {
        "meta": {"name": "treemap", "author": "ssst0n3@gmail.com", "version": "0.1"},
        "format": "node_tree",
        "data": {
          "id": 1,
          "topic": "init",
          "children": []
        },
      },
      tree: {}
    }
  },
  async created() {
    // this.root_node_id = this.$route.params.root_node_id
    await this.refresh()
  },
  methods: {
    // 获取选中标签的 ID
    get_selected_node_id() {
      let selected_node = this.$refs.js_mind.jm.get_selected_node();
      if (selected_node) {
        return selected_node.id;
      } else {
        return null;
      }
    },
    async refresh() {
      this.tree = await lightweightRestful.api.get(consts.api.v1.node.tree(this.root_node_id))
      this.mind.data = this.load(this.tree)
      this.refresh_key++
      this.$nextTick(() => {
        this.$refs.js_mind.jm.add_event_listener(this.jm_listener);
        this.$refs.js_mind.jm.view.add_event(this, 'click', this.click_handle);
      })
    },
    click_handle(e) {
      let element = e.target
      let node_id = this.$refs.js_mind.jm.view.get_binded_nodeid(element);
      // if (node_id) {
      console.log(node_id)
      this.$emit('js_mind_click_node', node_id)
      // }
    },
    async jm_listener(type, data) {
      if (type === consts.jm.type.edit) {
        switch (data.evt) {
          case "insert_node_after": {
            await this.insert_node(data)
            break
          }
          case "insert_node_before": {
            await this.insert_node(data)
            break
          }
          case "add_node": {
            if (data.data[2] === 'New Node') {
              this.add_queue[data.data[1]] = data.data
            } else {
              let parent = data.data[0]
              let name = data.data[2]
              await api.add_node(parent, name)
            }
            break
          }
          case "update_node": {
            let id = data.data[0]
            let name = data.data[1]
            if (id in this.add_queue) {
              let parent = this.add_queue[id][0]
              await api.add_node(parent, name)
              delete this.add_queue[id]
            } else {
              let new_node = {
                name: name
              }
              await api.update_node(id, new_node)
            }
            break
          }
          case "remove_node": {
            let id = data.data[0]
            await api.remove_node(id)
            break
          }
        }
        console.log(data)
      }
    },
    async insert_node(data) {
      let neighbour = data.data[0]
      let id = data.data[1]
      console.log(`insert: neighbour=${neighbour}`)
      let parent = this.$refs.js_mind.jm.mind.nodes[neighbour].parent.id
      data.data[0] = parent
      if (data.data[2] === 'New Node') {
        this.add_queue[id] = data.data
      } else {
        let name = data.data[2]
        await api.add_node(parent, name)
      }
    },
    load(node) {
      let format = {
        'id': node.id,
        'topic': node.name,
        'node_type': node.node_type,
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