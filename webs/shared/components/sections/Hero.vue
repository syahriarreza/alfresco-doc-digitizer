<template>
  <section id="hero" class="hero" :class="{'hero__border-wave':border=='wave'}">
    <v-img
      :height="minHeight"
      :src="src"
      class="white--text"
      :gradient="overlay?overlayGradient:''"
    >
      <v-container class="fill-height px-4 py-12" v-if="overlay">
        <slot></slot>
      </v-container>
      <img class="border-wave" src="~static/border-wave.svg" />
      <!-- <v-row  align="center" class="ma-0 fill-height border-ombak" justify="center">
          
      </v-row>-->
    </v-img>
  </section>
</template>

<script>
export default {
  name: "SectionHero",

  components: {
    BaseBody: () => import("../base/Body"),
    BaseHeading: () => import("../base/Heading"),
    BaseBtn: () => import("../base/Btn")
  },

  provide: {
    theme: { isDark: true }
  },

  computed: {
    minHeight() {
      // const height = this.$vuetify.breakpoint.mdAndUp ? "100vh" : "45vh";

      // return `calc(${height} - ${this.$vuetify.application.top}px)`;
      return `100vh`;
    }
  },

  props: {
    src: {
      type: String,
      default: ""
    },
    overlay: {
      type: Boolean,
      default: false
    },
    overlayGradient: {
      type: String,
      default: "to right, rgba(5, 11, 31, .8), rgba(5, 11, 31, .8)"
    },
    border: {
      type: String,
      default: "none"
    }
  },

  data: () => ({})
};
</script>
<style>
/* .hero.hero__border-wave::after {
  background: url("~static/border-wave.svg");
  background-repeat: no-repeat;
  background-size: 100%;
  display: block;
  width: 100%;
  margin-top: -140px;
  position: relative;
  height: 150px;
  content: " ";
} */
.border-wave {
  position: absolute;
  bottom: -1px;
  width: 100%;
}
</style>