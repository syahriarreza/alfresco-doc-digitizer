<template>
  <v-responsive>
    <v-row>
      <base-dialog
        :dialog="dialogShowHide"
        @hideDialog="refShowHideDialog"
        :title="title"
        :breadcrumbs="breadcrumbs"
      >
        <template v-slot:append-outer>
          <div class="container">
            <v-card elevation="24" class="mx-auto">
              <v-system-bar lights-out></v-system-bar>
              <v-carousel :show-arrows="true">
                <v-carousel-item v-for="(item,i) in product.src" :key="i" :src="item.src"></v-carousel-item>
              </v-carousel>
              <v-list two-line>
                <v-list-item>
                  <base-title
                    v-text="product.category + ' - ' + product.date"
                    class="text-uppercase"
                    space="1"
                  />
                </v-list-item>
                <v-list-item>
                  <v-list-item-content>
                    <base-subtitle v-html="product.content"></base-subtitle>
                  </v-list-item-content>
                </v-list-item>
              </v-list>
            </v-card>
          </div>
        </template>
      </base-dialog>
    </v-row>
  </v-responsive>
</template>


<script>
import { cloneDeep } from "lodash";
export default {
  props: {
    dialog: { type: Boolean, default: false },
  },
  components: {
    BaseDialog: () => import("../base/DialogFullscreen"),
    BaseTitle: () => import("../base/Title"),
    BaseSubtitle: () => import("../base/Subtitle"),
  },
  data() {
    return {
      colors: [
        "green",
        "secondary",
        "yellow darken-4",
        "red lighten-2",
        "orange darken-1",
      ],
      cycle: false,
      slides: ["First", "Second", "Third", "Fourth", "Fifth"],

      dialogShowHide: false,
      notifications: false,
      sound: true,
      widgets: false,
      title: "",
      breadcrumbs: [{ text: "Latest News" }, { text: "Detail News", to: "/" }],
    };
  },
  methods: {
    refShowHideDialog(val) {
      this.$emit("hidelog", false);
    },
  },
  computed: {
    product() {
      let datas = cloneDeep(this.$store.state.event.selectedEvent);
      this.title = datas.title;
      datas.src = [
        {
          src: datas.src,
        },
      ];
      return datas;
    },
  },
  watch: {
    dialog(newValue, oldValue) {
      this.dialogShowHide = this.dialog;
    },
  },
};
</script>