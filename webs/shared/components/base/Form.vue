<template>
  <ValidationObserver ref="formValidator">
    <v-form ref v-bind="$attrs" v-on="$listeners">
      <!-- pass through scoped slots -->
      <template v-for="(_, scopedSlotName) in $scopedSlots" v-slot:[scopedSlotName]="slotData">
        <slot :name="scopedSlotName" v-bind="slotData" />
      </template>
      <!-- pass through normal slots -->
      <template v-for="(_, slotName) in $slots" v-slot:[slotName]>
        <slot :name="slotName" />
      </template>
    </v-form>
  </ValidationObserver>
</template>
<script>
export default {
  name: "BaseForm",
  inheritAttrs: false,
  props: {},
  methods: {
    validate() {
      this.$refs.formValidator.validate();
    },
    validate() {
      this.$refs.formValidator.reset();
    },
    resetValidation() {
      this.$refs.formValidator.resetValidation();
    },
  },
};
</script>

<style lang="sass">
  //
</style>
