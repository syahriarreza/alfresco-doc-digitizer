<template>
  <div class="d-inline-flex">
    <v-col cols="12" sm="6" md="5" class="no-padding">
      <v-select
        v-model="item.text"
        :items="fieldDDL"
        item-text="text"
        item-value="value"
        label="Select"
        return-object
        single-line
        required
        @change="selectFiled(item)"
      ></v-select>
    </v-col>
    <v-col cols="12" sm="6" md="7" class="no-padding d-inline-flex">
      <v-select
        v-if="item.text.type == 'dropdown'"
        v-model="item.value"
        :items="item.data"
        item-text="text"
        item-value="value"
        label="Select"
        return-object
        single-line
        required
      ></v-select>
      <base-text-field v-else label="text" v-model="item.value"></base-text-field>

      <v-btn
        v-if="arrayField.length > 1"
        class="mx-2 btn-filed-filter"
        fab
        dark
        color="error"
        @click="removeFilter(item)"
      >
        <v-icon dark>mdi-minus</v-icon>
      </v-btn>
    </v-col>
  </div>
</template>

<script>
export default {
  props: {
    item: Object,
    fieldDDL: Array,
    arrayField: Array,
  },

  components: {
    BaseTextField: () => import('~shared/components/base/TextField'),
  },
  computed: {},
  methods: {
    selectFiled(item) {
      this.$emit('selectFiled', item)
    },
    removeFilter() {
      this.$emit('removeFilter')
    },
    addFilter() {},
  },
}
</script>
<style scoped>
.btn-filed-filter {
  height: 20px;
  width: 20px;
}

.btn-filed-filter .mdi-plus {
  font-size: 14px;
}

.btn-filed-filter .mdi-minus {
  font-size: 14px;
}

.v-select {
  padding-top: 0px;
  margin-top: 0px;
}
</style>