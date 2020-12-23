<template>
  <div :class="classes" class="pt-2">
    <!-- <div :class="horizontal && title && 'ml-6'"> -->
    <div>
      <base-title :title="title" class="text-uppercase" :space="spaceTitle" />

      <base-body
        v-if="text || $slots.default"
        :space="horizontal ? 0 : undefined"
        :text="text"
        class="mx-auto"
        max-width="700"
      >
        <slot></slot>
      </base-body>
    </div>
  </div>
</template>

<script>
// Mixins
import Heading from "@/mixins/heading";

export default {
  name: "BaseInfoDetail",

  components: {
    BaseTitle: () => import("./Title"),
    BaseBody: () => import("./Body")
  },

  mixins: [Heading],

  props: {
    align: {
      type: String,
      default: "left"
    },
    color: String,
    dark: Boolean,
    horizontal: Boolean,
    icon: String,
    outlined: {
      type: Boolean,
      default: true
    },
    space: {
      type: [Number, String],
      default: 8
    },
    spaceTitle: {
      type: [Number, String],
      default: 3
    },
    size: {
      type: [Number, String],
      default: 72
    },
    text: String,
    title: String
  },

  computed: {
    classes() {
      const classes = [`mb-${this.space}`];

      if (this.horizontal) {
        classes.push("d-flex");

        if (!this.$slots.default && !this.text) {
          classes.push("align-center");
        }
      }

      return classes;
    }
  }
};
</script>
