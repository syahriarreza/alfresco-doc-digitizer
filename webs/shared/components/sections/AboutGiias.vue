<template>
  <base-section id="theme-features" class="px-5 py-10">
    <base-section-heading
      title="About GIIAS"
    >Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor.</base-section-heading>

    <v-responsive class="mx-auto" max-width="1350">
      <v-container>
        <v-row>
          <v-col v-for="card in cards" :key="card.title" cols="12" sm="6" md="3">
            <base-avatar-about align="center" v-bind="card" />
          </v-col>
        </v-row>
      </v-container>
    </v-responsive>
  </base-section>
</template>

<script>
export default {
  name: "SectionAboutGiias",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    BaseAvatarAbout: () => import("../widget/AvatarAbout")
  },

  data: () => ({
    cards: [
      {
        icon: "mdi-monitor",
        title: "Exhibition",
        text:
          "Efficiently unleash media information without cross-media value. Quickly maximize value timely deliverables schemas."
      },
      {
        icon: "mdi-car-side",
        title: "Premiere Car",
        text:
          "Efficiently unleash media information without cross-media value. Quickly maximize value timely deliverables schemas."
      },
      {
        icon: "mdi-traffic-light",
        title: "Test Drive",
        text:
          "Efficiently unleash media information without cross-media value. Quickly maximize value timely deliverables schemas."
      },
      {
        icon: "mdi-wechat",
        title: "Speaker Talks",
        text:
          "Efficiently unleash media information without cross-media value. Quickly maximize value timely deliverables schemas."
      }
    ]
  })
};
</script>

<style scoped>
section >>> .v-avatar.base-avatar__avatar > i.v-icon {
  color: var(--v-primary-base) !important;
}
section >>> .v-avatar.base-avatar__avatar {
  border: solid 7px;
  background: transparent !important;
  border-color: currentColor !important;
  color: var(--v-primary-base) !important;
}
section >>> .v-avatar.base-avatar__outline {
  background: transparent !important;
}
</style>