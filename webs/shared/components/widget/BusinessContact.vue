<template>
  <v-theme-provider :dark="dark">
    <div>
      <base-info :title="title" color="primary">
        <slot />
      </base-info>

      <template v-for="({ icon, text, title: t }, i) in business">
        <base-avatar-card
          :key="i"
          :icon="icon"
          :outlined="false"
          :title="!dense ? t : undefined"
          color="transparent"
          horizontal
          space="0"
        >{{text}}</base-avatar-card>

        <v-divider v-if="i + 1 !== business.length" :key="`divider-${i}`" class="my-2" />
      </template>
    </div>
  </v-theme-provider>
</template>

<script>
export default {
  name: "BaseBusinessContact",

  components: {
    BaseInfo: () => import("../base/Info"),
    BaseAvatarCard: () => import("../base/AvatarCard")
  },

  props: {
    dark: Boolean,
    dense: Boolean,
    title: String
  },

  data: () => ({
    business: [
      {
        icon: "mdi-map-marker-outline",
        title: "Address",
        text: "8553 N. Beach St. Ste. 227 Fort Worth, Texas 76137"
      },
      {
        icon: "mdi-cellphone",
        title: "Phone",
        text: "01 (800) 433 744 01 (800) 433 633"
      },
      {
        icon: "mdi-email",
        title: "Email",
        text: "john@vuetifyjs.com heather@vuetifyjs.com"
      }
    ]
  })
};
</script>
