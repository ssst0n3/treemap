import Vue from 'vue'
import App from './App.vue'
import Router from 'vue-router'
import router from "@/router";
import treemap from "@/components/main"

Vue.use(Router)
Vue.use(treemap)
Vue.config.productionTip = false

new Vue({
    router,
    render: h => h(App),
}).$mount('#app')
