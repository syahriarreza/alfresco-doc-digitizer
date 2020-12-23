<template>
  <v-card>
    <v-card-title>
      <h4>Booking Details</h4>
    </v-card-title>
    <v-divider />
    <v-card-text class="py-0">
      <v-row>
        <v-col>
          <h3 class="color-black mb-1">Booking ID</h3>
          <h4>84397234799920</h4>
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <h4 class="color-black">4x GIIAS Jakarta - Personal Tickets</h4>
        </v-col>
      </v-row>
    </v-card-text>
    <v-divider />
    <v-card-text class="py-3">
      <v-row>
        <v-col class="py-0">
          <h3 class="color-black">Participant's</h3>
        </v-col>
      </v-row>
      <v-sheet
        v-for="(p, i) in participants"
        :key="i"
        :elevation="0"
        outlined
        rounded
        class="pa-5 my-2"
      >
        <p class="ma-0">{{ p.name }}</p>
        <p class="ma-0">{{ p.phone }}</p>
        <p class="ma-0">{{ p.email }}</p>
      </v-sheet>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  components: {
    BaseSubheading: () => import("~shared/components/base/Subheading"),
    BaseBoxDisplayItem: () => import("~shared/components/base/BoxDisplayItem"),
    BaseBody: () => import("~shared/components/base/Body"),
    BaseView: () => import("~shared/components/base/View"),
    BaseBtn: () => import("~shared/components/base/Btn"),
  },
  data() {
    return {
      participants: [
        {
          name: "Anggun Cahyaningtiyas",
          email: "081234728293",
          phone: "anggun@mail.com",
        },
        {
          name: "Angger Setyo",
          email: "081238934839",
          phone: "angger@mail.com",
        },
        {
          name: "Muhammad Cahyono",
          email: "0815933992003",
          phone: "cahyono@mail.com",
        },
        {
          name: "Agung Tiyo Nasution",
          email: "0811223474890",
          phone: "agungti@mail.com",
        },
      ],
    };
  },
};
</script>

<style scoped>
.color-black {
  color: black;
}
</style>