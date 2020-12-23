<template>
  <v-theme-provider dark>
    <section id="hero-alt">
      <base-img
        :height="height ? height : $vuetify.breakpoint.mdAndUp ? 350 : 225"
        :gradient="gradient"
        :src="src"
        color="#45516b"
        flat
        max-width="100%"
        tile
        :contain="contain"
      >
        <v-row v-if="title" align="center" class="ma-0 fill-height" :class="'text-'+align">
          <v-col :cols="defCols" :offset="defOffset">
            <base-heading :title="title" :align="align" space="2" weight="500" />

            <base-divider color="primary" :align="align" dense />
            <v-breadcrumbs :items="items" class="pa-0" :class="defBreadcrumbClass" divider=">" />
            <slot></slot>
          </v-col>
        </v-row>
      </base-img>
    </section>
  </v-theme-provider>
</template>

<script>
// Components
import { HexToRGBA, RGBAtoCSS } from "vuetify/lib/util/colorUtils";

export default {
  name: "SectionHeroPage",

  components: {
    BaseImg: () => import("../base/Img"),
    BaseHeading: () => import("../base/Heading"),
    BaseDivider: () => import("../base/Divider")
  },

  provide: {
    heading: { align: "center" }
  },

  data: () => ({}),

  computed: {
    gradient() {
      const color = `${this.$vuetify.theme.themes.light.secondary}CC`;
      const overlay = RGBAtoCSS(HexToRGBA(color));

      return `to top, ${overlay}, ${overlay}`;
    },
    defCols() {
      return this.align == "center" || this.$vuetify.breakpoint.smAndDown
        ? 12
        : 5;
    },
    defOffset() {
      switch (this.align) {
        case "left":
          return 1;
        case "right":
          return 6;
        default:
          return 0;
      }
    },
    defBreadcrumbClass() {
      switch (this.align) {
        case "center":
          return "justify-center";
        case "right":
          return "float-right";
        default:
          return "";
      }
    }
  },

  props: {
    items: { type: Array, default: () => [{ text: "HOME", to: "/" }] },
    src: { type: String, default: "" },
    title: { type: String, default: "" },
    contain: { type: Boolean, default: false },
    height: { type: [Number, String] },
    align: { type: String, default: "center" }
  }
};
</script>

<style lang="sass">
#hero-alt
  .v-breadcrumbs__item
    color: #FFFFFF
</style>
