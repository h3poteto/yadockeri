import Projects, { ProjectsModule } from './projects'

export type PagesModule = {
  projects: ProjectsModule
}

export default {
  namespaced: true,
  modules: {
    projects: Projects,
  },
}
