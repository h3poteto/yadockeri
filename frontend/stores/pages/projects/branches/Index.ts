import { MutationTree, ActionTree, Module } from 'vuex'
import Yadockeri, { Branch } from '@/lib/client'
import { RootState } from '@/store'

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
    const response = await Yadockeri.get<Array<Branch>>(
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
