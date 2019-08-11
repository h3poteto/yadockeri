<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="18" :offset="2" class="row">
        <h3 class="stack-name">
          {{ branch.stack_name }}
        </h3>
        <template v-if="branch.url.length > 0">
          <el-link :href="branch.url" target="_blank" type="primary">{{
            branch.url
          }}</el-link>
        </template>
        <div class="operation">
          <el-button type="danger" @click="startDelete">Delete</el-button>
          <el-button type="primary" @click="startDeploy">Deploy </el-button>
        </div>
        <el-divider>helm status</el-divider>
        <div class="helm-status">
          <pre>
        {{ status }}
      </pre
          >
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { createNamespacedHelpers } from 'vuex'

const { mapState, mapActions } = createNamespacedHelpers(
  'pages/projects/branches/show'
)

export default Vue.extend({
  name: 'show',
  props: ['project_id', 'branch_id'],
  computed: {
    ...mapState(['branch', 'status']),
  },
  async mounted() {
    await this.fetchBranch({
      projectID: this.project_id,
      id: this.branch_id,
    }).catch((err: Error) => {
      console.error(err)
      this.$message({
        message: 'Failed to get the branch',
        type: 'error',
      })
    })
    await this.fetchStatus({
      projectID: this.project_id,
      id: this.branch_id,
    })
  },
  methods: {
    ...mapActions(['fetchBranch', 'deploy', 'fetchStatus', 'delete']),
    async startDeploy() {
      await this.deploy({
        projectID: this.project_id,
        id: this.branch_id,
      }).catch((err: Error) => {
        console.error(err)
        this.$message({
          message: 'Failed to get deploy the branch',
          type: 'error',
        })
      })
    },
    startDelete() {
      this.$confirm(
        'This action cannot be undone. This will permanetnly delete this release.',
        'Are you absolutely sure?',
        {
          confirmButtonText: 'OK',
          cancelButtonText: 'Cancel',
          type: 'warning',
        }
      ).then(() => {
        this.delete({ projectID: this.project_id, id: this.branch_id }).catch(
          (err: Error) => {
            console.error(err)
            this.$message({
              message: 'Failed to delete this branch',
              type: 'error',
            })
          }
        )
      })
    },
  },
})
</script>

<style lang="scss" scoped>
.row {
  background-color: #fff;
  padding: 20px 12px;
}

.stack-name {
  text-decoration: underline;
}

.operation {
  display: block;
  text-align: right;
  margin: 12px 8px 24px auto;
}

.helm-status {
  background-color: #000;
  color: #e4e7ed;
  padding: 24px 8px;
}
</style>
