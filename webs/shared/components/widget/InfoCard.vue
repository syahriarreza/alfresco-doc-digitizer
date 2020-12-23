<template>
  <div>
    <base-img :src="src" contain max-width="128" min-height="78" width="100%" />

    <base-title size="body-1" space="4" :title="title" weight="regular" />

    <base-body>{{ body }}</base-body>

    <base-btn class="mb-12" color="white" outlined>{{ btnTitle }}</base-btn>
  </div>
</template>

<script>
export default {
  components: {
    BaseImg: () => import("../base/Img"),
    BaseTitle: () => import("../base/Title"),
    BaseBody: () => import("../base/Body"),
    BaseBtn: () => import("../base/Btn")
  },
  props: {
    src: String,
    title: String,
    body: String,
    btnTitle: String
  }
};
</script>

<style lang="scss" scoped>
</style>