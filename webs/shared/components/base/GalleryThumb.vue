<template>
  <v-card @click="clickImage">
    <v-img
      :src="isYoutube ? ytThumbnail : src"
      aspect-ratio="1"
      class="grey lighten-2 img-pointer"
      @click="clickImage"
    >
      <template v-slot:placeholder>
        <v-row class="fill-height ma-0" align="center" justify="center">
          <v-progress-circular indeterminate color="grey lighten-5"></v-progress-circular>
        </v-row>
      </template>
    </v-img>
    <v-overlay v-if="isYoutube" :absolute="true" :value="true" :opacity="0">
      <v-btn icon x-large>
        <v-img :src="ytPlayBtnSrc" aspect-ratio="1" class="img-pointer"></v-img>
      </v-btn>
    </v-overlay>
  </v-card>
</template>

<script>
import { getIdFromUrl } from "vue-youtube";

export default {
  data() {
    return {
      ytPlayBtnSrc: require("~/static/yt-play-button.png"),
    };
  },
  props: {
    src: {
      type: String,
      default: "", //TODO: add Broken Image src
    },
    urlSrc: {
      type: String,
      default: "",
    },
    index: {
      type: Number,
      default: -1,
    },
    contentType: {
      type: String,
      default: "file",
    },
  },
  methods: {
    clickImage() {
      this.$emit("clickImage", this.index);
    },
  },
  computed: {
    isYoutube() {
      return (
        (this.contentType.toLowerCase().indexOf("youtube") >= 0 &&
          this.urlSrc) ||
        this.urlSrc.toLowerCase().indexOf("youtube") >= 0
      );
    },
    ytThumbnail() {
      let videoID = getIdFromUrl(this.urlSrc);
      return `https://img.youtube.com/vi/${videoID}/sddefault.jpg`;
    },
  },
};
</script>

<style scoped>
.img-pointer {
  cursor: pointer !important;
}
</style>