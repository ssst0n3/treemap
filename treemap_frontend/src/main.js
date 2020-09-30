import Vue from 'vue'
import App from './App.vue'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import restful from 'vue-lightweight_restful'
import consts from "@/utils/consts";
import Router from 'vue-router'
import router from "@/router";
import jm from 'vue-jsmind'


restful.api.initClient(consts.BaseUrl)

Vue.use(Router)
Vue.use(jm)
Vue.use(BootstrapVue)
Vue.config.productionTip = false

new Vue({
    router,
    render: h => h(App),
}).$mount('#app')
