<template>
  <div id="app">
    <!--    <Example/>-->
    <TreemapItem :height="500" :root_node_id="1" v-on:js_mind_click_node="js_mind_click_node"/>
  </div>
</template>

<script>

import TreemapItem from "@/components/TreemapItem";

export default {
  name: 'App',
  components: {TreemapItem},
  methods: {
    js_mind_click_node(node_id) {
      console.log("clicked:", node_id)
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
