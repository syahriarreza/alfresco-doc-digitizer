<template>
  <div>
    <v-row>
      <v-col
        v-for="(img, index) in images"
        :key="index"
        class="d-flex child-flex pa-0"
        cols="6"
        sm="4"
        md="3"
      >
        <base-gallery-thumb
          v-on:clickImage="expandImage"
          :src="img.src"
          :urlSrc="img.url_source"
          :contentType="img.content_asset_kind"
          :index="index"
        ></base-gallery-thumb>
      </v-col>
    </v-row>

    <v-dialog
      class="gallery--dialog"
      v-model="imgPopup.isShow"
      :overlay-opacity="0.9"
      :width="$vuetify.breakpoint.mdAndUp ? '50%' : '100%'"
    >
      <v-carousel v-model="imgPopup.selectedIndex" hide-delimiters @change="carouselOnChange">
        <v-carousel-item
          v-for="(item,i) in images"
          :key="i"
          :src="item.src"
          contain
          class="dialog--carousel--item"
        >
          <youtube
            v-if="isYoutube(item)"
            :video-id="getYoutubeID(item.url_source)"
            :fit-parent="true"
            :ref="`youtube-${i}`"
          ></youtube>
        </v-carousel-item>
      </v-carousel>
    </v-dialog>
  </div>
</template>

<script>
import { getIdFromUrl } from "vue-youtube";

export default {
  components: {
    BaseGalleryThumb: () => import("~shared/components/base/GalleryThumb"),
  },
  props: {
    images: {
      type: Array,
      default: () => [
        {
          src: "#",
        },
      ],
    },
  },
  data() {
    return {
      imgPopup: {
        isShow: false,
        selectedIndex: 0,
      },
    };
  },
  methods: {
    expandImage(imgIdx) {
      this.imgPopup.selectedIndex = imgIdx;
      this.imgPopup.isShow = true;
      let selImg = this.images[imgIdx];
      if (this.isYoutube(selImg)) {
        this.playYoutube(imgIdx);
      }
    },
    carouselOnChange(imgIdx) {
      this.playYoutube(imgIdx - 1, true);
      this.playYoutube(imgIdx + 1, true);
    },
    isYoutube(item) {
      return (
        (item.content_asset_kind.toLowerCase().indexOf("youtube") >= 0 &&
          item.url_source) ||
        item.url_source.toLowerCase().indexOf("youtube") >= 0
      );
    },
    getYoutubeID(url) {
      return getIdFromUrl(url);
    },
    playYoutube(imgIdx, isPause) {
      setTimeout(() => {
        let _vueCompArr = this.$refs[`youtube-${imgIdx}`];
        if (_vueCompArr && _vueCompArr.length > 0 && _vueCompArr[0].player) {
          let player = _vueCompArr[0].player;
          if (isPause) {
            player.pauseVideo();
          } else {
            player.playVideo();
          }
        }
      }, 200);
    },
  },
};
</script>

<style scoped>
div >>> .v-dialog {
  box-shadow: none !important;
}

.dialog--carousel--item >>> .v-responsive__content {
  margin-top: auto !important;
  margin-bottom: auto !important;
}
</style>