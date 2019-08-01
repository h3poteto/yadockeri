import Index, { IndexModuleState } from './branches/Index'
import New, { NewModuleState } from './branches/New'
import Show, { ShowModuleState } from './branches/Show'

export type BranchesModule = {
  Index: IndexModuleState
  New: NewModuleState
  Show: ShowModuleState
}

export default {
  namespaced: true,
  modules: {
    index: Index,
    new: New,
    show: Show,
  },
}
