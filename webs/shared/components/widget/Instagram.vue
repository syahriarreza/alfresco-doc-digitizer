<template>
  <base-info title="Small Gallery">
    <v-row dense>
      <v-col v-for="n in 6" :key="n" cols="auto">
        <base-img :src="require(`@/static/insta-${n}.jpg`)" color="grey lighten-1" width="96" tile />
      </v-col>
    </v-row>
  </base-info>
</template>

<script>
export default {
  name: "Instagram",

  components: {
    BaseInfo: () => import("../base/Info"),
    BaseImg: () => import("../base/Img")
  }
};
</script>
