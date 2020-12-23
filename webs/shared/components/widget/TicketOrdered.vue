<template>
  <v-card class="mx-auto" outlined>
    <v-list-item three-line class="primary">
      <v-list-item-avatar tile size="80" color="grey" class="logo-card">
        <v-img contain :src="ticketLogoSrc" height="80" width="80"></v-img>
      </v-list-item-avatar>
      <v-list-item-content class="white--text h100">
        <v-list-item-title class="mb-1" v-text="event.name">GIIAS Jakarta-Daily Ticket</v-list-item-title>
        <v-list-item-subtitle class="white--text" v-text="event.location">ICE Building</v-list-item-subtitle>
        <v-list-item-subtitle
          class="white--text"
          v-text="this.$tool.dateFormat(
                event.date_from,
                'DD MMMM YYYY'
              ) "
        >10 Octocber 2020</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>
    <v-card-text class="text--primary color-body-ticket">
      <div class="d-flex flex-no-wrap justify-space-between">
        <v-list-item-content>
          <v-list-item-subtitle>Booking ID</v-list-item-subtitle>
          <v-list-item-subtitle v-text="item.sales_id">84397234799920</v-list-item-subtitle>
          <v-list-item-subtitle class="warning--text">Please show this QR Code to check in</v-list-item-subtitle>
          <div>
            <base-btn
              :class="'btn-print-ticket'"
              depressed
              small
              :color="''"
              :xlarge="false"
            >Print My Ticket</base-btn>
          </div>
        </v-list-item-content>
        <v-list-item-avatar tile size="80" color="grey">
          <img class="border-qr-code" :src="item.src" height="80" width="80" />
        </v-list-item-avatar>
      </div>
    </v-card-text>
    <v-card-text class="text--primary">
      <div class="d-flex flex-no-wrap justify-space-between footer-ticket">
        <img class="border-bulet" src="~static/bulet.svg" />
        <v-list-item-content>
          <v-list-item-subtitle>Ticket Holder</v-list-item-subtitle>
          <v-list-item-subtitle v-text="item.name">Anggun Cahyaningtiyas</v-list-item-subtitle>
          <v-list-item-subtitle v-text="item.phone">081234783999</v-list-item-subtitle>
          <v-list-item-subtitle v-text="item.email">anggun@mail.com</v-list-item-subtitle>
        </v-list-item-content>
      </div>
    </v-card-text>
  </v-card>
</template>

<script>
// Mixins
import Heading from "@/mixins/heading";

export default {
  name: "TicketOrdered",

  components: {
    BaseBtn: () => import("../base/Btn"),
    BaseTitle: () => import("../base/Title"),
    BaseIcon: () => import("../base/Icon"),
    BaseSelectField: () => import("../base/SelectField"),
  },

  mixins: [Heading],

  props: {
    item: {
      type: Object,
      default: {
        ticket_date: "10 Octocber 2020",
        sales_id: "2007271542488NUGO4Q1",
        name: "Anggun Cahyaningtiyas",
        phone: "081",
        email: "anggun@mail.com",
      },
    },
  },
  data: () => ({
    ticketLogoSrc: require("~/static/logo-giias.png"),
  }),
  computed: {
    event() {
      return this.$store.state.event.event;
    },
  },
  methods: {},
};
</script>

<style scoped>
.h100 {
  height: 100px;
}
.color-body-ticket {
  background-color: aliceblue !important;
}
.btn-print-ticket {
  background-color: white !important;
  font-size: 10px;
  border-radius: 4px;
  margin-top: 5px;
}
.footer-ticket {
  position: relative;
}
.border-bulet {
  position: absolute;
  top: -25px;
  width: 100%;
}
.border-qr-code {
  position: absolute;
  width: 100%;
}
.v-list-item__subtitle {
  white-space: pre-wrap;
}
.logo-card {
  background-color: #201162 !important;
  border: 1px solid #ffff;
  height: 80px;
  width: 80px;
}
</style>