<template>
  <div id="app">
    <Treemap :editable="true"></Treemap>
  </div>
</template>

<script>
import Treemap from "@/components/Treemap";

export default {
  name: 'App',
  components: {
    Treemap,
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
