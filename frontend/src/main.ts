import "@babel/polyfill";
import "mutationobserver-shim";
import Vue from "vue";
import "./plugins/bootstrap-vue";
import App from "./App.vue";
import router from "./router";
import { HttpService } from './plugins/http';
import { library } from '@fortawesome/fontawesome-svg-core'
import { faPlusSquare } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

Vue.config.productionTip = false;

library.add(faPlusSquare)

Vue.component('font-awesome-icon', FontAwesomeIcon)
Vue.use(HttpService);

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
