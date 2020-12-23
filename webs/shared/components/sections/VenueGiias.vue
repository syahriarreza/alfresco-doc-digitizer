<template>
  <base-section-background id="venue" :src="src">
    <template v-slot:append-outer>
      <v-container fluid>
        <v-row class="info-body">
          <v-col cols="12" sm="6" md="6"></v-col>
          <v-col cols="12" sm="6" md="6">
            <div class="d-flex">
              <base-heading title="GIIAS" class="mr-3" />
              <base-heading class="accent--text" title="VENUE" />
            </div>
            <p>Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor</p>

            <base-title
              class="text-uppercase accent--text mb-3 font-weight-bold"
              space="3"
            >Show days</base-title>
            <p>22 October - 1 November 2020</p>

            <base-title class="text-uppercase accent--text mb-3 font-weight-bold" space="3">Venue</base-title>
            <p>Indonesia Convention Exhibition</p>
            <p>(ICE) BSD</p>

            <base-btn class="btn-buy accent mt-2 ml-n1">View Location</base-btn>
          </v-col>
        </v-row>
      </v-container>
    </template>
  </base-section-background>
</template>

<script>
export default {
  name: "SectionAboutGiias",

  components: {
    BaseSectionBackground: () => import("../base/SectionBackground"),
    BaseHeading: () => import("../base/Heading"),
    BaseTitle: () => import("../base/Title"),
    BaseBtn: () => import("../base/Btn")
  },
  props: {},
  data: () => ({
    src: require("~/static/venue.jpg"),
    height: 550
  })
};
</script>
<style scoped>
.info-body {
  padding: 40px;
  color: white;
}
.btn-buy.primary {
  height: 50px !important;
  border-radius: 30px;
}
</style>
