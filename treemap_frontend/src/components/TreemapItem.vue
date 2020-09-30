<template>
  <div>
    <js-mind :key="refresh_key" :values="mind" :height="height" ref="js_mind" :v-if="renderComponent"></js-mind>
  </div>
</template>

<script>
import lightweightRestful from 'vue-lightweight_restful'
import consts from "@/utils/consts";

export default {
  name: "TreemapItem",
  data: function () {
    return {
      renderComponent: true,
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
  mounted() {
    this.$refs.js_mind.jm.add_event_listener(function (type, data) {
      console.log("listener:", type, data)
    });
  },
  methods: {
    async refresh() {
      this.tree = await lightweightRestful.api.get(consts.api.v1.node.tree(this.root_node_id))
      // this.$set(this.mind, 'data', this.load(this.tree))
      // Object.assign(this.mind.data, this.load(this.tree))
      this.mind.data = this.load(this.tree)
      this.refresh_key++
      this.$nextTick(
          ()=>{
            this.$refs.js_mind.jm.add_event_listener(function (type, data) {
              console.log("listener:", type, data)
            });
          }
      )
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
      if (node.children !== undefined) {
        node.children.forEach(
            child => {
              format.children.push(this.load(child))
            }
        )
      }
      return format
    },
    forceRerender() {
      // Remove my-component from the DOM
      this.renderComponent = false
      this.$nextTick(() => {
        // Add the component back in
        this.renderComponent = true
      })
    }
  }
}
</script>

<style scoped>

</style>