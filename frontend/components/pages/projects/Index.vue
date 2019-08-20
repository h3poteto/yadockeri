<template>
  <div class="dashboard">
    <el-container direction="vertical">
      <el-row>
        <el-col>
          <div class="new-repository">
            <el-button type="primary" @click="openNew"
              >New Repository</el-button
            >
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col class="repository-card-area">
          <el-card
            shadow="hover"
            class="repository-card"
            v-for="project in projects"
            :key="project.id"
          >
            <div class="repository" @click="openProject(project.id)">
              {{ project.title }}
            </div>

            <el-button
              type="text"
              icon="el-icon-edit"
              class="edit-button"
              @click="openEdit(project.id)"
            ></el-button>
          </el-card>
        </el-col>
      </el-row>
    </el-container>
    <router-view></router-view>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { createNamespacedHelpers } from 'vuex'

const { mapState, mapActions } = createNamespacedHelpers('pages/projects/index')

export default Vue.extend({
  name: 'index',
  computed: {
    ...mapState(['projects']),
  },
  methods: {
    ...mapActions(['fetchProjects']),
    openNew() {
      this.$router.push('/projects/new')
    },
    openProject(id: number) {
      this.$router.push(`/projects/${id}/branches`)
    },
    openEdit(id: number) {
      this.$router.push(`/projects/${id}/edit`)
    },
  },
  async mounted() {
    await this.fetchProjects().catch((err: Error) => {
      console.error(err)
      this.$message({
        message: 'Failed to get projects',
        type: 'error',
      })
    })
  },
})
</script>

<style lang="scss" scoped>
.repository-card-area {
  display: flex;
}
.repository-card {
  position: relative;
  width: 200px;
  margin: 10px 10px 10px 0;
  text-align: center;

  .edit-button {
    position: absolute;
    top: 0;
    right: 4px;
  }
}

.new-repository {
  display: block;
  text-align: right;
  margin: 12px 8px 24px auto;
}
</style>
