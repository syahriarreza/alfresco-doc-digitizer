<template>
  <base-section id="theme-features">
    <base-section-heading title="SPONSORS"></base-section-heading>

    <v-container>
      <v-row>
        <v-col v-for="(feature, i) in sponsores" :key="i" cols="12" sm="6" md="4" class="box-img">
          <base-box-display-item
            v-bind="feature"
            :item="feature"
            :width="300"
            :height="200"
            align="center"
          ></base-box-display-item>
        </v-col>
      </v-row>
    </v-container>
  </base-section>
</template>

<script>
import { cloneDeep } from "lodash";
export default {
  name: "SectionSponsors",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    BaseBoxDisplayItem: () => import("../base/BoxDisplayItem"),
  },

  data: () => ({
    features: [
      {
        title: "",
        src: require("~/static/sponsor/logo-spongsor-1.jpg"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/sponsor/logo-spongsor-2.jpg"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/sponsor/logo-spongsor-3.jpg"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/sponsor/logo-spongsor-4.jpg"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/sponsor/logo-spongsor-5.jpg"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/sponsor/logo-spongsor-6.jpg"),
        link: "",
        description: "",
        showBorder: true,
      },
    ],
  }),
  computed: {
    sponsores() {
      let datas = cloneDeep(this.$store.state.event.sponsors);
      return datas.map((_) => {
        _.src = this.$tool.getAssetURL(_.asset_id);
        _.description = "";
        _.showBorder = false;
        return _;
      });
      return _;
    },
  },
};
</script>
<style scoped>
.box-img {
  height: 100% !important;
}
section >>> .gt-logo .v-image__image--cover {
  background-size: auto 90% !important;
  background-color: white;
}
</style>
