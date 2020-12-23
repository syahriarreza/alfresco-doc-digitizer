<template>
  <div class="pt-2 gt-col">
    <div class="gt-inner">
      <div class="gt-logo">
        <base-img
          cover
          aspect-ratio="1.7"
          :src="item.src"
          color="grey lighten-1"
          :width="width"
          :height="height"
          :class="{ full: fullWidthImage, 'mb-2': marginBot }"
          :contain="contain"
          tile
          @click="clickImage"
        />
      </div>
      <v-col
        class="box-description"
        :justify="justify"
        v-if="showText && (item.title || item.description)"
        :cols="12"
      >
        <base-title
          v-if="item.title"
          align="align"
          :title="item.title"
          class="text-uppercase"
          space="1"
          @click="clickTitle"
        />
        <base-divider v-if="item.text" class="mr-auto" :color="color" space="1" />
        <base-body v-if="item.text" :text="text" space="1">
          <slot></slot>
        </base-body>
      </v-col>

      <v-col v-if="showText && item.description" cols="12">
        <base-subtitle align="align" :title="item.description" space="1" />
      </v-col>
      <div :class="item.showBorder == true ? 'line primary':'line'"></div>
    </div>
  </div>
</template>

<script>
// Mixins
import Heading from "@/mixins/heading";

export default {
  name: "BaseBannerBox",

  components: {
    BaseImg: () => import("./Img"),
    BaseTitle: () => import("./Title"),
    BaseDivider: () => import("./Divider"),
    BaseSubtitle: () => import("./Subtitle"),
    BaseBody: () => import("./Body"),
  },

  mixins: [Heading],

  props: {
    item: Object,
    customClass: String,
    width: { type: Number, default: 100 },
    height: { type: Number, default: 100 },
    contain: { type: Boolean, default: false },
    marginBot: { type: Boolean, default: false },
    showText: { type: Boolean, default: true },
    align: {
      type: String,
      default: "left",
    },
    color: {
      type: String,
      default: "primary",
    },
  },
  data: () => ({
    fullWidthImage: true,
  }),
  computed: {
    classes() {
      const classes = [`mb-${this.space}`];
      classes.push("text-center");
      return classes;
    },
  },
  methods: {
    clickImage() {
      this.$emit("clickImage", this.item);
    },
    clickTitle() {
      this.$emit("clickTitle", this.item);
    },
  },
};
</script>

<style scoped>
.gt-col {
  width: 100%;
  position: relative;
  display: inline-block;
}
.gt-col > .gt-inner {
  display: inline-block;
  height: 100%;
  width: 100%;
}
.gt-logo {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-wrap: wrap;
  text-align: center;
  /* padding: 10px; */
  box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
  -webkit-box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
  -moz-box-shadow: 0px -4px 14px 0px rgba(215, 215, 215, 0.349019607843137);
  height: 100%;
  width: 100%;
  border-radius: 5px 5px 0px 0px;
  overflow: hidden;
  background-color: #f9f9f9;
}
.line {
  width: 100%;
  height: 5px;
  border-radius: 0px 0px 5px 5px;
}
</style>
<style>
</style>