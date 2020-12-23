<template>
  <!-- grid toolbar -->
  <v-toolbar flat color="transparent" class="v-toolbar--grid mt-n5 mb-n2">
    <!-- grid filter -->
    <v-select
      v-if="!hide_search"
      class="btn-field"
      :items="field"
      v-model="inputValue.selected"
      item-text="text"
      item-value="value"
      label="Select Field"
      :solo="true"
    ></v-select>
    <!-- grid search -->
    <v-text-field
      v-if="!hide_search"
      append-icon="mdi-magnify"
      label="Cari ..."
      hide-details
      single-line
      solo
      clearable
      @keyup.enter="$emit('on-search')"
      @click:append="$emit('on-search')"
      @click:clear="onClear"
      v-model="inputValue.search"
    ></v-text-field>
    <v-spacer />
    <!-- grid button -->
    <slot name="button" />
  </v-toolbar>
</template>

<script>
export default {
  props: {
    value: {
      selected: '',
      search: '',
    },
    field: Array,
    hide_search: {
      type: Boolean,
      default: false,
    },
  },
  computed: {
    inputValue: {
      get() {
        return this.value
      },
      set(nv) {
        this.$emit('input', nv)
      },
    },
  },
  methods: {
    onClear() {
      this.inputValue = {
        selected: 'all',
        search: '',
      }
      this.$emit('on-search')
    },
  },
}
</script>

<style lang="scss">
.v-toolbar--grid {
  .v-toolbar__content {
    padding-left: 0px !important;
    padding-right: 0px !important;
  }
  .btn-field {
    height: 48px;
    flex: none;
    width: 150px;
    margin-right: 10px;
  }
}
</style>