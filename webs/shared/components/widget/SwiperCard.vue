<template>
  <swiper class="swiper pa-2" :options="swiperOption">
    <swiper-slide v-for="(item, i) in items" :key="i" class="mr-5">
      <v-card
        :elevation="i == index ? 0 : 5"
        max-width="200"
        @click="clickCard(item, i)"
        :color="i == index ? 'primary-gradient' : ''"
        :class="{ 'white--text': i == index }"
      >
        <v-card-title class="pa-2">
          <span class="mx-auto">{{ item[titleProp] }}</span>
        </v-card-title>
        <v-card-text class="px-0 pt-0">
          <v-img
            :src="$tool.getAssetURL(item[assetIDProp])"
            contain
            max-height="100"
            max-width="200"
          >
            <template v-slot:placeholder>
              <v-row class="fill-height ma-0" align="center" justify="center">
                <v-progress-circular indeterminate color="grey lighten-5"></v-progress-circular>
              </v-row>
            </template>
          </v-img>
        </v-card-text>
      </v-card>
    </swiper-slide>
    <div class="swiper-pagination" slot="pagination"></div>
    <!-- <div class="swiper-button-prev" slot="button-prev"></div>
    <div class="swiper-button-next" slot="button-next"></div>-->
  </swiper>
</template>

<script>
import { Swiper, SwiperSlide } from "vue-awesome-swiper";
import "~shared/assets/swiper/swiper-base.min.css";
import "~shared/assets/swiper/swiper-v5.3.6.min.css";
import "~shared/assets/swiper/swiper-v5.3.6.min.js";

export default {
  components: {
    Swiper,
    SwiperSlide,
  },
  props: {
    swiperOption: {
      type: Object,
      default: () => ({
        slidesPerView: 3,
        spaceBetween: 30,
        freeMode: true,
        pagination: {
          el: ".swiper-pagination",
          clickable: true,
        },
        // navigation: {
        //   nextEl: ".swiper-button-next",
        //   prevEl: ".swiper-button-prev",
        // },
      }),
    },
    items: {
      type: Array,
      default: () => [
        {
          title: "No Title",
          asset_id: "",
        },
      ],
    },
    titleProp: {
      type: String,
      default: "title",
    },
    assetIDProp: {
      type: String,
      default: "asset_id",
    },
    index: {
      type: Number,
      default: -1,
    },
  },
  methods: {
    clickCard(item, index) {
      this.$emit("click", { item, index });
    },
  },
};
</script>

<style scoped>
.swiper-slide {
  max-width: 210px !important;
}
.v-card {
  cursor: grab;
}
</style>