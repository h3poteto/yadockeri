import { MutationTree, ActionTree, Module } from 'vuex'
import axios from 'axios'
import { RootState } from '@/store'
import router from '../../../../router'

type Branch = {
  id: number
  project_id: number
  user_id: number
  name: string
  url: string
  stack_name: string
}

type State = {
  branch: Branch
  status: string
}

const initalState = (): State => ({
  branch: {
    id: 0,
    project_id: 0,
    user_id: 0,
    name: '',
    url: '',
    stack_name: '',
  },
  status: '',
})

const MUTATION_TYPES = {
  SET_BRANCH: 'SET_BRANCH',
  UPDATE_STATUS: 'UPDATE_STATUS',
}

const mutations: MutationTree<State> = {
  [MUTATION_TYPES.SET_BRANCH]: (state, branch: Branch) => {
    state.branch = branch
  },
  [MUTATION_TYPES.UPDATE_STATUS]: (state, status: string) => {
    state.status = status
  },
}

const actions: ActionTree<State, RootState> = {
  fetchBranch: async ({ commit }, { projectID, id }) => {
    const response = await axios.get<Branch>(
      `/api/v1/projects/${projectID}/branches/${id}`
    )
    commit(MUTATION_TYPES.SET_BRANCH, response.data)
  },
  deploy: async ({ commit }, { projectID, id }) => {
    const response = await axios.patch(
      `/api/v1/projects/${projectID}/branches/${id}/deploy`
    )
    commit(MUTATION_TYPES.UPDATE_STATUS, response.data.status)
  },
  fetchStatus: async ({ commit }, { projectID, id }) => {
    try {
      const response = await axios.get(
        `/api/v1/projects/${projectID}/branches/${id}/status`
      )
      commit(MUTATION_TYPES.UPDATE_STATUS, response.data.status)
    } catch (err) {
      console.error(err)
      commit(MUTATION_TYPES.UPDATE_STATUS, 'release does not exist')
    }
  },
  delete: async (_, { projectID, id }) => {
    await axios.delete(`/api/v1/projects/${projectID}/branches/${id}`)
    router.push(`/projects/${projectID}/branches`)
  },
}

export type ShowModuleState = State

export default {
  namespaced: true,
  state: initalState(),
  mutations: mutations,
  actions: actions,
} as Module<State, RootState>
