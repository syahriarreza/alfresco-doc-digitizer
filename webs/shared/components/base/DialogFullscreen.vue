<template>
  <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
    <section-hero-page :src="src" :title="title" :items="breadcrumbs"></section-hero-page>
    <v-card>
      <v-toolbar dark color="primary">
        <v-spacer></v-spacer>
        <v-toolbar-items>
          <v-btn icon dark @click="hideDialogBox">
            <v-icon>mdi-close</v-icon>
          </v-btn>
        </v-toolbar-items>
      </v-toolbar>
      <slot name="append-outer" />
    </v-card>
  </v-dialog>
</template>

<script>
export default {
  props: {
    dialog: { type: Boolean, default: false },
    title: { type: String, default: "Detail Product" },
    breadcrumbs: {
      type: Array,
      default: [
        { text: "Product", to: "/" },
        { text: "Detail Product", to: "/" },
      ],
    },
  },
  components: {
    SectionHeroPage: () => import("~shared/components/sections/HeroPage"),
  },
  data() {
    return {
      src: require("~/static/about.jpg"),
      notifications: false,
      sound: true,
      widgets: false,
    };
  },
  methods: {
    hideDialogBox() {
      this.$emit("hideDialog", false);
    },
  },
};
</script>
<style scoped>
.v-toolbar {
  height: 64px;
  position: fixed;
  width: 100%;
  z-index: 1;
  top: 0;
  background: transparent;
  opacity: 0.5;
}
</style>