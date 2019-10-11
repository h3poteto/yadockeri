<template>
  <div>
    <el-dialog
      title="New Repository"
      :visible.sync="visible"
      width="40%"
      :before-close="handleClose"
      v-loading="loadingCreateProject"
    >
      <ValidationObserver ref="observer">
        <el-form label-width="150px" ref="form" slot-scope="{ validate }">
          <ValidationProvider rules="required" name="selectedRepository">
            <el-form-item
              label="Repository"
              slot-scope="{ errors }"
              :error="errors[0]"
            >
              <el-select
                v-model="selectedRepository"
                filterable
                placeholder="Select"
                :loading="loadingGithubRepo"
              >
                <el-option
                  v-for="item in githubRepos"
                  :key="item.html_url"
                  :label="item.full_name"
                  :value="item"
                ></el-option>
              </el-select>
            </el-form-item>
          </ValidationProvider>
          <ValidationProvider rules="required" name="selectedHelmRepository">
            <el-form-item
              label="Helm Repository"
              slot-scope="{ errors }"
              :error="errors[0]"
            >
              <el-select
                v-model="selectedHelmRepositoryUrl"
                filterable
                placeholder="Select"
                :loading="loadingGithubRepo"
              >
                <el-option
                  v-for="item in githubRepos"
                  :key="item.repository_url"
                  :label="item.full_name"
                  :value="item.html_url"
                ></el-option>
              </el-select>
            </el-form-item>
          </ValidationProvider>
          <el-form-item label="Helm Directory">
            <el-input
              :value="helmDirectory"
              placeholder="Helm Directory"
              @input="setHelmDirectory"
            ></el-input>
          </el-form-item>
          <ValidationProvider rules="required" name="baseURL">
            <el-form-item
              label="Base URL"
              slot-scope="{ errors }"
              :error="errors[0]"
            >
              <el-input
                v-model="baseURL"
                placeholder="https://dev.scout.lapras.com"
              ></el-input>
            </el-form-item>
          </ValidationProvider>
          <ValidationProvider rules="required" name="namespace">
            <el-form-item
              label="Namespace"
              slot-scope="{ errors }"
              :error="errors[0]"
            >
              <el-input v-model="namespace" placeholder="default"></el-input>
            </el-form-item>
          </ValidationProvider>
          <el-form-item label="Override Values">
            <ul class="value-list">
              <li class="value" v-for="(value, id) in values" :key="id">
                <el-input
                  :value="value.key"
                  @input="v => updateValueKey({ newKey: v, key: value.key })"
                  placeholder="image.tag"
                  size="small"
                ></el-input>
                <div>:</div>
                <el-input
                  :value="value.value"
                  @input="
                    v => updateValueOption({ newOption: v, key: value.key })
                  "
                  :placeholder="'{{.CommitSHA1}}'"
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
            <el-button type="primary" @click="validate().then(onSubmit)"
              >Create</el-button
            >
            <el-button @click="handleClose">Cancel</el-button>
          </el-form-item>
        </el-form>
      </ValidationObserver>
    </el-dialog>
  </div>
</template>

<script>
import Vue from 'vue'
import { createNamespacedHelpers } from 'vuex'
import { ValidationObserver, ValidationProvider } from 'vee-validate'

const { mapState, mapActions } = createNamespacedHelpers('pages/projects/new')

export default Vue.extend({
  name: 'new-project',
  components: {
    ValidationObserver,
    ValidationProvider,
  },
  data() {
    return {
      visible: true,
    }
  },
  computed: {
    ...mapState([
      'githubRepos',
      'helmDirectory',
      'loadingGithubRepo',
      'loadingCreateProject',
      'values',
    ]),
    selectedHelmRepositoryUrl: {
      get() {
        return this.$store.state.pages.projects.new.selectedHelmRepositoryUrl
      },
      set(value) {
        this.changeHelmRepository(value)
      },
    },
    selectedRepository: {
      get() {
        return this.$store.state.pages.projects.new.selectedRepository
      },
      set(value) {
        this.changeRepository(value)
      },
    },
    baseURL: {
      get() {
        return this.$store.state.pages.projects.new.baseURL
      },
      set(value) {
        this.setBaseURL(value)
      },
    },
    namespace: {
      get() {
        return this.$store.state.pages.projects.new.namespace
      },
      set(value) {
        this.setNamespace(value)
      },
    },
  },
  async created() {
    await this.fetchGithubRepos().catch(err => {
      console.error(err)
      this.$message({
        message: 'Failed to get github repositories',
        type: 'error',
      })
    })
  },
  methods: {
    ...mapActions([
      'fetchGithubRepos',
      'changeRepository',
      'changeHelmRepository',
      'setHelmDirectory',
      'setBaseURL',
      'setNamespace',
      'submit',
      'addValue',
      'removeValue',
      'updateValueKey',
      'updateValueOption',
    ]),
    handleClose() {
      this.$router.push('/')
    },
    async onSubmit() {
      this.values.map(v => {
        if (v.key.length === 0 || v.value.length === 0) {
          this.$message({
            message: 'Key and value are not allowed empty in Override Values',
            type: 'error',
          })
          throw new Error('key is empty')
        }
      })
      await this.submit().catch(err => {
        console.error(err)
        this.$message({
          message: 'Failed to create a new project',
          type: 'error',
        })
      })
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
