<template>
  <div>
    <el-form ref="form" label-width="150px">
      <el-form-item label="Branch">
        <el-select
          :value="selectedBranch"
          filterable
          placeholder="Select"
          :loading="loadingGithubBranch"
          @change="changeBranch"
        >
          <el-option
            v-for="item in githubBranches"
            :key="item.name"
            :label="item.name"
            :value="item.name"
          >
          </el-option>
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">Create</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { createNamespacedHelpers } from 'vuex'

const { mapState, mapActions } = createNamespacedHelpers(
  'pages/projects/branches/new'
)

export default Vue.extend({
  name: 'new',
  props: ['project_id'],
  computed: {
    ...mapState(['loadingGithubBranch', 'githubBranches', 'selectedBranch']),
  },
  async mounted() {
    const project = await this.fetchProject(this.project_id).catch(
      (err: Error) => {
        console.error(err)
        this.$message({
          message: 'Failed to get the project',
          type: 'error',
        })
      }
    )
    await this.fetchGithubBranches({
      owner: project.repository_owner,
      repo: project.repository_name,
    }).catch((err: Error) => {
      console.error(err)
      this.$message({
        message: 'Failed to get github branches',
        type: 'error',
      })
    })
  },
  methods: {
    ...mapActions([
      'fetchProject',
      'fetchGithubBranches',
      'changeBranch',
      'submit',
    ]),
    async onSubmit() {
      await this.submit().catch((err: Error) => {
        console.error(err)
        this.$message({
          message: 'Failed to create a new branch',
          type: 'error',
        })
      })
    },
  },
})
</script>

<style lang="scss" scoped></style>
