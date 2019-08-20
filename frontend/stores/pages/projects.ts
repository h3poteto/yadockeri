import Branches, { BranchesModule } from './projects/branches'
import Index, { ProjectsIndexModuleState } from './projects/Index'
import New, { ProjectsNewModuleState } from './projects/New'
import Edit, { ProjectsEditModuleState } from './projects/Edit'

export type ProjectsModule = {
  Index: ProjectsIndexModuleState
  New: ProjectsNewModuleState
  Edit: ProjectsEditModuleState
  Branches: BranchesModule
}

export default {
  namespaced: true,
  modules: {
    index: Index,
    new: New,
    edit: Edit,
    branches: Branches,
  },
}
