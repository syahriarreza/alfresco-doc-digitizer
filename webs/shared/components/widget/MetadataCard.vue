<template>
  <v-card>
    <v-card-title>Event Details</v-card-title>
    <v-divider></v-divider>
    <v-card-text>
      <base-icon-card
        v-for="(m, i) in metas"
        :key="i"
        :icon="m.icon"
        :outlined="false"
        :title="m.title"
        :color="color"
        size="subtitle"
        horizontal
        space="0"
        spaceTitle="0"
      >{{m.text}}</base-icon-card>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  components: {
    BaseIconCard: () => import("../base/IconCard")
  },
  props: {
    color: {
      type: String,
      default: "primary"
    },
    metas: {
      type: Array,
      default: () => [
        {
          title: "No Title",
          text: "No Content",
          icon: "mdi-border-none-variant"
        }
      ]
    }
  }
};
</script>

<style lang="scss" scoped>
</style>