<template>
  <div>
    <js-mind :key="refresh" :values="mind" :height="height" ref="js_mind"></js-mind>
  </div>
</template>

<script>
// import lightweightRestful from 'vue-lightweight_restful'
// import consts from "@/utils/consts";


export default {
  name: "Treemap",
  components: {},
  data: function () {
    return {
      is_mounted: false,
      refresh: 1,
      selected: 0,
      options: [
        {value: null, text: 'Please select an option'},
      ],
      tree_example: {
        "id": 1,
        "topic": "网络安全",
        "children": [
          {"id": 2, "topic": "Web安全", "children": []},
          {"id": 3, "topic": "Linux", "children": []},
          {
            "id": 4, "topic": "CTF", "children": [
              {
                "id": 5,
                "topic": "<a href='xxx' style='color: white'> <i class='fas fa-book'> Web </i></a>",
                "children": []
              },
            ]
          },
        ]
      },
      jm: {},
      mind: {
        "meta": {"name": "treemap", "author": "ssst0n3@gmail.com", "version": "0.1"},
        "format": "node_tree",
        "data": {},
      },
      old: {},
    }
  },
  props: {
    tree: Object,
    editable: Boolean,
    height: String
  },
  computed: {
    jm_value() {
      if (!this.is_mounted)
        return;
      return this.get_jm_data()
    }
  },
  async created() {
    this.$set(this.mind, 'data', this.load(this.tree))
    // this.mind.data = this.load(this.tree == null ? this.tree_example : this.tree)
  },
  watch: {
    tree: function () {
      // this.mind.data = this.load(this.tree == null ? this.tree_example : this.tree)
      this.$set(this.mind, 'data', this.load(this.tree))
      this.refresh += 1
      console.log('jm_value tree:', this.jm_value)
      console.log(this.get_jm_data())
    },
    jm_value: {
      handler() {
        console.log('jm_value watch:', this.jm_value)
      }, deep: true, immediate: true
    }
  },
  mounted() {
    this.is_mounted = true
    // this.$watch('$refs.js_mind.jm.mind.nodes', (newValue) => {
    //   console.log("jm.get_data", this.$refs.js_mind.jm.get_data())
    //   console.log('new:', this.record(newValue))
    //   console.log('old:', this.old)
    //   if (Object.keys(this.old).length > 0) {
    //     this.diff(this.record(newValue), this.old)
    //   }
    //   this.old = this.record(newValue)
    // }, {deep: true, immediate: true})
  },
  methods: {
    get_jm_data() {
      return this.$refs.js_mind.jm.get_data()
    },
    diff(n, o) {
      Object.keys(n).forEach(key => {
        if (o.key !== undefined) {
          console.log(key, 'exist', o[key])
        } else {
          let newNode = {
            'name': n[key].topic,
            'parent': n[key].id
          }
          console.log("add", newNode)
          // lightweightRestful.api.post(consts.api.v1.node, null, newNode)
        }
      })
    },
    record(nodes) {
      let r = {}
      Object.keys(nodes).forEach(key => {
            if (nodes[key].topic !== undefined) {
              r[key] = nodes[key].topic
            }
          }
      )
      return r
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
    }
  }
}
</script>

<style scoped>

</style>