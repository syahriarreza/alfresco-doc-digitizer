<template>
  <v-card>
    <v-card-title>Sponsors</v-card-title>
    <v-divider></v-divider>
    <widget-sponsor-grid :items="items" :cols="6"></widget-sponsor-grid>
  </v-card>
</template>

<script>
export default {
  components: {
    WidgetSponsorGrid: () => import("./SponsorGrid")
  },
  props: {
    items: {
      type: Array,
      default: () => [
        {
          title: "",
          src: require("~/static/sponsor/logo-spongsor-1.jpg"),
          link: "",
          description: "",
          showBorder: false
        }
      ]
    }
  }
};
</script>

<style lang="scss" scoped>
</style>