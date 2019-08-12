<template>
  <div class="dashboard">
    <div class="new-branch">
      <router-link :to="`/projects/${project_id}/branches/new`">
        <el-button type="primary">
          New
        </el-button>
      </router-link>
    </div>
    <el-table :data="branches">
      <el-table-column label="Branch" prop="name"></el-table-column>
      <el-table-column label="URL" prop="url"> </el-table-column>
      <el-table-column label="Operations" width="300">
        <template slot-scope="scope">
          <router-link :to="`/projects/${project_id}/branches/${scope.row.id}`">
            <el-button size="small" type="text">Details</el-button>
          </router-link>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { createNamespacedHelpers } from 'vuex'
const { mapState, mapActions } = createNamespacedHelpers(
  'pages/projects/branches/index'
)

export default Vue.extend({
  name: 'Branches',
  props: ['project_id'],
  computed: {
    ...mapState(['branches']),
  },
  async mounted() {
    await this.fetchBranches(this.project_id).catch((err: Error) => {
      console.error(err)
      this.$message({
        message: 'Failed to get branches',
        type: 'error',
      })
    })
  },
  methods: {
    ...mapActions(['fetchBranches']),
  },
})
</script>

<style lang="scss" scoped>
.new-branch {
  display: block;
  text-align: right;
  margin: 12px 8px 24px auto;
}
</style>
