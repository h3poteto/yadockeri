import Vue from 'vue'
import { loadProgressBar } from 'axios-progress-bar'
import 'axios-progress-bar/dist/nprogress.css'
import VeeValidate from 'vee-validate'

import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'

Vue.config.productionTip = false

// The default filed name is conflicted, so we have to override it.
// https://github.com/ElemeFE/element/issues/4720#issuecomment-299617405
const VeeConfig = {
  inject: false,
  fieldsBagName: 'veefields',
}

Vue.use(VeeValidate, VeeConfig)

loadProgressBar()

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
