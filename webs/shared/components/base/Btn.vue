<template>
  <v-btn
    :color="color"
    :depressed="depressed"
    :min-width="minWidth"
    :tile="tile"
    :class="btnClass"
    :x-large="xlarge"
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
  </v-btn>
</template>

<script>
export default {
  name: "BaseBtn",

  props: {
    color: {
      type: String,
      default: "primary",
    },
    depressed: {
      type: Boolean,
      default: true,
    },
    minWidth: {
      type: [Number, String],
      default: 164,
    },
    tile: {
      type: Boolean,
      default: true,
    },
    xlarge: { type: Boolean, default: true },
    btnClass: { type: String, default: "font-weight-bold" },
  },
};
</script>
