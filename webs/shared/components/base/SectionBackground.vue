<template>
  <section v-bind="$attrs" :style="styles" v-on="$listeners">
    <base-img :height="height" :src="src" :color="color" flat max-width="100%" eager tile>
      <slot name="append-outer" />
    </base-img>
  </section>
</template>

<script>
// Components
import Measurable from "vuetify/lib/mixins/measurable";

export default {
  name: "BaseSection",

  mixins: [Measurable],
  components: {
    BaseImg: () => import("../base/Img")
  },

  props: {
    src: { type: String, default: "" },
    color: { type: String, default: "#45516b" },
    height: { type: Number, default: 550 },
    space: {
      type: [Number, String],
      default: 0
    }
  },

  computed: {
    styles() {
      const space = this.$vuetify.breakpoint.mdAndUp
        ? this.space
        : this.space / 2;

      return {
        ...this.measurableStyles,
        padding: `${space}px 0`
      };
    }
  }
};
</script>
