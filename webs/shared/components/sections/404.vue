<template>
  <base-section id="404">
    <base-heading title="404" />

    <base-subheading class="primary--text" title="Oh!" />

    <base-subheading space="8" title="Looks like this page doesn't exist" />

    <base-img
      :src="require('@/static/logo.svg')"
      class="mx-auto mb-12"
      contain
      height="350"
      width="250"
    />

    <div class="text-center">
      <base-btn :to="{ name: 'Home' }">Get me Out of Here</base-btn>
    </div>
  </base-section>
</template>

<script>
export default {
  name: "FourOhFour",

  components: {
    BaseSection: () => import("../base/Section"),
    BaseHeading: () => import("../base/Heading"),
    BaseSubheading: () => import("../base/Subheading"),
    BaseImg: () => import("../base/Img"),
    BaseBtn: () => import("../base/Btn")
  },

  provide: {
    heading: { align: "center" }
  }
};
</script>
