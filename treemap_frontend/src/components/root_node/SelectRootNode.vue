<template>
  <div>
    <div class="col">
      <b-form-select v-model="selected" :options="options" size="sm" class="mt-3"/>
      <div class="mt-3">Select root node <strong>{{ selected }}</strong></div>
    </div>
  </div>
</template>

<script>
import lightweightRestful from "vue-lightweight_restful";
import consts from "@/utils/consts";

export default {
  name: "SelectRootNode",
  props: {
    default: Number
  },
  data: function () {
    return {
      selected: this.default === undefined ? 1 : this.default,
      options: [
        {value: null, text: 'Please select an option'},
      ],
    }
  },
  async created() {
    await this.generate_options()
  },
  watch: {
    selected: function (n, o) {
      this.$emit('root_node_changed', n, o)
    }
  },
  methods: {
    async generate_options() {
      let options = []
      let nodes = await lightweightRestful.api.get(consts.api.v1.node.root)
      nodes.forEach(node => {
        options.push({'value': node.id, 'text': node.name})
      })
      this.options = options
    }
  }
}
</script>

<style scoped>

</style>