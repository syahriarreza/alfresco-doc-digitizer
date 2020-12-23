<template>
  <v-overlay :value="overlay" :dark="false" color="#ffffff" opacity="1" class="v-overlay-loading">
    <img src="~static/loading-Infinity.svg" height="100" width="100" />
  </v-overlay>
</template>

<script>
export default {
  props: {
    overlay: { type: Boolean, default: false },
  },
  components: {},
  data: () => ({
    absolute: true,
  }),
  watch: {},
  methods: {},
  computed: {},
};
</script>
<style scoped>
</style>