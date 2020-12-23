<template>
  <base-section id="ticket" class="grey lighten-5">
    <base-section-heading title="BUY TICKET"></base-section-heading>

    <v-container>
      <v-row class="justify-center">
        <v-col v-for="(ticket, i) in tickets" :key="i" cols="12" sm="6" md="3" class="box-img">
          <widget-ticket-item
            v-bind="ticket"
            :item="ticket"
            align="center"
            :buttonLabel="buttonLabel"
            @buyTicket="SubmitbuyTicket"
          ></widget-ticket-item>
        </v-col>
      </v-row>
    </v-container>
  </base-section>
</template>

<script>
import { cloneDeep } from "lodash";
export default {
  name: "SectionTicket",

  components: {
    WidgetTicketItem: () => import("../widget/TicketItem"),
    BaseSection: () => import("../base/Section"),
    BaseSectionHeading: () => import("../base/SectionHeading"),
  },

  data: () => ({
    buttonLabel: "BUY TICKET",
  }),
  methods: {
    SubmitbuyTicket(item, qty) {
      this.$emit("buyTicket", item, qty);
    },
  },
  computed: {
    tickets() {
      let datas = cloneDeep(this.$store.state.ticket.tickets);
      return datas.map((_) => {
        _.qtyList = [];
        for (let i = 0; i < _.max_order; i++) {
          _.qtyList.push((i + 1).toString());
        }
        return _;
      });
    },
  },
};
</script>
<style scoped>
.box-img {
  height: 100% !important;
}
</style>
