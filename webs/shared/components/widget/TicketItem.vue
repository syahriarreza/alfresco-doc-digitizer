<template>
  <div class="pt-2 gt-col">
    <div class="gt-inner">
      <div :class="item.is_recomended == false ? 'gt-event-ticket':'priority primary'">
        <div class="gt-ticket-inner">
          <base-title
            :class="item.is_recomended == false ? 'text-uppercase primary--text':'text-uppercase priority--text'"
            space="3"
          >{{item.name}}</base-title>
          <div class="gt-price">
            <span class="amount">
              <span
                :class="item.is_recomended == false ? 'currencySymbol':'currencySymbol--priority'"
              >{{item.curr}}</span>
              <span
                :class="item.is_recomended == false ? 'primary--text':'price--priority'"
              >{{$tool.formatDenomination(item.price).number}}</span>
              <span
                :class="item.is_recomended == false ? 'currencySymbol':'currencySymbol--priority'"
              >{{$tool.formatDenomination(item.price).denom}}</span>
            </span>
          </div>
          <div class="gt-ticket-features">
            <div v-if="typeof item.notes == 'object'">
              <template v-for="(desc,index) in item.notes">
                <div v-bind:key="'desc-' + index" class="d-flex desc-ticket">
                  <p :class="item.is_recomended == true ? 'priority--text':''">
                    <i class="mdi mdi-check-circle check-tiket" aria-hidden="true"></i>
                    {{ desc }}
                  </p>
                </div>
              </template>
            </div>
            <div v-else class="d-flex desc-ticket">
              <p v-html="item.notes" :class="item.is_recomended == true ? 'priority--text':''"></p>
            </div>
          </div>
          <div class="gt-buy-now gt-woocommerce-button">
            <div class="gt-quantity">
              <base-select-field
                :items="item.qtyList"
                label="Quantity"
                v-model="qtyTicket"
                class="ddl-qty-tiket"
              ></base-select-field>
            </div>
          </div>
        </div>
      </div>
      <base-btn
        class="btn-buy"
        :color="item.is_recomended == true ? 'btn-priority':'primary'"
        @click="buyTicket(item)"
      >{{buttonLabel}}</base-btn>
    </div>
  </div>
</template>

<script>
// Mixins
import Heading from "@/mixins/heading";

export default {
  name: "BaseTicketItem",

  components: {
    BaseBtn: () => import("../base/Btn"),
    BaseTitle: () => import("../base/Title"),
    BaseIcon: () => import("../base/Icon"),
    BaseSelectField: () => import("../base/SelectField"),
  },

  mixins: [Heading],

  props: {
    item: Object,
    buttonLabel: { type: String, default: "BUY NOW" },
  },
  data: () => ({
    qtyTicket: 0,
  }),
  computed: {},
  methods: {
    buyTicket(item) {
      this.$emit("buyTicket", item, this.qtyTicket);
    },
  },
};
</script>

<style scoped>
.gt-col {
  width: 100%;
  position: relative;
  display: inline-block;
}
.gt-col > .gt-inner {
  display: inline-block;
  height: 100%;
  width: 100%;
}
.gt-event-ticket {
  background-color: #fff;
  box-shadow: 0 0 30px rgba(0, 0, 0, 0.1);
  padding: 30px;
  border-radius: 5px 5px 0px 0px;
  text-align: center;
  min-height: 100px;
}
.priority {
  box-shadow: 0 0 30px rgba(0, 0, 0, 0.1);
  padding: 30px;
  border-radius: 5px 5px 0px 0px;
  text-align: center;
  min-height: 100px;
}
.priority--text {
  color: #fff;
}
.gt-price {
  font-weight: 600;
  font-size: 3rem;
  line-height: 1;
  margin: 15px 0 0;
}
.currencySymbol {
  font-size: 20px;
  color: #180000;
}
.gt-ticket-features {
  margin: 30px 0 0;
  color: #888;
}

.currencySymbol--priority {
  font-size: 20px;
  color: #fff;
}
.gt-ticket-features--priority {
  margin: 30px 0 0;
  color: #fff;
}
.price--priority {
  color: #ffc80d;
}
.btn-priority {
  background-color: #ffc80d !important;
  width: 100%;
  color: #fff;
  border-radius: 0px 0px 5px 5px;
}
.btn-buy.primary {
  width: 100%;
  height: 50px !important;
  border-radius: 0px 0px 5px 5px;
}
.check-tiket {
  padding-right: 5px;
  font-size: 15px;
  font-weight: 400;
  /* color: #132877; */
}
.desc-ticket {
  justify-content: center;
}
</style>
<style>
.gt-logo .v-responsive__sizer {
  padding-bottom: 0% !important;
}
</style>