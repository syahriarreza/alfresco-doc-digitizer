<template>
  <v-theme-provider :dark="dark">
    <v-container class="pa-0">
      <v-row :justify="justify" no-gutters>
        <v-col v-if="icon" :class="`text-${align}`" cols="12" class="mb-4">
          <base-avatar
            v-if="icon"
            :color="color"
            :dark="dark"
            :icon="icon"
            :outlined="outlined"
            :size="size"
            class="mb-3"
          />
        </v-col>

        <v-col v-if="title || subtitle" :cols="callout ? 9 : 12">
          <base-subtitle v-if="subtitle" :title="subtitle" space="1" />

          <base-title :title="title" class="text-uppercase" space="1" />
          <base-divider :color="color" />

          <base-body v-if="text || $slots.default" :text="text" space="6">
            <slot></slot>
          </base-body>
        </v-col>

        <v-col v-if="callout" cols="2">
          <div class="display-3 grey--text text--lighten-4 font-weight-bold pr-8" v-text="callout" />
        </v-col>
      </v-row>
    </v-container>
  </v-theme-provider>
</template>

<script>
// Mixins
import Heading from "@/mixins/heading";

export default {
  name: "BaseAvatarCard",

  components: {
    BaseBody: () => import("../base/Body"),
    BaseDivider: () => import("../base/Divider"),
    BaseIcon: () => import("../base/Icon"),
    BaseTitle: () => import("../base/Title"),
    BaseSubtitle: () => import("../base/Subtitle"),
    BaseAvatar: () => import("../base/Avatar")
  },

  mixins: [Heading],

  props: {
    dark: Boolean,
    callout: String,
    color: {
      type: String,
      default: "primary"
    },
    icon: String,
    subtitle: String,
    text: String,
    title: String,
    size: {
      type: [Number, String],
      default: 120
    },
    outlined: {
      type: Boolean,
      default: true
    }
  }
};
</script>
