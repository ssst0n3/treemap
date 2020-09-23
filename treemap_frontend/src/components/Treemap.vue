<template>
  <div>
    <!--    <div id="jsmind_container" ref="jsmind_container"/>-->
    <js-mind :values="mind" height="1000px"></js-mind>
    <radial-menu
        style="position: fixed; z-index: 3; margin-top: 1%; margin-left:85%; background-color: white"
        :itemSize="60"
        :radius="120"
        :angle-restriction="180"
        :rotate="180">
      <radial-menu-item
          v-for="(item, index) in items"
          :key="index"
          style="background-color: white"
          @click="() => handleClick(item)">
        <span>{{ item }}</span>
      </radial-menu-item>
    </radial-menu>
  </div>
</template>

<script>
import jsMind from 'jsmind'
// import 'jsmind/style/jsmind.css'

import Vue from 'vue'
import jm from 'vue-jsmind'
import {RadialMenu, RadialMenuItem} from 'vue-radial-menu'

Vue.use(jm)

export default {
  name: "Treemap",
  components: {
    RadialMenu,
    RadialMenuItem
  },
  data: function () {
    return {
      tree_example: {
        "topic": "网络安全",
        "children": [
          {"topic": "Web安全", "children": []},
          {"topic": "Linux", "children": []},
          {
            "topic": "CTF", "children": [
              {"topic": "<a href='xxx' style='color: white'> <i class='fas fa-book'> Web </i></a>", "children": []},
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
  created() {
    this.mind.data = this.load(this.tree == null ? this.tree_example : this.tree)
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
    set_container_size() {
      let container = this.$refs.jsmind_container;
      container.style.height = window.innerHeight - 50 + 'px'
      container.style.width = window.innerWidth - 8 + 'px'
    },
    load(node) {
      let format = {
        'id': jsMind.util.uuid.newid(),
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