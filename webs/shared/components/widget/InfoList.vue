<template>
  <v-theme-provider :dark="dark">
    <v-container class="pa-0">
      <v-row :justify="justify" no-gutters>
        <v-col cols="12">
          <v-list color="transparent" rounded dense>
            <div>
              <base-title :title="title" class="text-uppercase" space="1" />
              <base-divider :color="color" />
            </div>
            <v-list-item-group v-model="item" color="primary" class="mt-n1">
              <v-list-item v-for="(item, i) in items" :key="i">
                <v-list-item-icon v-if="item.icon">
                  <v-icon v-text="item.icon"></v-icon>
                </v-list-item-icon>
                <v-list-item-content>
                  <v-list-item-title v-text="item.title"></v-list-item-title>
                  <v-list-item-title v-if="item.subtitle" v-text="item.subtitle"></v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list-item-group>
          </v-list>
        </v-col>
      </v-row>
    </v-container>
  </v-theme-provider>
</template>

<script>
// Mixins
import Heading from "@/mixins/heading";

export default {
  name: "WidgetInfoList",

  components: {
    BaseTitle: () => import("../base/Title"),
    BaseDivider: () => import("../base/Divider"),
  },

  mixins: [Heading],

  props: {
    dark: Boolean,
    color: {
      type: String,
      default: "primary",
    },
    to: {
      type: String,
    },
    title: String,
    items: Array,
  },

  data: () => ({
    item: {
      title: "",
      subtitle: "",
      icon: "",
      to: "",
      action: (d) => {},
    },
  }),
};
</script>

<style scoped>
ol {
  list-style: none;
}
</style>
