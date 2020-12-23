<template>
  <v-textarea
    :outlined="outlined"
    :dense="dense"
    :single-line="singleLine"
    :hide-details="hideDetails"
    v-bind="{
      rows: 6,
      ...$attrs,
    }"
    v-on="$listeners"
  />
</template>

<script>
export default {
  name: 'BaseTextarea',
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
