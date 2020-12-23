<template>
  <v-theme-provider dark>
    <base-section id="social" class="accent text-center" space="56">
      <base-icon class="mb-8">{{ icon }}</base-icon>

      <base-section-heading color="transparent" :title="title">{{ description }}</base-section-heading>

      <a
        class="d-inline-block mb-8"
        :href="url"
        style="text-decoration: none;"
        target="_blank"
      >{{ url }}</a>

      <div class="py-4" />

      <v-btn
        class="font-weight-bold"
        color="white"
        :href="url"
        light
        min-width="168"
        target="_blank"
        x-large
      >
        {{ btnName }}
        <v-icon right>{{ btnIcon }}</v-icon>
      </v-btn>
    </base-section>
  </v-theme-provider>
</template>

<script>
export default {
  name: "SectionSocialMedia",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
    BaseIcon: () => import("../base/Icon")
  },

  provide: {
    heading: { align: "center" }
  },

  props: {
    icon: {
      type: String,
      default: "mdi-twitter"
    },
    title: {
      type: String,
      default: "Social Media"
    },
    description: {
      type: String,
      default:
        "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Excepturi aspernatur recusandae aut repudiandae illo error obcaecati dolores voluptate, tempore."
    },
    url: {
      type: String,
      default: "https://twitter.com/vuetifyjs"
    },
    btnName: {
      type: String,
      default: "Follow Us"
    },
    btnIcon: {
      type: String,
      default: "mdi-open-in-new"
    }
  }
};
</script>
