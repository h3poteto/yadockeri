<template>
  <div>
    <ValidationObserver ref="observer">
      <el-form ref="form" label-width="150px" slot-scope="{ validate }">
        <ValidationProvider rules="required" name="selectedBranch">
          <el-form-item
            label="Branch"
            slot-scope="{ errors }"
            :error="errors[0]"
          >
            <el-select
              v-model="selectedBranch"
              filterable
              placeholder="Select"
              :loading="loadingGithubBranch"
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
        </ValidationProvider>
        <el-form-item>
          <el-button type="primary" @click="validate().then(onSubmit)"
            >Create</el-button
          >
        </el-form-item>
      </el-form>
    </ValidationObserver>
  </div>
</template>

<script>
import Vue from 'vue'
import { createNamespacedHelpers } from 'vuex'
import { ValidationObserver, ValidationProvider } from 'vee-validate'

const { mapState, mapActions } = createNamespacedHelpers(
  'pages/projects/branches/new'
)

export default Vue.extend({
  name: 'new',
  components: {
    ValidationObserver,
    ValidationProvider,
  },
  props: ['project_id'],
  computed: {
    ...mapState(['loadingGithubBranch', 'githubBranches']),
    selectedBranch: {
      get() {
        return this.$store.state.pages.projects.branches.new.selectedBranch
      },
      set(value) {
        this.changeBranch(value)
      },
    },
  },
  async mounted() {
    const project = await this.fetchProject(this.project_id).catch(err => {
      console.error(err)
      this.$message({
        message: 'Failed to get the project',
        type: 'error',
      })
    })
    await this.fetchGithubBranches({
      owner: project.repository_owner,
      repo: project.repository_name,
    }).catch(err => {
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
      await this.submit().catch(err => {
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
