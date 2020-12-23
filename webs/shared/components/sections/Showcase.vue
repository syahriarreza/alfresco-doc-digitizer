<template>
  <div>
    <v-row class="pt-4 pb-11" style="background-color: white;">
      <v-col>
        <base-section-heading title="SHOWCASE"></base-section-heading>
        <v-carousel hide-delimiters height="100%" width="100%">
          <v-carousel-item :key="i" v-for="(gs, i) in groupedShowcase">
            <v-layout row>
              <v-flex :key="j" v-for="(s, j) in gs">
                <v-img class="img-pointer" :src="s.src" aspect-ratio="1.7" @click="expandImage(s)"></v-img>
              </v-flex>
            </v-layout>
          </v-carousel-item>
        </v-carousel>
      </v-col>
    </v-row>

    <v-dialog
      class="gallery--dialog"
      v-model="imgPopup.isShow"
      :overlay-opacity="0.9"
      :width="$vuetify.breakpoint.mdAndUp ? '50%' : '100%'"
    >
      <v-carousel v-model="imgPopup.selectedIndex" hide-delimiters>
        <v-carousel-item v-for="(item,i) in showcases" :key="i" :src="item.src" contain></v-carousel-item>
      </v-carousel>
    </v-dialog>
  </div>
</template>

<script>
import { chunk, findIndex } from "lodash";

export default {
  components: {
    BaseSectionHeading: () => import("~shared/components/base/SectionHeading"),
  },
  props: {
    showcasePerPage: {
      type: Number,
      default: 5,
    },
    showcases: {
      type: Array,
      default: () => [{ _id: 0, src: require("~/static/u856.jpg") }],
    },
  },
  computed: {
    groupedShowcase() {
      return chunk(
        this.showcases,
        this.$vuetify.breakpoint.smAndDown ? 1 : this.showcasePerPage
      );
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
    expandImage(sc) {
      console.log("showcase sel:", sc);
      let selIdx = findIndex(this.showcases, (_) => _._id == sc._id);
      console.log("selIdx:", selIdx);
      this.imgPopup.selectedIndex = selIdx;
      this.imgPopup.isShow = true;
    },
  },
};
</script>

<style scoped>
div >>> .v-dialog {
  box-shadow: none !important;
}
.img-pointer {
  cursor: pointer !important;
}
</style>