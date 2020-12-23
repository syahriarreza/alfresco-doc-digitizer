<template>
  <base-section id="gallery">
    <base-section-heading title="Gallery"></base-section-heading>

    <v-responsive class="mx-auto">
      <v-container fluid>
        <widget-gallery :images="galleryImages"></widget-gallery>
      </v-container>
    </v-responsive>
  </base-section>
</template>

<script>
export default {
  name: "SectionGallery",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    WidgetGallery: () => import("~shared/components/widget/Gallery"),
  },

  data: () => ({}),
  computed: {
    galleryImages() {
      let datas = this.$store.state.event.gallery;
      return datas.map((_) => {
        _.src = this.$tool.getAssetURL(_.thumbnail_asset_id);
        return _;
      });
    },
  },
};
</script>
