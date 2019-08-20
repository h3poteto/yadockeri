import { Module, MutationTree, ActionTree } from 'vuex'
import Yadockeri, {
  Project,
  AuthenticationError,
  OverrideValue,
} from '@/lib/client'
import { RootState } from '@/store'
import router from '@/router'

export type ProjectsEditState = {
  project: Project | null
  helmDirectory: string
  baseURL: string
  namespace: string
  values: Array<OverrideValue>
}

const initialState = (): ProjectsEditState => ({
  project: null,
  helmDirectory: '',
  baseURL: '',
  namespace: '',
  values: [],
})

const MUTATION_TYPES = {
  SET_PROJECT: 'SET_PROJECT',
  SET_HELM_DIRECTORY: 'SET_HELM_DIRECTORY',
  SET_BASE_URL: 'SET_BASE_URL',
  SET_NAMESPACE: 'SET_NAMESPACE',
  UPDATE_VALUES_KEY: 'UPDATE_VALUES_KEY',
  UPDATE_VALUES_OPTION: 'UPDATE_VALUES_OPTION',
  ADD_VALUE: 'ADD_VALUE',
  REMOVE_VALUE: 'REMOVE_VALUE',
}

const mutations: MutationTree<ProjectsEditState> = {
  [MUTATION_TYPES.SET_PROJECT]: (state, project: Project) => {
    state.project = project
    state.helmDirectory = project.helm_directory_name
    state.baseURL = project.base_url
    state.namespace = project.namespace
    state.values = project.values
  },
  [MUTATION_TYPES.SET_HELM_DIRECTORY]: (state, directory: string) => {
    state.helmDirectory = directory
  },
  [MUTATION_TYPES.SET_BASE_URL]: (state, url: string) => {
    state.baseURL = url
  },
  [MUTATION_TYPES.SET_NAMESPACE]: (state, namespace: string) => {
    state.namespace = namespace
  },
  [MUTATION_TYPES.UPDATE_VALUES_KEY]: (state, { newKey, key }) => {
    state.values = state.values.map(v => {
      if (v.key === key) {
        return { key: newKey, value: v.value }
      }
      return v
    })
  },
  [MUTATION_TYPES.UPDATE_VALUES_OPTION]: (state, { newOption, key }) => {
    state.values = state.values.map(v => {
      if (v.key === key) {
        return { key: key, value: newOption }
      }
      return v
    })
  },
  [MUTATION_TYPES.ADD_VALUE]: state => {
    state.values = state.values.concat([{ key: '', value: '' }])
  },
  [MUTATION_TYPES.REMOVE_VALUE]: (state, key: string) => {
    state.values = state.values.filter(v => v.key !== key)
  },
}

const actions: ActionTree<ProjectsEditState, RootState> = {
  fetchProject: async ({ commit }, id: string) => {
    try {
      const response = await Yadockeri.get<Project>(`/api/v1/projects/${id}`)
      commit(MUTATION_TYPES.SET_PROJECT, response.data)
    } catch (err) {
      if (err instanceof AuthenticationError) {
        window.location.href = '/login'
      } else {
        throw err
      }
    }
  },
  changeHelmDirectory: ({ commit }, directory: string) => {
    commit(MUTATION_TYPES.SET_HELM_DIRECTORY, directory)
  },
  changeBaseULR: ({ commit }, url: string) => {
    commit(MUTATION_TYPES.SET_BASE_URL, url)
  },
  changeNamespace: ({ commit }, namespace: string) => {
    commit(MUTATION_TYPES.SET_NAMESPACE, namespace)
  },
  submit: async ({ state }, id: string) => {
    try {
      await Yadockeri.patch<Project>(`/api/v1/projects/${id}`, {
        base_url: state.baseURL,
        helm_directory_name: state.helmDirectory,
        namespace: state.namespace,
        value_options: state.values,
      })
      router.push('/')
    } catch (err) {
      if (err instanceof AuthenticationError) {
        window.location.href = '/login'
      } else {
        throw err
      }
    }
  },
  addValue: ({ commit }) => {
    commit(MUTATION_TYPES.ADD_VALUE)
  },
  removeValue: ({ commit }, id) => {
    commit(MUTATION_TYPES.REMOVE_VALUE, id)
  },
  updateValueKey: ({ commit }, value) => {
    commit(MUTATION_TYPES.UPDATE_VALUES_KEY, value)
  },
  updateValueOption: ({ commit }, value) => {
    commit(MUTATION_TYPES.UPDATE_VALUES_OPTION, value)
  },
}

export type ProjectsEditModuleState = ProjectsEditState

export default {
  namespaced: true,
  state: initialState(),
  mutations: mutations,
  actions: actions,
} as Module<ProjectsEditState, RootState>
