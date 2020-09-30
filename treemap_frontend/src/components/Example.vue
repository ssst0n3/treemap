<template>
  <div>
    <CreateRootNode v-on:create_new_root_node="create_new_root_node"/>
    <SelectRootNode ref="select_root_node" v-on:root_node_changed="root_node_changes"/>
    <Treemap height="500px" :tree="tree"/>
  </div>
</template>

<script>
import SelectRootNode from "@/components/root_node/SelectRootNode";
import CreateRootNode from "@/components/root_node/CreateRootNode";
import Treemap from "@/components/Treemap";
import lightweightRestful from 'vue-lightweight_restful'
import consts from "@/utils/consts";

export default {
  name: "Example",
  components: {Treemap, CreateRootNode, SelectRootNode},
  data() {
    return {
      tree: {
        "id": 1,
        "name": "网络安全",
        "children": []
      }
    }
  },
  async created() {
    await this.refresh_tree(1)
  },
  mounted() {
    // TODO: add to blog
    // this.$watch('$refs.select_root_node.selected', (newValue, oldValue) => {
    //       console.log(newValue, oldValue)
    //     }, {deep: true}
    // )
  },
  methods: {
    async root_node_changes(n, o) {
      console.log(n, o)
      await this.refresh_tree(n)
    },
    create_new_root_node(id) {
      this.$refs.select_root_node.generate_options()
      this.$refs.select_root_node.selected = id
    },
    async refresh_tree(id) {
      let response = await lightweightRestful.api.get(consts.api.v1.node.tree(id))
      // console.log(response)
      // this.tree = response
      // Object.assign(this.tree, response)
      this.tree = response
    },
    set_container_size() {
      let container = this.$refs.jsmind_container;
      container.style.height = window.innerHeight - 50 + 'px'
      container.style.width = window.innerWidth - 8 + 'px'
    }
  }
}
</script>

<style scoped>

</style>