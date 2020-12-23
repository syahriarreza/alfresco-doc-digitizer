<template>
  <!-- grid toolbar -->
  <v-toolbar flat color="transparent" class="v-toolbar--grid mt-n5 mb-n2">
    <!-- grid filter -->
    <v-menu
      v-if="showFilter"
      offset-y
      :close-on-content-click="false"
      :max-width="maxWidth"
      :min-width="maxWidth"
    >
      <template v-slot:activator="{ on }">
        <!-- filter button -->
        <v-btn class="mr-2" large v-on="on" :class="{'primary--text':false}">
          <v-icon>mdi-tune</v-icon>
        </v-btn>
      </template>
      <v-card>
        <v-form ref="form">
          <v-card-title>Filter</v-card-title>
          <v-divider />
          <v-card-text>
            <slot name="filter-form" />
          </v-card-text>
          <v-divider />
          <v-card-actions>
            <v-btn text @click="$emit('on-clear')">Ulang</v-btn>
            <v-spacer />
            <v-btn color="primary" @click="$emit('on-filter')" text>Cari</v-btn>
          </v-card-actions>
        </v-form>
      </v-card>
    </v-menu>

    <!-- grid search -->
    <v-text-field
      append-icon="mdi-magnify"
      label="Cari ..."
      hide-details
      single-line
      solo
      clearable
      @keyup.enter="$emit('on-search')"
      @click:append="$emit('on-search')"
      @click:clear="onClear"
      v-model="inputValue"
    ></v-text-field>
    <v-spacer />

    <!-- grid button -->
    <slot name="button" />
  </v-toolbar>
</template>

<script>
export default {
  props: {
    value: { type: String, default: '' },
    maxWidth: { type: String, default: '300px' },
    showFilter: { type: Boolean, default: true },
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
    hasDefaultSlot() {
      return !!this.$slots.default
    },
    hasFilterFormSlot() {
      return !!this.$slots['filter-form']
    },
  },
  methods: {
    onClear() {
      this.inputValue = ''
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
}
</style>