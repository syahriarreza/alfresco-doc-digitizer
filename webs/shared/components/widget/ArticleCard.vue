<template>
  <v-card>
    <v-card-title v-if="headerSrc" style="padding: 0 0 16px 0;">
      <v-img :src="headerSrc"></v-img>
    </v-card-title>
    <v-card-text>
      <slot></slot>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  props: {
    headerSrc: {
      type: String
    }
  }
};
</script>

<style lang="scss" scoped>
</style>