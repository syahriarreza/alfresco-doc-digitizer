<template>
  <v-text-field
    :outlined="outlined"
    :dense="dense"
    :single-line="singleLine"
    :hide-details="hideDetails"
    v-bind="$attrs"
    v-on="$listeners"
  >
    <!-- pass through scoped slots -->
    <template v-for="(_, scopedSlotName) in $scopedSlots" v-slot:[scopedSlotName]="slotData">
      <slot :name="scopedSlotName" v-bind="slotData" />
    </template>
    <!-- pass through normal slots -->
    <template v-for="(_, slotName) in $slots" v-slot:[slotName]>
      <slot :name="slotName" />
    </template>
  </v-text-field>
</template>
<script>
export default {
  name: 'BaseTextField',
  props: {
    outlined: { type: Boolean, default: true },
    dense: { type: Boolean, default: true },
    singleLine: { type: Boolean, default: false },
    hideDetails: { type: [String, Boolean], default: 'auto' },
  },
}
</script>

<style lang="sass">
  //
</style>
