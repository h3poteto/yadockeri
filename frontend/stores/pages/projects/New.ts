import { Module, MutationTree, ActionTree } from 'vuex'
import Yadockeri, {
  GitHubRepo,
  Project,
  AuthenticationError,
} from '@/lib/client'
import { RootState } from '@/store'
import router from '@/router'

export type ProjectsNewState = {
  githubRepos: Array<GitHubRepo>
  selectedRepository: string
  selectedRepositoryOwner: string
  selectedHelmRepositoryUrl: string
  helmDirectory: string
  baseURL: string
  namespace: string
  loadingGithubRepo: boolean
  loadingCreateProject: boolean
  values: Array<OverrideValue>
}

export type OverrideValue = {
  key: string
  value: string
}

const initialState = (): ProjectsNewState => ({
  githubRepos: [],
  selectedRepository: '',
  selectedRepositoryOwner: '',
  selectedHelmRepositoryUrl: '',
  helmDirectory: '',
  baseURL: '',
  namespace: 'default',
  loadingGithubRepo: false,
  loadingCreateProject: false,
  values: [{ key: '', value: '' }],
})

const MUTATION_TYPES = {
  SET_GITHUB_REPOS: 'SET_GITHUB_REPOS',
  SET_SELECTED_REPOSITORY: 'SET_SELECTED_REPOSITORY',
  SET_SELECTED_REPOSITORY_OWNER: 'SET_SELECTED_REPOSITORY_OWNER',
  SET_SELECTED_HELM_REPOSITORY_URL: 'SET_SELECTED_HELM_REPOSITORY_URL',
  SET_HELM_DIRECTORY: 'SET_HELM_DIRECTORY',
  SET_BASE_URL: 'SET_BASE_URL',
  SET_NAMESPACE: 'SET_NAMESPACE',
  TOGGLE_LOADING_GITHUB_REPO: 'TOGGLE_LOADING_GITHUB_REPO',
  TOGGLE_LOADING_CREATE_PROJECT: 'TOGGLE_LOADING_CREATE_PROJECT',
  UPDATE_VALUES_KEY: 'UPDATE_VALUES_KEY',
  UPDATE_VALUES_OPTION: 'UPDATE_VALUES_OPTION',
  ADD_VALUE: 'ADD_VALUE',
  REMOVE_VALUE: 'REMOVE_VALUE',
}

const mutations: MutationTree<ProjectsNewState> = {
  [MUTATION_TYPES.SET_GITHUB_REPOS]: (
    state,
    githubRepos: Array<GitHubRepo>
  ) => {
    state.githubRepos = githubRepos
  },
  [MUTATION_TYPES.SET_SELECTED_REPOSITORY]: (
    state,
    selectedRepository: string
  ) => {
    state.selectedRepository = selectedRepository
  },
  [MUTATION_TYPES.SET_SELECTED_REPOSITORY_OWNER]: (
    state,
    selectedRepositoryOwner: string
  ) => {
    state.selectedRepositoryOwner = selectedRepositoryOwner
  },
  [MUTATION_TYPES.SET_SELECTED_HELM_REPOSITORY_URL]: (
    state,
    selectedHelmRepositoryUrl: string
  ) => {
    state.selectedHelmRepositoryUrl = selectedHelmRepositoryUrl
  },
  [MUTATION_TYPES.SET_HELM_DIRECTORY]: (state, helmDirectory: string) => {
    state.helmDirectory = helmDirectory
  },
  [MUTATION_TYPES.SET_BASE_URL]: (state, url: string) => {
    state.baseURL = url
  },
  [MUTATION_TYPES.SET_NAMESPACE]: (state, namespace: string) => {
    state.namespace = namespace
  },
  [MUTATION_TYPES.TOGGLE_LOADING_GITHUB_REPO]: state => {
    state.loadingGithubRepo = !state.loadingGithubRepo
  },
  [MUTATION_TYPES.TOGGLE_LOADING_CREATE_PROJECT]: state => {
    state.loadingCreateProject = !state.loadingCreateProject
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

const actions: ActionTree<ProjectsNewState, RootState> = {
  fetchGithubRepos: async ({ commit }) => {
    commit(MUTATION_TYPES.TOGGLE_LOADING_GITHUB_REPO)
    try {
      const response = await Yadockeri.get<Array<GitHubRepo>>(
        '/api/v1/github/repos'
      )
      commit(MUTATION_TYPES.SET_GITHUB_REPOS, response.data)
    } catch (err) {
      if (err instanceof AuthenticationError) {
        window.location.href = '/login'
      } else {
        throw err
      }
    } finally {
      commit(MUTATION_TYPES.TOGGLE_LOADING_GITHUB_REPO)
    }
  },
  changeRepository: ({ commit }, item: GitHubRepo) => {
    commit(MUTATION_TYPES.SET_SELECTED_REPOSITORY, item.name)
    commit(MUTATION_TYPES.SET_SELECTED_REPOSITORY_OWNER, item.owner.login)
  },
  changeHelmRepository: ({ commit }, selectedHelmRepositoryUrl: string) => {
    commit(
      MUTATION_TYPES.SET_SELECTED_HELM_REPOSITORY_URL,
      selectedHelmRepositoryUrl
    )
  },
  setHelmDirectory: ({ commit }, helmDirectory: string) => {
    commit(MUTATION_TYPES.SET_HELM_DIRECTORY, helmDirectory)
  },
  setBaseURL: ({ commit }, url: string) => {
    commit(MUTATION_TYPES.SET_BASE_URL, url)
  },
  setNamespace: ({ commit }, namespace: string) => {
    commit(MUTATION_TYPES.SET_NAMESPACE, namespace)
  },
  submit: async ({ commit, state, dispatch }) => {
    commit(MUTATION_TYPES.TOGGLE_LOADING_CREATE_PROJECT)
    try {
      await Yadockeri.post<Project>('/api/v1/projects', {
        title: state.selectedRepository,
        base_url: state.baseURL,
        repository_owner: state.selectedRepositoryOwner,
        repository_name: state.selectedRepository,
        helm_repository_url: state.selectedHelmRepositoryUrl,
        helm_directory_name: state.helmDirectory,
        namespace: state.namespace,
        value_options: state.values,
      })
      dispatch('pages/projects/index/fetchProjects', {}, { root: true })
      router.push('/')
    } catch (err) {
      if (err instanceof AuthenticationError) {
        window.location.href = '/login'
      } else {
        throw err
      }
    } finally {
      commit(MUTATION_TYPES.TOGGLE_LOADING_CREATE_PROJECT)
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

export type ProjectsNewModuleState = ProjectsNewState

export default {
  namespaced: true,
  state: initialState(),
  mutations: mutations,
  actions: actions,
} as Module<ProjectsNewState, RootState>
