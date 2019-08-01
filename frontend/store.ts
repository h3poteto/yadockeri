import Vue from 'vue'
import Vuex from 'vuex'
import createLogger from 'vuex/dist/logger'

import pages, { PagesModule } from './stores/pages/Index'

Vue.use(Vuex)

export type RootState = {
  pages: PagesModule
}

export default new Vuex.Store<RootState>({
  strict: true,
  plugins: process.env.NODE_ENV !== 'production' ? [createLogger({})] : [],
  modules: {
    pages,
  },
})
