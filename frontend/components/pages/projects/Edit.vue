<template>
  <div>
    <el-dialog
      title="Edit Repository"
      :visible.sync="visible"
      width="40%"
      :before-close="handleClose"
    >
      <ValidationObserver ref="observer">
        <el-form label-width="150px" ref="form" slot-scope="{ validate }">
          <template v-if="project">
            <el-form-item label="Repository">
              <el-input :value="project.repository_name" disabled></el-input>
            </el-form-item>
            <el-form-item label="Helm Repository">
              <el-input
                :value="project.helm_repository_url"
                disabled
              ></el-input>
            </el-form-item>
            <el-form-item label="Helm Directory">
              <el-input
                v-model="helmDirectory"
                placeholder="Helm Directory"
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
              <el-button type="primary" @click="validate().then(onSubmit)"
                >Update</el-button
              >
              <el-button @click="handleClose">Cancel</el-button>
            </el-form-item>
          </template>
        </el-form>
      </ValidationObserver>
    </el-dialog>
  </div>
</template>

<script>
import Vue from 'vue'
import { createNamespacedHelpers } from 'vuex'
import { ValidationObserver, ValidationProvider } from 'vee-validate'

const { mapState, mapActions } = createNamespacedHelpers('pages/projects/edit')

export default Vue.extend({
  name: 'edit-project',
  components: {
    ValidationObserver,
    ValidationProvider,
  },
  props: ['project_id'],
  data() {
    return {
      visible: true,
    }
  },
  computed: {
    ...mapState(['project', 'values']),
    helmDirectory: {
      get() {
        return this.$store.state.pages.projects.edit.helmDirectory
      },
      set(value) {
        this.changeHelmDirectory(value)
      },
    },
    baseURL: {
      get() {
        return this.$store.state.pages.projects.edit.baseURL
      },
      set(value) {
        this.changeBaseURL(value)
      },
    },
    namespace: {
      get() {
        return this.$store.state.pages.projects.edit.namespace
      },
      set(value) {
        this.changeNamespace(value)
      },
    },
  },
  async created() {
    await this.fetchProject(this.project_id)
  },
  methods: {
    ...mapActions([
      'fetchProject',
      'changeHelmDirectory',
      'changeBaseURL',
      'changeNamespace',
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
      await this.submit(this.project_id).catch(err => {
        console.error(err)
        this,
          $message({
            message: 'Failed to update this project',
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
