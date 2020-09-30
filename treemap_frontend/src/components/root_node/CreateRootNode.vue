<template>
  <div>
    <b-button v-b-modal.modal-add-skill_tree><i class="fa fa-plus"></i></b-button>
    <b-modal id="modal-add-skill_tree" title="Create SkillTree" @ok="saveResource">
      <b-form>
        <b-form-group label="Name">
          <b-form-input type="text" v-model="model.name"></b-form-input>
        </b-form-group>
        <b-form-group label="Description">
          <b-form-input type="text" v-model="model.description"></b-form-input>
        </b-form-group>
      </b-form>
    </b-modal>
  </div>
</template>

<script>
import lightweightRestful from 'vue-lightweight_restful'
import consts from '@/utils/consts'

export default {
  name: 'CreateRootNode',
  data() {
    return {
      model: {
        'name': '',
        'node_type': consts.model.node.node_type.root,
        'description': '',
      }
    }
  },
  methods: {
    async saveResource() {
      let response = await lightweightRestful.api.createResource(consts.api.v1.node.root, this.model)
      this.$emit('create_new_root_node', response.id)
    }
  }
}
</script>

<style scoped>

</style>
