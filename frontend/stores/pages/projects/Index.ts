import { Module, ActionTree, MutationTree } from 'vuex'
import { RootState } from '@/store'
import Yadockeri, { Project } from '@/lib/client'
import Branches, { BranchesModule } from './branches'

export type ProjectsIndexState = {
  projects: Array<Project>
}

const initialState = (): ProjectsIndexState => ({
  projects: [],
})

const MUTATION_TYPE = {
  SET_PROJECTS: 'SET_PROJECTS',
}

const mutations: MutationTree<ProjectsIndexState> = {
  [MUTATION_TYPE.SET_PROJECTS]: (
    state: ProjectsIndexState,
    projects: Array<Project>
  ) => {
    state.projects = projects
  },
}

const actions: ActionTree<ProjectsIndexState, RootState> = {
  fetchProjects: async ({ commit }) => {
    const response = await Yadockeri.get<Project>('/api/v1/projects')
    commit(MUTATION_TYPE.SET_PROJECTS, response.data)
  },
}

type ProjectsIndexModule = {
  Branches: BranchesModule
}

export type ProjectsIndexModuleState = ProjectsIndexState & ProjectsIndexModule

export default {
  namespaced: true,
  modules: {
    branches: Branches,
  },
  state: initialState(),
  mutations,
  actions,
} as Module<ProjectsIndexState, RootState>
