<template>
  <v-card flat tile>
    <v-card-text>
      <base-subheading :align="'left'" space="2" color="primary">
        <span class="schedulle-heading">Show days</span>
      </base-subheading>
      <base-body>{{ $moment(schedule.date_from).format('DD MMMM YYYY') }} - {{ $moment(schedule.date_to).format('DD MMMM YYYY') }}</base-body>
      <base-subheading :align="'left'" space="2">
        <p class="schedulle-heading">Venue</p>
      </base-subheading>
      <base-body>{{ schedule.location }}</base-body>
    </v-card-text>
  </v-card>
</template>

<script>
export default {
  components: {
    BaseSubheading: () => import("~shared/components/base/Subheading"),
    BaseBody: () => import("~shared/components/base/Body"),
  },
  props: {
    schedule: {
      type: Object,
      default: {
        date_from: new Date(),
        date_to: new Date(),
        location: "No Venue",
      },
    },
  },
};
</script>

<style scoped>
.schedulle-heading {
  font-family: "Arial Bold", "Arial";
  font-weight: 700;
  font-size: 20px;
  color: #f4bc29;
  margin-bottom: 0;
}
</style>