<template>
  <base-section id="theme-features" class="grey lighten-5">
    <base-section-heading title="Latest News"></base-section-heading>

    <v-container>
      <v-row>
        <v-col v-for="(item, i) in newsInfo" :key="i" cols="12" sm="6" md="6" class="box-img">
          <base-news v-bind="item" :item="item" :height="350" @latestNews="detailNews"></base-news>
        </v-col>
      </v-row>
    </v-container>
    <base-dialog-detailLatestNews :dialog="dialogLatestNews" @hidelog="hidelog" />
  </base-section>
</template>

<script>
export default {
  name: "SectionLatestNews",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    BaseNews: () => import("../base/News"),
    BaseDialogDetailLatestNews: () =>
      import("../widget/DialogDetailLatestNews"),
  },

  data: () => ({
    dialogLatestNews: false,
  }),
  computed: {
    newsInfo() {
      let datas = this.$store.state.event.news;
      return datas.map((_) => {
        _.src = this.$tool.getAssetURL(_.featured_image_asset_id);
        _.date = this.$tool.dateFormat(_.updated_date, "DD MMM, YYYY");
        _.link = "";
        _.showBorder = false;
        return _;
      });
    },
  },
  methods: {
    detailNews(item) {
      this.$store.dispatch("event/setSelectedEvent", item);
      this.dialogLatestNews = true;
    },
    hidelog(item) {
      this.dialogLatestNews = item;
    },
  },
};
</script>
<style scoped>
.box-img {
  height: 100% !important;
}
</style>
