<template>
  <base-section id="theme-features">
    <base-section-heading :title="title" v-if="title"></base-section-heading>
    <v-row>
      <v-col v-for="(feature, i) in features" :key="i" cols="6" md="4" class="box-img">
        <base-box-display-item
          v-bind="feature"
          :item="feature"
          :height="150"
          align="center"
          contain
          v-on:clickImage="onClick"
          v-on:clickTitle="onClick"
        ></base-box-display-item>
      </v-col>
    </v-row>
  </base-section>
</template>

<script>
export default {
  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    BaseBoxDisplayItem: () => import("../base/BoxDisplayItem")
  },
  props: {
    title: String,
    features: {
      type: Array,
      default: () => [
        {
          _id: 0,
          title: "",
          src: require("~/static/carlogo/logo-car-1.png"),
          link: "",
          description: "",
          showBorder: true
        }
      ]
    }
  },
  methods: {
    onClick(item) {
      this.$router.push({ name: "product-id", params: { id: item._id } });
    }
  }
};
</script>
<style scoped>
.box-img {
  height: 100% !important;
}

.box-img >>> .gt-logo {
  background-color: white !important;
}
.box-img >>> .box-description {
  background-color: white !important;
}

>>> .gt-inner {
  cursor: pointer;
}
</style>
