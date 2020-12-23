<template>
  <base-section id="theme-features" class="grey lighten-5">
    <base-section-heading title="Floor Plan"></base-section-heading>

    <v-responsive class="mx-auto" max-width="1350">
      <v-container fluid>
        <v-col
          v-for="(floorPlan, i) in floorPlans"
          :key="i"
          cols="12"
          sm="12"
          md="12"
          class="box-img"
        >
          <v-card>
            <v-img :src="floorPlan.src" :key="i"></v-img>
          </v-card>
          <!-- <base-box-display-item v-bind="feature" :item="feature" :height="400" align="center"></base-box-display-item> -->
        </v-col>
      </v-container>
    </v-responsive>
  </base-section>
</template>

<script>
import { cloneDeep } from "lodash";
export default {
  name: "SectionFloorPlan",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    BaseBoxDisplayItem: () => import("../base/BoxDisplayItem"),
  },

  data: () => ({
    features: [
      {
        title: "",
        src: require("~/static/floor-plan.jpg"),
        link: "",
        description: "",
        showBorder: false,
      },
    ],
  }),
  computed: {
    floorPlans() {
      let datas = cloneDeep(this.$store.state.event.floorPlan);
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
