<template>
  <v-container>
    <v-row>
      <v-col v-for="(feature, i) in features" :key="i" cols="6" md="3" class="box-img">
        <base-box-display-item
          v-bind="feature"
          :item="feature"
          :width="width"
          :height="height"
          :contain="contain"
          :marginBot="marginBot"
          :showText="showText"
          align="center"
          v-on:clickImage="onClick"
          v-on:clickTitle="onClick"
        ></base-box-display-item>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  components: {
    BaseBoxDisplayItem: () => import("../base/BoxDisplayItem"),
  },
  props: {
    width: { type: Number, default: 100 },
    height: { type: Number, default: 100 },
    contain: { type: Boolean, default: false },
    marginBot: { type: Boolean, default: false },
    showText: { type: Boolean, default: true },
    features: {
      type: Array,
      default: () => [
        {
          eventid: "",
          tenantid: "",
          title: "",
          src: require("~/static/carlogo/logo-car-1.png"),
          link: "",
          description: "",
          showBorder: true,
        },
      ],
    },
  },
  methods: {
    onClick(item) {
      this.$router.push({
        name: "event-eventid-tenantid",
        params: { eventid: item.event_id, tenantid: item._id },
      });
    },
  },
};
</script>

<style scoped>
>>> .gt-inner {
  cursor: pointer;
}
</style>