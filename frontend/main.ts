import Vue from 'vue'
import { loadProgressBar } from 'axios-progress-bar'
import 'axios-progress-bar/dist/nprogress.css'

import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'

Vue.config.productionTip = false
loadProgressBar()

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
