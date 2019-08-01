import Branches, { BranchesModule } from './projects/branches'
import Index, { ProjectsIndexModuleState } from './projects/Index'
import New, { ProjectsNewModuleState } from './projects/New'

export type ProjectsModule = {
  Index: ProjectsIndexModuleState
  New: ProjectsNewModuleState
  Branches: BranchesModule
}

export default {
  namespaced: true,
  modules: {
    index: Index,
    new: New,
    branches: Branches,
  },
}
