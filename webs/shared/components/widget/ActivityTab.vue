<template>
  <v-tabs background-color="primary" class="elevation-2" dark>
    <v-tabs-slider></v-tabs-slider>

    <v-tab
      v-for="(actArray, date_from) in groupedActivities"
      :key="date_from"
      :href="`#actTab-${date_from}`"
    >{{ $moment(date_from).format('DD MMM YYYY') }}</v-tab>

    <v-tab-item
      v-for="(actArray, date_from) in groupedActivities"
      :key="date_from"
      :value="'actTab-'+date_from"
    >
      <v-expansion-panels>
        <v-expansion-panel v-for="(actItem, actItemIdx) in actArray" :key="actItemIdx">
          <v-expansion-panel-header>
            <v-list-item-content v-if="$vuetify.breakpoint.smAndDown">
              <v-list-item-subtitle class="text--primary" v-text="actItem.name"></v-list-item-subtitle>
              <v-list-item-subtitle
                v-text="$tool.int2time(actItem.hour_from) + ' - ' + $tool.int2time(actItem.hour_to)"
              ></v-list-item-subtitle>
            </v-list-item-content>
            <v-row v-else>
              <v-col
                cols="2"
              >{{ $tool.int2time(actItem.hour_from) + ' - ' + $tool.int2time(actItem.hour_to) }}</v-col>
              <v-col>{{ actItem.name }}</v-col>
            </v-row>
          </v-expansion-panel-header>
          <!-- TODO : letter-spacing -->
          <!-- <v-expansion-panel-content style="letter-spacing:0px">{{ actItem.description }}</v-expansion-panel-content> -->
          <v-expansion-panel-content>
            <base-body v-html="actItem.description"></base-body>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-tab-item>
  </v-tabs>
</template>

<script>
import { groupBy } from "lodash";

export default {
  components: {
    BaseBody: () => import("../base/Body"),
  },
  props: {
    activities: {
      type: Array,
      default: () => [
        {
          title: "No Title",
          description: "No Description",
          date_from: new Date(),
          date_to: new Date(),
          hour_from: 900,
          hour_to: 1000,
          is_free: false,
        },
      ],
    },
  },
  computed: {
    groupedActivities() {
      return groupBy(this.activities, (_) => {
        return this.$moment(_.date_from).format("DD MMM YYYY");
      });
    },
  },
};
</script>

<style lang="scss" scoped>
</style>