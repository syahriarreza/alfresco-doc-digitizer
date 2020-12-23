<template>
  <v-data-table v-bind="$attrs" v-on="$listeners">
    <!-- pass through scoped slots -->
    <template v-for="(_, scopedSlotName) in $scopedSlots" v-slot:[scopedSlotName]="slotData">
      <slot :name="scopedSlotName" v-bind="slotData" />
    </template>
    <!-- pass through normal slots -->
    <template v-for="(_, slotName) in $slots" v-slot:[slotName]>
      <slot :name="slotName" />
    </template>
  </v-data-table>
</template>
<script>
export default {
  name: "BaseDataTable",
  props: {},
};
</script>

<style lang="sass">
  //
</style>
