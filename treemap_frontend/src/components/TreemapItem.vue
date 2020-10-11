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
    decorate: Function,
    until: {
      type: Boolean,
      required: true
    }
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
    await this.wait_until()
    await this.refresh()
  },
  methods: {
    sleep(wait) {
      return new Promise((resolve) => setTimeout(resolve, +wait || 0))
    },
    async wait_until(wait = 0) {
      console.log('waiting:', wait)
      if (this.until || wait === 8) {
        console.log('stop wait')
        return
      }
      await this.sleep(100 * Math.pow(2, wait)).then(
          // await this.wait_until(wait + 1)
      )
      await this.wait_until(wait + 1)
    },
    node_content_is_init_status(node_id) {
      return this.$refs.js_mind.jm.mind.nodes[node_id].data.content_type === consts.model.node.content_type.default
    },
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
      let response = await lightweightRestful.api.get(consts.api.v1.node.tree(this.root_node_id), null, this, 'please login first.')
      if (response.success !== false) {
        this.tree = response
      } else {
        return
      }
      this.mind.data = this.load(this.tree)
      this.refresh_key++
      this.$nextTick(() => {
        this.$refs.js_mind.jm.add_event_listener(this.jm_listener);
        this.$refs.js_mind.jm.view.add_event(this, 'click', this.click_handle);
      })
    },
    click_handle(e) {
      let element = e.target
      let is_expander = this.$refs.js_mind.jm.view.is_expander(element)
      if (!is_expander) {
        let node_id = this.$refs.js_mind.jm.view.get_binded_nodeid(element);
        if (node_id) {
          node_id = this.$refs.js_mind.jm.mind.nodes[node_id].id
        }
        // console.log(node_id)
        this.$emit('js_mind_click_node', node_id)
      }
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
              let new_id = await api.add_node(parent, name)
              delete this.add_queue[id]
              let node = this.$refs.js_mind.jm.mind.nodes[id]
              node.id = new_id
              this.$refs.js_mind.jm.mind.nodes[new_id] = node
            } else {
              await api.update_name(id, name)
            }
            break
          }
          case "remove_node": {
            let id = data.data[0]
            await api.remove_node(id)
            break
          }
          case "move_node": {
            let node_id = data.data[0]
            let before_id = data.data[1]
            let parent_id = data.data[2]
            await api.move_node(node_id, before_id, parent_id)
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
      let topic = this.decorate(node.name, node.content_type, node.content_id)
      let format = {
        'id': node.id,
        'topic': topic,
        'node_type': node.node_type,
        'children': [],
        'index': node.index,
        'content_type': node.content_type,
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