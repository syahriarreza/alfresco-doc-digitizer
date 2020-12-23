<template>
  <v-app>
    <v-main>
      <nuxt />
    </v-main>
  </v-app>
</template>

<script>
export default {
  components: {},
  data() {
    return {
      isOnDevEnv: false,
    }
  },
  computed: {},
  methods: {},
  mounted() {
    this.isOnDevEnv = process.env.NODE_ENV === 'development'
  },
}
</script>

<style>
.vuedl-notification {
  top: 72px !important;
  right: 12px !important;
}

.v-btn.primary-gradient {
  color: white !important;
  background: var(--v-primary-base);
  background: linear-gradient(
    90deg,
    var(--v-primary-base) 30%,
    var(--v-tertiary-base) 100%
  );
}
</style>
