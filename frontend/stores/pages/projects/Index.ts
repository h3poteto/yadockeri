import { ActionContext } from 'vuex'
import axios from 'axios'
import { RootState } from '@/store'
import Branches, { BranchesModule } from './branches'

type Context = ActionContext<ProjectsIndexState, RootState>

type Project = {
  id: number
  userId: number
  title: string
  baseURL: string
  repositoryOwner: string
  repositoryName: string
  helmRepositoryUrl: string
  helmDirectoryName: string
  namespace: string
}

export type ProjectsIndexState = {
  projects: Array<Project>
}

const initialState = (): ProjectsIndexState => ({
  projects: [],
})

const actions = {
  async fetchProjects({ commit }: Context) {
    const response = await axios.get('/api/v1/projects')
    commit(MUTATION_TYPE.SET_PROJECTS, response.data)
  },
}

const MUTATION_TYPE = {
  SET_PROJECTS: 'SET_PROJECTS',
}

const mutations = {
  [MUTATION_TYPE.SET_PROJECTS](
    state: ProjectsIndexState,
    projects: Array<Project>
  ) {
    state.projects = projects
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
}