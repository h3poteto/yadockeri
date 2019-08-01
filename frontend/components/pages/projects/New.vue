<template>
  <div>
    <el-dialog
      title="New Repository"
      :visible.sync="visible"
      width="40%"
      :before-close="handleClose"
      v-loading="loadingCreateProject"
    >
      <el-form label-width="150px">
        <el-form-item label="Repository">
          <el-select
            :value="selectedRepository"
            filterable
            placeholder="Select"
            :loading="loadingGithubRepo"
            @change="changeRepository"
          >
            <el-option
              v-for="item in githubRepos"
              :key="item.repositoryUrl"
              :label="item.fullName"
              :value="item"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Helm Repository">
          <el-select
            :value="selectedHelmRepositoryUrl"
            filterable
            placeholder="Select"
            :loading="loadingGithubRepo"
            @change="changeHelmRepository"
          >
            <el-option
              v-for="item in githubRepos"
              :key="item.repositoryUrl"
              :label="item.fullName"
              :value="item.repositoryUrl"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="Helm Directory">
          <el-input
            :value="helmDirectory"
            placeholder="Helm Directory"
            @input="setHelmDirectory"
          ></el-input>
        </el-form-item>
        <el-form-item label="Base URL">
          <el-input
            :value="baseURL"
            placeholder="https://dev.scout.lapras.com"
            @input="setBaseURL"
          ></el-input>
        </el-form-item>
        <el-form-item label="Namespace">
          <el-input
            :value="namespace"
            placeholder="default"
            @input="setNamespace"
          ></el-input>
        </el-form-item>

        <el-form-item>
          <ul class="value-list">
            <li class="value" v-for="(value, id) in values" :key="id">
              <el-input
                :value="value.key"
                @input="v => updateValueKey({ newKey: v, key: value.key })"
                placeholder="Override key"
                size="small"
              ></el-input>
              <div>:</div>
              <el-input
                :value="value.value"
                @input="
                  v => updateValueOption({ newOption: v, key: value.key })
                "
                placeholder="Override value"
                size="small"
              ></el-input>
              <el-button
                type="text"
                icon="el-icon-close"
                @click="removeValue(value.key)"
              ></el-button>
            </li>
          </ul>
          <div class="add-value-area">
            <el-button
              class="add-value-option"
              type="info"
              size="small"
              @click="addValue"
              >Add value</el-button
            >
          </div>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="onSubmit">Create</el-button>
          <el-button @click="handleClose">Cancel</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { createNamespacedHelpers } from 'vuex'

const { mapState, mapActions } = createNamespacedHelpers('pages/projects/new')

export default Vue.extend({
  name: 'new-project',
  data() {
    return {
      visible: true,
    }
  },
  computed: {
    ...mapState([
      'githubRepos',
      'selectedRepository',
      'selectedHelmRepositoryUrl',
      'helmDirectory',
      'baseURL',
      'namespace',
      'loadingGithubRepo',
      'loadingCreateProject',
      'values',
    ]),
  },
  created() {
    this.fetchGithubRepos()
  },
  methods: {
    ...mapActions([
      'fetchGithubRepos',
      'changeRepository',
      'changeHelmRepository',
      'setHelmDirectory',
      'setBaseURL',
      'setNamespace',
      'onSubmit',
      'addValue',
      'removeValue',
      'updateValueKey',
      'updateValueOption',
    ]),
    handleClose() {
      this.$router.push('/')
    },
  },
})
</script>

<style lang="scss" scoped>
.value-list {
  list-style: none;

  .value {
    display: flex;
  }
}

.add-value-area {
  display: flex;
  justify-content: flex-end;
  margin: 12px 8px 12px auto;
}
</style>