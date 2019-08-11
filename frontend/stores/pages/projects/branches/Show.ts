import { MutationTree, ActionTree, Module } from 'vuex'
import Yadockeri, { Branch, Status, AuthenticationError } from '@/lib/client'
import { RootState } from '@/store'
import router from '@/router'

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
    try {
      const response = await Yadockeri.get<Branch>(
        `/api/v1/projects/${projectID}/branches/${id}`
      )
      commit(MUTATION_TYPES.SET_BRANCH, response.data)
    } catch (err) {
      if (err instanceof AuthenticationError) {
        window.location.href = '/login'
      } else {
        throw err
      }
    }
  },
  deploy: async ({ commit }, { projectID, id }) => {
    try {
      const response = await Yadockeri.patch<Status>(
        `/api/v1/projects/${projectID}/branches/${id}/deploy`
      )
      commit(MUTATION_TYPES.UPDATE_STATUS, response.data.status)
    } catch (err) {
      if (err instanceof AuthenticationError) {
        window.location.href = '/login'
      } else {
        throw err
      }
    }
  },
  fetchStatus: async ({ commit }, { projectID, id }) => {
    try {
      const response = await Yadockeri.get<Status>(
        `/api/v1/projects/${projectID}/branches/${id}/status`
      )
      commit(MUTATION_TYPES.UPDATE_STATUS, response.data.status)
    } catch (err) {
      if (err instanceof AuthenticationError) {
        window.location.href = '/login'
      } else {
        commit(MUTATION_TYPES.UPDATE_STATUS, 'release does not exist')
      }
    }
  },
  delete: async (_, { projectID, id }) => {
    try {
      await Yadockeri.delete(`/api/v1/projects/${projectID}/branches/${id}`)
    } catch (err) {
      if (err instanceof AuthenticationError) {
        window.location.href = '/login'
      } else {
        throw err
      }
    }
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
