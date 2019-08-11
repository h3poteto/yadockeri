import { Module, MutationTree, ActionTree } from 'vuex'
import Yadockeri, { Project, GitHubBranch } from '@/lib/client'
import { RootState } from '@/store'
import router from '@/router'

type State = {
  project: Project | null
  loadingGtihubBranch: boolean
  githubBranches: Array<GitHubBranch>
  selectedBranch: string
}

const initialState = (): State => ({
  project: null,
  loadingGtihubBranch: false,
  githubBranches: [],
  selectedBranch: '',
})

const MUTATION_TYPES = {
  SET_PROJECT: 'SET_PROJECT',
  CHANGE_LOADING_GITHUB_BRANCH: 'CHANGE_LOADING_GITHUB_BRANCH',
  SET_GITHUB_BRANCHES: 'SET_GITHUB_BRANCHES',
  SET_SELECTED_BRANCH: 'SET_SELECTED_BRANCH',
}

const mutations: MutationTree<State> = {
  [MUTATION_TYPES.SET_PROJECT]: (state, project: Project) => {
    state.project = project
  },
  [MUTATION_TYPES.CHANGE_LOADING_GITHUB_BRANCH]: (state, loading: boolean) => {
    state.loadingGtihubBranch = loading
  },
  [MUTATION_TYPES.SET_GITHUB_BRANCHES]: (
    state,
    branches: Array<GitHubBranch>
  ) => {
    state.githubBranches = branches
  },
  [MUTATION_TYPES.SET_SELECTED_BRANCH]: (state, branch: string) => {
    state.selectedBranch = branch
  },
}

const actions: ActionTree<State, RootState> = {
  fetchProject: async ({ commit }, id: number): Promise<Project> => {
    const response = await Yadockeri.get<Project>(`/api/v1/projects/${id}`)
    commit(MUTATION_TYPES.SET_PROJECT, response.data)
    return response.data
  },
  fetchGithubBranches: async ({ commit }, { owner, repo }) => {
    commit(MUTATION_TYPES.CHANGE_LOADING_GITHUB_BRANCH, true)
    const response = await Yadockeri.get<Array<GitHubBranch>>(
      `/api/v1/github/branches`,
      {
        params: {
          owner: owner,
          repo: repo,
        },
      }
    ).finally(() => {
      commit(MUTATION_TYPES.CHANGE_LOADING_GITHUB_BRANCH, false)
    })
    commit(MUTATION_TYPES.SET_GITHUB_BRANCHES, response.data)
  },
  changeBranch: ({ commit }, name: string) => {
    commit(MUTATION_TYPES.SET_SELECTED_BRANCH, name)
  },
  submit: async ({ state }) => {
    if (state.selectedBranch.length == 0) {
      throw new Error('branch is blank')
    }
    await Yadockeri.post(`/api/v1/projects/${state.project!.id}/branches`, {
      name: state.selectedBranch,
    })
    return router.push(`/projects/${state.project!.id}/branches`)
  },
}

export type NewModuleState = State

export default {
  namespaced: true,
  state: initialState(),
  mutations: mutations,
  actions: actions,
} as Module<State, RootState>
