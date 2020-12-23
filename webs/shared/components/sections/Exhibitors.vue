<template>
  <base-section id="exhibitor" class="grey lighten-5">
    <base-section-heading title="GIIAS 2020 EXHIBITORS"></base-section-heading>

    <v-container>
      <v-row>
        <v-col v-for="(feature, i) in exhibitors" :key="i" cols="6" md="3" class="box-img">
          <base-box-display-item v-bind="feature" :item="feature" :height="150" align="center"></base-box-display-item>
        </v-col>
      </v-row>
    </v-container>
  </base-section>
</template>

<script>
import { cloneDeep } from "lodash";
export default {
  name: "SectionExhibitors",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    BaseBoxDisplayItem: () => import("../base/BoxDisplayItem"),
  },

  data: () => ({
    features: [
      {
        title: "",
        src: require("~/static/carlogo/logo-car-1.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-2.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-3.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-4.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-5.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-6.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-7.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-8.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-9.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-10.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-11.png"),
        link: "",
        description: "",
        showBorder: true,
      },
      {
        title: "",
        src: require("~/static/carlogo/logo-car-12.png"),
        link: "",
        description: "",
        showBorder: true,
      },
    ],
  }),
  computed: {
    exhibitors() {
      let datas = cloneDeep(this.$store.state.tenant.tenants);
      return datas.map((_) => {
        _.src = this.$tool.getAssetURL(_.brand_image_asset_id);
        _.showBorder = true;
        _.description = "";
        return _;
      });
    },
  },
};
</script>
<style scoped>
.box-img {
  height: 100% !important;
}

section >>> .gt-logo .v-image__image--cover {
  background-size: auto 90px !important;
}
</style>
