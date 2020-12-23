<template>
  <base-section id="theme-features">
    <base-section-heading :title="title">{{ description }}</base-section-heading>

    <v-container>
      <v-row>
        <v-col v-for="(feature, i) in features" :key="i" cols="12" md="6">
          <base-avatar-card v-bind="feature" align="left" horizontal>{{ feature.description }}</base-avatar-card>
        </v-col>
      </v-row>
    </v-container>
  </base-section>
</template>

<script>
export default {
  name: "SectionThemeFeatures",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    BaseAvatarCard: () => import("../base/AvatarCard")
  },

  props: {
    title: {
      type: String,
      default: "Theme Features"
    },
    description: {
      type: String,
      default:
        "Lorem ipsum dolor sit amet consectetur adipisicing elit. Libero, in! Asperiores, impedit libero. Veniam rerum saepe unde nihil possimus quibusdam esse accusamus mollitia magni fuga."
    },
    features: {
      type: Array,
      default: () => [
        {
          color: "primary",
          dark: true,
          title: "Easily Customizable",
          icon: "mdi-pencil-box-outline",
          description:
            "Lorem ipsum dolor sit amet consectetur adipisicing elit. Veritatis illum veniam cum dolores ratione commodi beatae quas maxime, laboriosam excepturi solut!"
        }
      ]
    }
  }
};
</script>
