<template>
  <v-container>
    <v-row>
      <v-col v-for="(item, i) in items" :key="i" :cols="cols" class="box-img">
        <base-box-display-item
          v-bind="item"
          :item="item"
          :width="width"
          :height="height"
          align="center"
        ></base-box-display-item>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  components: {
    BaseBoxDisplayItem: () => import("../base/BoxDisplayItem")
  },
  props: {
    cols: {
      type: Number,
      default: 4
    },
    width: { type: Number, default: 100 },
    height: { type: Number, default: 100 },
    items: {
      type: Array,
      default: [
        {
          title: "",
          src: require("~/static/sponsor/logo-spongsor-1.jpg"),
          link: "",
          description: "",
          showBorder: true
        }
      ]
    }
  }
};
</script>

<style scoped>
</style>