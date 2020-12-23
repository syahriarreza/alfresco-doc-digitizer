<template>
  <v-theme-provider dark>
    <base-section id="info-alt" class="grey darken-4" space="48">
      <v-responsive class="mx-auto" max-width="1400">
        <v-container fluid>
          <v-row>
            <v-col cols="12" md="3">
              <base-img
                :src="require('@/static/logo-giias.png')"
                contain
                max-width="128"
                min-height="78"
                width="100%"
              />

              <template v-for="(feature, i) in features">
                <base-info-detail
                  :key="i"
                  v-bind="feature"
                  align="left"
                  horizontal
                  :space="4"
                  :spaceTitle="0"
                >{{feature.subtitle}}</base-info-detail>
              </template>
              <base-btn class="mt-6" color="white" outlined>More Info</base-btn>
            </v-col>

            <v-col cols="12" md="3">
              <widget-info-list v-bind="about"></widget-info-list>
              <widget-info-list v-bind="visitor"></widget-info-list>
              <widget-info-list v-bind="exhibitor"></widget-info-list>
            </v-col>
            <v-col cols="12" md="3">
              <widget-info-list v-bind="program"></widget-info-list>
            </v-col>

            <v-col cols="12" md="3">
              <widget-info-list v-bind="series"></widget-info-list>
              <!-- <widget-instagram /> -->
            </v-col>
          </v-row>
        </v-container>
      </v-responsive>
    </base-section>
  </v-theme-provider>
</template>

<script>
export default {
  name: "SectionFooterInfo",

  components: {
    WidgetInfoList: () => import("../widget/InfoList"),
    WidgetInstagram: () => import("../widget/Instagram"),
    WidgetBusinessContact: () => import("../widget/BusinessContact"),
    BaseSection: () => import("../base/Section"),
    BaseInfo: () => import("../base/Info"),
    BaseImg: () => import("../base/Img"),
    BaseBody: () => import("../base/Body"),
    BaseBtn: () => import("../base/Btn"),
    BaseTitle: () => import("../base/Title"),
    BaseInfoDetail: () => import("../base/InfoDetail"),
  },
  props: {},
  data: () => ({
    features: [
      {
        title: "Show days",
        subtitle: "22 October - 1 November 2020",
      },
      {
        title: "Venue",
        subtitle: "Indonesia Convention Exhibition (ICE) BSD",
      },
    ],
    about: {
      title: "ABOUT",
      items: [{ title: "CONTACT US", to: "#", icon: "mdi-chevron-right" }],
    },
    visitor: {
      title: "VISITOR",
      items: [
        { title: "DATE & VENUE", to: "#", icon: "mdi-chevron-right" },
        { title: "LOCATION & DIRECTIONS", to: "#", icon: "mdi-chevron-right" },
        { title: "FLOOR PLAN", to: "#", icon: "mdi-chevron-right" },
        { title: "EXHIBITOR LIST", to: "#", icon: "mdi-chevron-right" },
        { title: "SHUTTLE BUS SERVICES", to: "#", icon: "mdi-chevron-right" },
        { title: "GIIAS FACILITIES", to: "#", icon: "mdi-chevron-right" },
      ],
    },
    exhibitor: {
      title: "EXHIBITOR",
      items: [
        { title: "FLOOR PLAN", to: "#", icon: "mdi-chevron-right" },
        { title: "EXHIBITOR LIST", to: "#", icon: "mdi-chevron-right" },
      ],
    },
    program: {
      title: "PROGRAMS",
      items: [
        { title: "THE 14th IIAC", to: "#", icon: "mdi-chevron-right" },
        { title: "BEYOND MOBILITY AREA", to: "#", icon: "mdi-chevron-right" },
        { title: "TEST DRIVE", to: "#", icon: "mdi-chevron-right" },
        { title: "TEST RIDE", to: "#", icon: "mdi-chevron-right" },
        { title: "BLOOD DONATION", to: "#", icon: "mdi-chevron-right" },
        { title: "CSR PAVILION", to: "#", icon: "mdi-chevron-right" },
        { title: "DEMO AREA", to: "#", icon: "mdi-chevron-right" },
        { title: "EDUCATION DAY", to: "#", icon: "mdi-chevron-right" },
        {
          title: "FAVORITE CAR & MOTORCYCLE",
          to: "#",
          icon: "mdi-chevron-right",
        },
        { title: "FAVORITE BOOTH", to: "#", icon: "mdi-chevron-right" },
        { title: "MISS AUTO SHOW", to: "#", icon: "mdi-chevron-right" },
        { title: "WORKING BAY", to: "#", icon: "mdi-chevron-right" },
      ],
    },
    series: {
      title: "GIIAS THE SERIES",
      items: [
        { title: "GIIAS SURABAYA", to: "#", icon: "mdi-chevron-right" },
        { title: "GIIAS MAKASSAR", to: "#", icon: "mdi-chevron-right" },
        { title: "GIIAS MEDAN", to: "#", icon: "mdi-chevron-right" },
      ],
    },
  }),
};
</script>

<style lang="sass">
#info-alt a
  text-decoration: none
</style>
