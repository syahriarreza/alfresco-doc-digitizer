<template>
  <base-card color="transparent">
    <base-img
      :src="item.src"
      cover
      aspect-ratio="1.7"
      color="grey lighten-1"
      class="rounded"
      tile
      @click="detailNews(item)"
    >
      <template v-slot:placeholder>
        <v-row class="fill-height ma-0" align="center" justify="center">
          <v-progress-circular indeterminate color="grey lighten-5"></v-progress-circular>
        </v-row>
      </template>
    </base-img>
    <base-subtitle :title="item.date" class="mb-4 mt-4" />
    <base-divider color="primary" />
    <!-- TODO : define card spacing text -->
    <h3 class="mt-n2">{{item.title}}</h3>
  </base-card>
</template>

<script>
export default {
  components: {
    BaseTitle: () => import("./Title"),
    BaseSubtitle: () => import("./Subtitle"),
    BaseDivider: () => import("./Divider"),
    BaseCard: () => import("../base/Card"),
    BaseImg: () => import("../base/Img"),
    BaseBtn: () => import("../base/Btn"),
  },
  props: {
    title: {
      type: String,
      default: "No Title",
    },
    text: {
      type: String,
      default: "No Content",
    },
    date: {
      type: [String, Date],
    },
    src: {
      type: String,
      default: "",
    },
    url: {
      type: String,
      default: "#",
    },
    item: { type: Object, default: null },
  },
  methods: {
    detailNews(val) {
      this.$emit("latestNews", val);
    },
  },
};
</script>

<style scoped>
a.news-title {
  font-size: 1.615rem;
  font-weight: 700;
  color: black;
  text-decoration: none;
  font-weight: bold;
}
</style>

<style>
.title-no-spacing {
  letter-spacing: 0 !important;
}
</style>