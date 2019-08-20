import Vue from 'vue'
import Router from 'vue-router'
import Projects from './components/pages/projects/Index.vue'
import NewProject from './components/pages/projects/New.vue'
import EditProject from './components/pages/projects/Edit.vue'
import Branches from './components/pages/projects/branches/Index.vue'
import NewBranch from './components/pages/projects/branches/New.vue'
import ShowBranch from './components/pages/projects/branches/Show.vue'

Vue.use(Router)

export default new Router({
  mode: 'history',
  base: '/',
  routes: [
    {
      path: '/',
      name: 'Dashboard',
      component: Projects,
    },
    {
      path: '/projects/',
      name: 'Projects',
      component: Projects,
      children: [
        {
          path: 'new',
          name: 'NewProject',
          component: NewProject,
        },
        {
          path: ':project_id/edit',
          name: 'EditProject',
          component: EditProject,
          props: true,
        },
      ],
    },
    {
      path: '/projects/:project_id/branches',
      name: 'Branches',
      component: Branches,
      props: true,
    },
    {
      path: '/projects/:project_id/branches/new',
      name: 'NewBranch',
      component: NewBranch,
      props: true,
    },
    {
      path: '/projects/:project_id/branches/:branch_id',
      name: 'ShowBranch',
      component: ShowBranch,
      props: true,
    },
  ],
})
