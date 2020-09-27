<template>
  <div>
    <!--    <div id="jsmind_container" ref="jsmind_container"/>-->
    <div class="container">
      <div class="row">
        <div class="col">
          <b-form-select v-model="selected" :options="options" size="sm" class="mt-3"/>
          <div class="mt-3">Select root node <strong>{{ selected }}</strong></div>
        </div>
      </div>
    </div>
    <js-mind :values="mind" height="1000px"></js-mind>

  </div>
</template>

<script>
import Vue from 'vue'
import jm from 'vue-jsmind'
import lightweightRestful from 'vue-lightweight_restful'
import consts from "@/utils/consts";

Vue.use(jm)

export default {
  name: "Treemap",
  components: {},
  data: function () {
    return {
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
    }
  },
  props: {
    tree: Object,
    editable: Boolean,
  },
  computed: {},
  async created() {
    this.mind.data = this.load(this.tree == null ? this.tree_example : this.tree)
    await this.generate_options()
  },
  mounted() {
    // this.set_container_size()
    // let options = {
    //   container: 'jsmind_container',
    //   editable: this.editable,
    //   theme: 'asphalt'
    // }
    // this.jm = jsMind.show(options, this.mind)
  },
  methods: {
    async generate_options() {
      let options = []
      let nodes = await lightweightRestful.api.get(consts.api.v1.node)
      nodes.forEach(node => {
        console.log(node)
        options.push({'value': node.id, 'text': node.name})
      })
      this.options = options
    },
    set_container_size() {
      let container = this.$refs.jsmind_container;
      container.style.height = window.innerHeight - 50 + 'px'
      container.style.width = window.innerWidth - 8 + 'px'
    },
    load(node) {
      let format = {
        'id': node.id,
        'topic': node.topic,
        'children': []
      }
      node.children.forEach(
          child => {
            format.children.push(this.load(child))
          }
      )
      return format
    },
    async handleClick(item) {
      console.log(item)
      if (item === 'expand_all') {
        this.jm.expand_all()
      } else if (item === 'collapse_all') {
        this.jm.collapse_all()
      }
    }
  }
}
</script>

<style scoped>

</style>