import { MutationTree, ActionTree, Module } from 'vuex'
import axios from 'axios'
import { RootState } from '@/store'

type Branch = {
  id: number
  project_id: number
  user_id: number
  name: string
  url: string
  stack_name: string
}

type State = {
  branches: Array<Branch>
}

const initialState = (): State => ({
  branches: [],
})

const MUTATION_TYPES = {
  SET_BRANCHES: 'SET_BRANCHES',
}

const mutations: MutationTree<State> = {
  [MUTATION_TYPES.SET_BRANCHES]: (state, branches: Array<Branch>) => {
    state.branches = branches
  },
}

const actions: ActionTree<State, RootState> = {
  fetchBranches: async ({ commit }, projectId: number) => {
    const response = await axios.get<Array<Branch>>(
      `/api/v1/projects/${projectId}/branches`
    )
    commit(MUTATION_TYPES.SET_BRANCHES, response.data)
  },
}

export type IndexModuleState = State

export default {
  namespaced: true,
  state: initialState(),
  mutations: mutations,
  actions: actions,
} as Module<State, RootState>
