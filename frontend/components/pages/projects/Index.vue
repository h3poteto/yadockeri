<template>
  <div class="dashboard">
    <el-container direction="vertical">
      <el-row>
        <el-col>
          <div class="new-repository">
            <el-button type="primary" @click="handleOpen"
              >New Repository</el-button
            >
          </div>
        </el-col>
      </el-row>
      <el-row>
        <el-col class="repository-card-area">
          <router-link
            v-for="project in projects"
            :to="`/projects/${project.id}/branches`"
            :key="project.id"
          >
            <el-card shadow="hover" class="repository-card">{{
              project.title
            }}</el-card>
          </router-link>
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
    handleOpen() {
      this.$router.push('/projects/new')
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
  width: 200px;
  margin: 10px 10px 10px 0;
  text-align: center;
}

.new-repository {
  display: block;
  text-align: right;
  margin: 12px 8px 24px auto;
}
</style>
